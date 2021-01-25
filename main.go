package main

import (
	"flag"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/server"
)

var (
	configFlag = flag.String("config", "./config/local.yaml", "path to configuration")
)

func main() {
	flag.Parse()
	server.Run(config.Init(*configFlag))
}
