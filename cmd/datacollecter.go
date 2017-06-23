package main
import (
	"flag"
	"os"
	"fmt"
	"github.com/liuliuzi/datacollecter/cmd/app"
)

func main() {
	mode         := flag.Int("mode", 0, "0:default mode, heapster get metrics value from this app; 1 : this app send metrics value to influxdb dircectly")
	localIP      := flag.String("localIP", "0.0.0.0", "local IP")
	localPort    := flag.String("localPort", "8060", "local Port opened")
	liveTime     := flag.String("liveTime", "60s", "metric sample live time")
	influxdbIP   := flag.String("influxdbIP", "10.140.163.102", "influxdb IP effective in mode 1")
	influxdbPort := flag.String("influxdbPort", "8086", "influxdb Port effective in mode 1")
	flag.Parse()

	if err := app.Run(*mode,*localIP,*localPort,*liveTime,*influxdbIP,*influxdbPort); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
