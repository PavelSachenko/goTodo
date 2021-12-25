package http

import (
	"github.com/gin-gonic/gin"
	"newExp/internal/config"
	v1 "newExp/internal/controller/http/v1"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	handler := gin.Default()
	h.InitApi(handler)
	return handler
}

func (h *Handler) InitApi(handler *gin.Engine) {
	routesV1 := v1.Handler{}
	api := handler.Group("/api")
	{
		routesV1.Init(api)
	}
}
