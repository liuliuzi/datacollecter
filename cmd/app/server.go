package app
import (
	"net/http"
	"net/url"
	"time"
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-swagger12"
	"github.com/liuliuzi/datacollecter/pkg/api"
	"k8s.io/heapster/common/flags"
	"k8s.io/heapster/metrics/sinks/influxdb"
	"k8s.io/heapster/metrics/core"
)




func Run(mode int,localIP string,localPort string,liveTime string ,influxdbIP string,influxdbPort string) error {
	liveTimeDur,err := time.ParseDuration(liveTime)
	if err!=nil{
		return err
	}

	var InfluxdbSink core.DataSink
	if mode==1{
		influxdbURLString:="http://"+influxdbIP+":"+influxdbPort
		influxdbURL,err  :=url.Parse(influxdbURLString)
		if err != nil {
			return err
		}
		uri              :=flags.Uri{"influxdb",*influxdbURL}
		InfluxdbSink,err  =influxdb.CreateInfluxdbSink(&uri.Val)
		if err != nil {
			return err
		}
	}

	metricsApi := api.MetricsService{make(map[string]api.Metric),liveTimeDur,mode,InfluxdbSink}
    metricsApi.Register()

    config := swagger.Config{
        WebServices:    restful.RegisteredWebServices(),
        WebServicesUrl: "http://"+localIP+":"+localPort,
        ApiPath:        "/apidocs.json",
        //SwaggerPath:     "/apidocs/",
        //SwaggerFilePath: "/Users/emicklei/Projects/swagger-ui/dist"
    }

    swagger.InstallSwaggerService(config)
    if mode==1{
    	fmt.Println("start app in mode 1")
    }else{
    	fmt.Println("start app in mode 0")
    }
    err = http.ListenAndServe(localIP+":"+localPort, nil)
    if err != nil {
		return err
	}
	return nil
}