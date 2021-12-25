package v1

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) Init(api *gin.RouterGroup) {
	handler := api.Group("v1")
	{
		h.initTodoList(handler)

	}
}
