package util
import (
	"os"
)


func BuildLabel() map[string]string{
	labels:=make(map[string]string)

    labels["pod_name"]=os.Getenv("HOSTNAME")
    labels["namespace_name"]="default"
    return labels
    // TODO get more labels
    //labels["app"]="kubia"
    //get labels["pod_name"] from k8s api
}