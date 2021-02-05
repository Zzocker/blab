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
	api "github.com/Zzocker/blab/internal/http"
	"github.com/Zzocker/blab/internal/logger"
	"github.com/Zzocker/blab/internal/middleware"
	"github.com/gin-gonic/gin"
)

const (
	serverLoggerPrefix = "[server] %v"
)

// BuildAndRun : builds and run server with given configuration
func BuildAndRun(conf config.ApplicationConf) {
	logger.L.Info(serverLoggerPrefix, "building server")

	// gin engine
	engine := gin.New()
	engine.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)
	engine.Use(middleware.AccessLog())
	// oauth and non-oauth group
	oauth := engine.Group("/")
	noOauth := engine.Group("/public")

	// register routers
	if err := api.RegisterRouters(conf, oauth, noOauth); err != nil {
		logger.L.Error(serverLoggerPrefix, fmt.Sprintf("failed to register routers : %v", err.Error()))
		os.Exit(1)
	}

	// start server
	startServer(engine, conf.Port)
}

func startServer(engine *gin.Engine, port int) {
	logger.L.Info(serverLoggerPrefix, fmt.Sprintf("starting http server on port=%d", port))
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: engine,
	}
	logger.L.Info(serverLoggerPrefix, fmt.Sprintf("http server started on port=%d", port))
	go srv.ListenAndServe()

	quit := make(chan os.Signal, 0)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.L.Error(serverLoggerPrefix, "forceing server to shutdown")
	}
	logger.L.Info(serverLoggerPrefix, "existing server...")
}
