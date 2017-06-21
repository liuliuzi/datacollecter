package api
import (
	"net/http"
	"github.com/emicklei/go-restful"
	"fmt"
	"time"
)

type MetricsService struct {
	MetricsCache  map[string]Metric
	LiveSeconds    time.Duration
}

type Metric struct {
	MetricValue string
	Timestamp time.Time
}

func (ms MetricsService)valideTime(sampleTime time.Time) bool{
	curTime:=time.Now()
	if curTime.After(sampleTime.Add(ms.LiveSeconds)){
		fmt.Println(curTime.String()+" is not after " +sampleTime.Add(ms.LiveSeconds).String())
		return false
	}else{
		return true
	}
}

func (ms MetricsService) Register() {
	ws := new(restful.WebService)
	ws.
		Path("/metrics").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(ms.findmetric))
	ws.Route(ws.POST("/{metric}").To(ms.createmetric))
	restful.Add(ws)
}

func (ms MetricsService) findmetric(request *restful.Request, response *restful.Response) {
	fmt.Println("get Metrics request ")
	retString:=""
	if ms.MetricsCache!= nil {
		for metricType, value := range ms.MetricsCache{
			if value.MetricValue!=""{
				if ms.valideTime(ms.MetricsCache[metricType].Timestamp){
					retString+="# TYPE "+metricType+" gauge \n"+metricType+" "+string(value.MetricValue)+"\n"
					delete(ms.MetricsCache, metricType);
				}else{
					delete(ms.MetricsCache, metricType);
				}
			}
		}
		//response.WriteEntity("# TYPE cpu gauge \n cpu "+string(ms.MetricsCache["cpu"].Metric))
		response.WriteErrorString(http.StatusOK,retString)
	} else {
		response.WriteErrorString(http.StatusNotFound, "valid metric could not be found.")
	}
}

func (ms *MetricsService) createmetric(request *restful.Request, response *restful.Response) {
	metricType := request.PathParameter("metric")
	fmt.Println("post Metrics request type "+metricType)

	metric := new(Metric)
	err := request.ReadEntity(&metric)
	if err == nil {
		ms.MetricsCache[metricType]=*metric
		response.WriteHeaderAndEntity(http.StatusCreated, metric)
	} else {
		fmt.Println("cannot pares request body")
		fmt.Println(err)
		response.WriteError(http.StatusInternalServerError, err)
	}
}


