package v1

import (
	"github.com/gin-gonic/gin"
	"newExp/internal/usecase"
)

type Handler struct {
	service *usecase.SuperService
}

func (h *Handler) Init(api *gin.RouterGroup, service *usecase.SuperService) {
	handler := api.Group("v1")
	{
		h.initTodoList(handler)

	}
}
