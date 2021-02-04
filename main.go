package main

import (
	"flag"
	"os"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/internal/logger"
)

var (
	confPath = flag.String("conf", "config/local.yaml", "configuration path")
)

func main() {
	flag.Parse()
	conf, err := config.ReadConf(*confPath)
	if err != nil {
		logger.L.Error(err.Error())
		os.Exit(1)
	}
	_ = conf
}
