package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"newExp/internal/config"
	v1 "newExp/internal/controller/http/v1"
	"newExp/internal/usecase"

	_ "newExp/docs"
)

type Handler struct {
	service *usecase.SuperService
}

func NewHandler(service *usecase.SuperService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	handler := gin.New()
	h.InitApi(handler)
	return handler
}

func (h *Handler) InitApi(handler *gin.Engine) {
	routesV1 := v1.Handler{}
	handler.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := handler.Group("/api", h.cors)
	{
		routesV1.Init(api, h.service)
	}
}
