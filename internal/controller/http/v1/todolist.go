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
			lists.POST("/", h.createList)
			lists.Use(h.checkAccessRightForList)
			{
				lists.GET("/:list_id", h.getList)
				lists.PUT("/:list_id", h.updateList)
				lists.DELETE("/:list_id", h.deleteList)
			}
		}
		item := todo.Group("/list/:list_id", h.checkAccessRightForList)
		{
			item.GET("/items/:item_id", h.getItem)
			item.GET("/items/", h.getItems)
			item.POST("/items", h.createItems)
			item.PUT("/items/:item_id", h.updateItems)
			item.DELETE("/items/:item_id", h.deleteItems)
		}
	}
}

//---------Lists---------

func (h *Handler) getLists(c *gin.Context) {
	result, err := h.service.List.SearchLists(getUserId(c))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) getList(c *gin.Context) {
	result, err := h.service.List.GetList(getListId(c), getUserId(c))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) createList(c *gin.Context) {
	var input todo.InputListRequest
	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	userId := getUserId(c)
	result, err := h.service.List.CreateList(userId, input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]uint64{
		"id": result,
	})
}

func (h *Handler) updateList(c *gin.Context) {
	var input todo.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	err := h.service.List.UpdateList(input, getListId(c), getUserId(c))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) deleteList(c *gin.Context) {
	userId := getUserId(c)
	err := h.service.List.DeleteList(getListId(c), userId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.Status(http.StatusNoContent)
}

//---------Items---------

func (h *Handler) getItem(c *gin.Context) {
	itemId, err := strconv.ParseUint(c.Param("item_id"), 10, 64)
	if err != nil {
	}
	item, err := h.service.Item.GetById(itemId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) getItems(c *gin.Context) {
	items, err := h.service.Item.GetAllFromList(getListId(c))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) createItems(c *gin.Context) {
	var input todo.InputItemRequest
	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	itemId, err := h.service.Item.CreateItem(getListId(c), &input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusCreated, map[string]uint64{
		"id": itemId,
	})
}

func (h *Handler) updateItems(c *gin.Context) {
	var input todo.UpdateItem
	itemId, err := strconv.ParseUint(c.Param("item_id"), 10, 64)
	err = c.BindJSON(&input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	if err := h.service.Item.UpdateItem(input, itemId); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) deleteItems(c *gin.Context) {
	err := h.service.Item.DeleteItem(getListId(c))
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.Status(http.StatusNoContent)
}
