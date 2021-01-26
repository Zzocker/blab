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
	"github.com/Zzocker/blab/internal/health"
	"github.com/Zzocker/blab/internal/user"
	"github.com/Zzocker/blab/pkg/accesslog"
	"github.com/Zzocker/blab/pkg/log"
	"github.com/gin-gonic/gin"
)

// Run start server with give configuration
func Run(conf *config.C) {
	logger := log.New()

	router := gin.New() //
	// Add more configs to this router TODO

	addMiddlewares(router, logger)
	initializeRouters(router)

	start(router, logger, conf)
}

func initializeRouters(r *gin.Engine) {
	health.RegisterHandlers(r)
	user.RegisterHandlers(r)
}

func addMiddlewares(r *gin.Engine, l log.Logger) {
	r.Use(accesslog.Handler(l))
}

func start(r *gin.Engine, l log.Logger, conf *config.C) {
	address := fmt.Sprintf(":%d", conf.Port)
	srv := http.Server{
		Addr:    address,
		Handler: r,
	}
	go srv.ListenAndServe()
	l.Infof("Server started on port : %d", conf.Port)

	quit := make(chan os.Signal, 0)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	l.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		l.Warn("Server forced to shutdown")
	}
	l.Info("server exiting...")
}
