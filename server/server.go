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
	"github.com/Zzocker/blab/core"
	stub "github.com/Zzocker/blab/internal/http"
	"github.com/Zzocker/blab/internal/logger"
	"github.com/Zzocker/blab/internal/middleware"
	"github.com/gin-gonic/gin"
)

const (
	pkgName = "server"
)

// BuildAndRun :
func BuildAndRun(conf *config.ApplicationConf) {
	logger.L.Infof(-1, pkgName, "building server")

	engin := gin.New()
	engin.Use(gin.Recovery())
	engin.Use(middleware.Access())

	public := engin.Group("/public")
	private := engin.Group("/api")

	// register routers
	core.BuildAll(conf)
	stub.RegisterRouters(conf, public, private)
	start(engin, conf.Port)
}

func start(engine *gin.Engine, port int) {
	logger.L.Infof(-1, pkgName, "starting http server on port=%d", port)
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: engine,
	}
	logger.L.Infof(-1, pkgName, "http server started on port=%d", port)
	go srv.ListenAndServe()

	quit := make(chan os.Signal, 0)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.L.Errorf(-2, pkgName, "forceing server to shutdown")
	}
	logger.L.Infof(-2, pkgName, "existing server...")
}
