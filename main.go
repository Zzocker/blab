package main

import "github.com/Zzocker/blab/internal/logger"

func main() {
	logger.L.Info("show me info message")
	logger.L.Error("show me error message")
}
