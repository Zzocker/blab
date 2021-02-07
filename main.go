package main

import (
	"flag"
	"os"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/internal/logger"
)

var (
	configPath = flag.String("-conf", "config/local.yaml", "application configuration file")
)

func main() {
	flag.Parse()
	conf, err := config.LoadConfig(*configPath)
	if err != nil {
		os.Exit(1)
	}
	logger.Register(conf.LogLevel)
	logger.L.Infof(-1, "main", "application logger registered")
}
