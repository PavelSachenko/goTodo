package v1

import (
	"errors"
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

func getUserId(c *gin.Context) (uint64, error) {
	id, ok := c.Get(userId)
	if !ok {
		return 0, errors.New("user not found")
	}

	uId, ok := id.(uint64)
	if !ok {
		return 0, errors.New("id is of invalid type")
	}

	return uId, nil
}
