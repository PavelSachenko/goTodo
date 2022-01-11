package v1

import (
	"github.com/gin-gonic/gin"
	"newExp/internal/usecase"
)

type Handler struct {
	service *usecase.SuperService
}

func (h *Handler) Init(api *gin.RouterGroup, service *usecase.SuperService) {
	h.service = service
	handler := api.Group("v1")
	{
		h.initTodoList(handler)
		h.initUser(handler)
	}
}

func getUserId(c *gin.Context) uint64 {
	id, _ := c.Get(userId)
	return id.(uint64)
}

func getListId(c *gin.Context) uint64 {
	id, _ := c.Get(listId)
	return id.(uint64)
}
