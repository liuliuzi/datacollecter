package options

import (
	"github.com/spf13/pflag"
)

type DatacollecterRunOptions struct {
	Mode         int
	LocalIP      string
	LocalPort    string
	LiveTime     string
	InfluxdbIP   string
	InfluxdbPort string
	Version      bool
}

func NewDatacollecterRunOptions() *DatacollecterRunOptions {
	return &DatacollecterRunOptions{
	}
}

func (d *DatacollecterRunOptions) AddFlags(fs *pflag.FlagSet) {
	fs.IntVar(&d.Mode,"mode", 0, "0:default mode, heapster get metrics value from this app; 1 : this app send metrics value to influxdb dircectly")
	fs.StringVar(&d.LocalIP,"localIP", "0.0.0.0", "local IP")
	fs.StringVar(&d.LocalPort,"localPort", "8060", "local Port opened")
	fs.StringVar(&d.LiveTime,"liveTime", "60s", "metric sample live time")
	fs.StringVar(&d.InfluxdbIP,"influxdbIP", "10.140.163.102", "influxdb IP effective in mode 1")
	fs.StringVar(&d.InfluxdbPort,"influxdbPort", "8086", "influxdb Port effective in mode 1")
	fs.BoolVar(&d.Version, "version", false, "print version info and exit")
}