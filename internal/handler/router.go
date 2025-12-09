package handler

import "github.com/gin-gonic/gin"

func (h *handler) Router(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.POST("/generate", h.Generate)
	}
}
