package server

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/internal/health"
	"github.com/Zzocker/blab/pkg/log"
	"github.com/gorilla/mux"
)

// Run start server with give configuration
func Run(conf *config.C) {
	logger := log.New()

	router := mux.NewRouter() //
	// Add more configs to this router TODO

	initializeRouters(router)

	start(router, logger, conf)
}

func initializeRouters(r *mux.Router) {
	health.RegisterHandlers(r)
}

func start(r *mux.Router, l log.Logger, conf *config.C) {
	address := fmt.Sprintf(":%d", conf.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		l.Error(err)
		os.Exit(1)
	}
	go http.Serve(lis, r)
	l.Infof("Server started on port : %d", conf.Port)

	done := make(chan os.Signal, 0)
	signal.Notify(done, os.Interrupt)

	<-done
	l.Info("server interrupted")
}
