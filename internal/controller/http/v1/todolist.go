package v1

import (
	"database/sql"
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
			lists.Use(h.checkAccessRight)
			{
				lists.GET("/:list_id", h.getList)
				lists.PUT("/:list_id", h.updateList)
				lists.DELETE("/:list_id", h.deleteList)
			}
		}
		item := todo.Group("/list/:list_id", h.checkAccessRight)
		{
			item.GET("/items/:item_id", h.getItem)
			item.GET("/items/", h.getItems)
			item.POST("/items", h.createItems)
			item.PUT("/items/:item_id", h.updateItems)
			item.DELETE("/items/:item_id", h.deleteItems)
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
	id, err := strconv.ParseUint(c.Param("list_id"), 10, 64)
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
	id, err := strconv.ParseUint(c.Param("list_id"), 10, 64)
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
	id, err := strconv.ParseUint(c.Param("list_id"), 10, 64)
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

func (h *Handler) getItem(c *gin.Context) {
	itemId, err := strconv.ParseUint(c.Param("item_id"), 10, 64)
	if err != nil {
	}
	item, err := h.service.Item.GetById(itemId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, response.ItemResponse{
		ID:      item.ID,
		Title:   item.Title,
		Text:    item.Text,
		DueDate: item.DueDate.String,
		Checked: item.Checked,
	})
}

func (h *Handler) getItems(c *gin.Context) {
	listId, err := strconv.ParseUint(c.Param("list_id"), 10, 64)
	if err != nil {
	}
	items, err := h.service.Item.GetAllFromList(listId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	var result []*response.ItemResponse
	for _, item := range items {
		responseItem := &response.ItemResponse{
			ID:      item.ID,
			Title:   item.Title,
			Text:    item.Text,
			DueDate: item.DueDate.String,
			Checked: item.Checked,
		}
		result = append(result, responseItem)
	}
	c.JSON(http.StatusOK, result)
}

type InputItemRequest struct {
	Title   string         `form:"title" json:"title" binding:"required"`
	Text    string         `form:"text" json:"text" binding:"required"`
	DueDate sql.NullString `form:"due_date" json:"due_date"`
}

func (h *Handler) createItems(c *gin.Context) {
	var input InputItemRequest
	listId, err := strconv.ParseUint(c.Param("list_id"), 10, 64)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	item := &todo.Item{
		Title:   input.Title,
		Text:    input.Text,
		DueDate: input.DueDate,
	}
	itemId, err := h.service.Item.CreateItem(listId, item)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, map[string]uint64{
		"id": itemId,
	})
}

func (h *Handler) updateItems(c *gin.Context) {
	var input todo.UpdateItem
	itemId, err := strconv.ParseUint(c.Param("item_id"), 10, 64)
	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.service.Item.UpdateItem(input, itemId); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) deleteItems(c *gin.Context) {
	itemId, err := strconv.ParseUint(c.Param("item_id"), 10, 64)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.service.Item.DeleteItem(itemId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusNoContent)
}
