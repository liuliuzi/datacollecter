package main
import (
	"flag"
	"fmt"
	"os"
	"net/http"
	"time"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-swagger12"
	"../pkg/api"
)

func main() {
	localIP   := flag.String("localIP", "0.0.0.0", "local IP")
	localPort := flag.String("localPort", "8060", "local Port opened")
	liveTime  := flag.String("liveTime", "60s", "metric sample live time")

	flag.Parse()

	liveTimeDur,err := time.ParseDuration(*liveTime)
	if err!=nil{
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	metricsApi := api.MetricsService{make(map[string]api.Metric),liveTimeDur}
    metricsApi.Register()

    config := swagger.Config{
        WebServices:    restful.RegisteredWebServices(),
        WebServicesUrl: "http://"+*localIP+":"+*localPort,
        ApiPath:        "/apidocs.json",
        //SwaggerPath:     "/apidocs/",
        //SwaggerFilePath: "/Users/emicklei/Projects/swagger-ui/dist"
    }

    swagger.InstallSwaggerService(config)
    err = http.ListenAndServe(*localIP+":"+*localPort, nil)
    if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}