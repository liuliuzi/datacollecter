package app
import (
	"net/url"
	"net/http"
	"time"
	"fmt"
	"errors"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-swagger12"
	"github.com/liuliuzi/datacollecter/pkg/api"
	"github.com/liuliuzi/datacollecter/cmd/app/options"
	"k8s.io/heapster/common/flags"
	"k8s.io/heapster/metrics/sinks/influxdb"
	"k8s.io/heapster/metrics/core"
)



func Run(opt *options.DatacollecterRunOptions) error {
//func Run(mode int,localIP string,localPort string,liveTime string ,influxdbIP string,influxdbPort string) error {
	liveTimeDur,err := time.ParseDuration(opt.LiveTime)
	if err!=nil{
		return err
	}
	if opt.Mode!=0 && opt.Mode!=1{
		errors.New("invalid mode")
		return err
	}

	var influxdbSink core.DataSink
	if opt.Mode==1{
		influxdbURLString:="http://"+opt.InfluxdbIP+":"+opt.InfluxdbPort
		influxdbURL,err  :=url.Parse(influxdbURLString)
		if err != nil {
			return err
		}
		uri              :=flags.Uri{"influxdb",*influxdbURL}
		influxdbSink,err  =influxdb.CreateInfluxdbSink(&uri.Val)
		if err != nil {
			return err
		}
	}
	metricsApi := api.MetricsService{make(map[string]api.Metric),liveTimeDur,opt.Mode,influxdbSink}
    metricsApi.Register()

    config := swagger.Config{
        WebServices:    restful.RegisteredWebServices(),
        WebServicesUrl: "http://"+opt.LocalIP+":"+opt.LocalPort,
        ApiPath:        "/apidocs.json",
    }

    swagger.InstallSwaggerService(config)
    if opt.Mode==1{
    	fmt.Println("start app in mode directly")
    }else{
    	fmt.Println("start app in mode indirectly")
    }
    err = http.ListenAndServe(opt.LocalIP+":"+opt.LocalPort, nil)
    if err != nil {
		return err
	}
	return nil
}