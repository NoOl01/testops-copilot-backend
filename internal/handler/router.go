package handler

import (
	"testops_copilot/internal/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "testops_copilot/docs"
)

func (h *handler) Router(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.POST("/generate", h.Generate)

		if config.Env.Debug {
			api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}
}
