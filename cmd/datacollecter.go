package main
import (
	"os"
	"fmt"
	"github.com/spf13/pflag"
	goflag "flag"
	"github.com/liuliuzi/datacollecter/cmd/app"
	"github.com/liuliuzi/datacollecter/cmd/app/options"
	"github.com/liuliuzi/datacollecter/version"
)

func main() {
	opt := options.NewDatacollecterRunOptions()
	opt.AddFlags(pflag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	pflag.Parse()

	if opt.Version {
		fmt.Println(version.VersionInfo())
		os.Exit(0)
	}
	if err := app.Run(opt); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
