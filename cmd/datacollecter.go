package main
import (
	"flag"
	"os"
	"fmt"
	"github.com/liuliuzi/datacollecter/cmd/app"
)

func main() {
	localIP   := flag.String("localIP", "0.0.0.0", "local IP")
	localPort := flag.String("localPort", "8060", "local Port opened")
	liveTime  := flag.String("liveTime", "60s", "metric sample live time")

	flag.Parse()

	if err := app.Run(*localIP,*localPort,*liveTime); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
