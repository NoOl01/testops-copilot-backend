package handler

import (
	"testops_copilot/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Router(r *gin.Engine)
	Generate(ctx *gin.Context)
}

type handler struct {
	Service service.Service
}

func NewHandler(service service.Service) Handler {
	return &handler{
		Service: service,
	}
}
