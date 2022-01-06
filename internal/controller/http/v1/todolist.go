package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"newExp/internal/controller/http/v1/response"
	"newExp/internal/model/todo"
	"strconv"
)

func (h *Handler) initTodoList(api *gin.RouterGroup) {
	todo := api.Group("/todo", h.auth, h.checkUserId)
	{
		lists := todo.Group("/lists")
		{
			lists.GET("/", h.getLists)
			lists.GET("/:id", h.getList)
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
	userId := getUserId(c)
	result, err := h.service.List.SearchLists(userId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) getList(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userId := getUserId(c)
	result, err := h.service.List.GetList(id, userId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}

type InputListRequest struct {
	Title       string `form:"title" json:"title" binding:"required"`
	Description string `form:"description" json:"description"`
}

func (h *Handler) createList(c *gin.Context) {
	var input InputListRequest
	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	list := &todo.List{
		Title:       input.Title,
		Description: input.Description,
	}
	userId := getUserId(c)
	result, err := h.service.List.CreateList(userId, list)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]uint64{
		"id": result,
	})
}

func (h *Handler) updateList(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userId := getUserId(c)
	var input todo.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	err = h.service.List.UpdateList(input, id, userId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) deleteList(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userId := getUserId(c)
	err = h.service.List.DeleteList(id, userId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.Status(http.StatusNoContent)
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
