package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/internal/httpapi"
	"github.com/Zzocker/blab/internal/middleware"
	"github.com/Zzocker/blab/pkg/log"
	"github.com/gin-gonic/gin"
)

// Run start server with give configuration
func Run(conf *config.C) {
	r := gin.New()

	r.Use(gin.Recovery())
	middleware.BuildMiddleware(*conf)

	auth := r.Group("/a", middleware.OAuth())
	noAuth := r.Group("/n")

	httpapi.BuildAllRouter(*conf, noAuth, auth)
	start(r, conf)
}

func start(r *gin.Engine, conf *config.C) {
	address := fmt.Sprintf(":%d", conf.Port)
	srv := http.Server{
		Addr:    address,
		Handler: r,
	}
	go srv.ListenAndServe()
	log.L.Infof("Server started on port : %d", conf.Port)

	quit := make(chan os.Signal, 0)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.L.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.L.Warn("Server forced to shutdown")
	}
	log.L.Info("server exiting...")
}
