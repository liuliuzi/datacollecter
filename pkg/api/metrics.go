package api
import (
	"net/http"
	"strconv"
	"fmt"
	"time"
	"github.com/emicklei/go-restful"
	"k8s.io/heapster/metrics/core"
	"github.com/liuliuzi/datacollecter/pkg/util"
)

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
		if ms.Mode==0 {
			ms.MetricsCache[metricType]=*metric
			response.WriteHeaderAndEntity(http.StatusCreated, metric)
		}else{
			ms.Pushinfluxdb(metricType,*metric)
			response.WriteHeaderAndEntity(http.StatusCreated, metric)
		}

	} else {
		fmt.Println("cannot pares request body")
		fmt.Println(err)
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func (ms *MetricsService)Pushinfluxdb(metricType string,metric Metric) {
	fmt.Println("Push data to influxdb")
	metricValues:=make(map[string]core.MetricValue)
	MetricValueF64,_:=strconv.ParseFloat(metric.MetricValue, 32)

	metricValues["custom/"+metricType]=core.MetricValue{0,float32(MetricValueF64),core.MetricGauge,core.ValueFloat}
	labels:=util.BuildLabel()

	MetricSet:=core.MetricSet{time.Now(),metric.Timestamp,metricValues,labels,nil}
	MetricSets:=make(map[string]*core.MetricSet)
	MetricSets["1"]=&MetricSet
	databatch:=core.DataBatch{time.Now(), MetricSets}

	ms.InfluxdbSink.ExportData(&databatch)

}




