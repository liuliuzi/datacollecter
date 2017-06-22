package app
import (
	"net/http"
	"time"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-swagger12"
	"github.com/liuliuzi/datacollecter/pkg/api"
)



func Run(localIP string,localPort string,liveTime string) error {
	liveTimeDur,err := time.ParseDuration(liveTime)
	if err!=nil{
		return err
	}

	metricsApi := api.MetricsService{make(map[string]api.Metric),liveTimeDur}
    metricsApi.Register()

    config := swagger.Config{
        WebServices:    restful.RegisteredWebServices(),
        WebServicesUrl: "http://"+localIP+":"+localPort,
        ApiPath:        "/apidocs.json",
        //SwaggerPath:     "/apidocs/",
        //SwaggerFilePath: "/Users/emicklei/Projects/swagger-ui/dist"
    }

    swagger.InstallSwaggerService(config)
    err = http.ListenAndServe(localIP+":"+localPort, nil)
    if err != nil {
		return err
	}
	return nil
}