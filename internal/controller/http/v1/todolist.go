package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initTodoList(api *gin.RouterGroup) {
	todo := api.Group("/todo")
	{
		lists := todo.Group("/lists")
		{
			lists.GET("/", h.getLists)
			lists.POST("/", h.createList)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
		}
		item := todo.Group("/items")
		{
			item.GET("/", h.getItems)
			item.POST("/", h.createItems)
			item.PUT("/:id", h.updateItems)
			item.DELETE("/:id", h.deleteItems)
		}
	}
}

func (h *Handler) getLists(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *Handler) updateList(c *gin.Context) {
	c.Status(http.StatusOK)
}
func (h *Handler) createList(c *gin.Context) {
	c.Status(http.StatusOK)
}
func (h *Handler) deleteList(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *Handler) getItems(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *Handler) createItems(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *Handler) updateItems(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h *Handler) deleteItems(c *gin.Context) {
	c.Status(http.StatusOK)
}
