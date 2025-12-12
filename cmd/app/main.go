package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testops_copilot/internal/config"
	"testops_copilot/internal/consts"
	"testops_copilot/internal/handler"
	"testops_copilot/internal/service"
	"testops_copilot/pkg/logger"
	"time"

	"github.com/NoOl01/golog/pkg/golog"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title TestOps Copilot
// @version 0.1.0
// @BasePath /api/v1
func main() {
	config.EnvLoad()
	logger.InitLogger()
	defer golog.Stop()

	logger.Log.Info(consts.Server, "server starting")

	srv := service.NewService()
	h := handler.NewHandler(srv)

	if !config.Env.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(cors.Default())
	h.Router(r)

	r.NoRoute(func(c *gin.Context) {
		c.File("./static/index.html")
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Env.ServerPort),
		Handler: r,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Log.Info(consts.Server, "server started on: "+config.Env.ServerPort)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Error(consts.Server, "server running error: "+err.Error())
		}
	}()

	<-quit
	logger.Log.Info(consts.Server, "server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Log.Error(consts.Server, "server forced to shutdown: "+err.Error())
	} else {
		logger.Log.Info(consts.Server, "server stopped successfully")
	}
}
