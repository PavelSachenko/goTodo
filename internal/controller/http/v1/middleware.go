package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

var (
	userId = "userId"
	listId = "listId"
)

func (h *Handler) auth(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
			"Authorization": "not found Authorization header",
		})
		return
	}
	token := strings.Split(header, "Bearer ")
	if token[1] == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
			"Authorization": "token is empty",
		})
		return
	}

	id, err := h.service.Auth.ParseToke(token[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
			"Authorization": "wrong token",
		})
		return
	}

	c.Set(userId, id)
}

func (h *Handler) checkUserId(c *gin.Context) {
	id, ok := c.Get(userId)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
			"Authorization": "user not found",
		})
	}

	_, ok = id.(uint64)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
			"error": "id is of invalid type",
		})
	}
}

func (h *Handler) checkAccessRightForList(c *gin.Context) {
	list_id, err := strconv.ParseUint(c.Param("list_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"error": "list_id is of invalid type",
		})
	}
	c.Set(listId, list_id)
	result, err := h.service.List.CheckAccessRight(list_id, getUserId(c))
	if err != nil || result != true {
		c.AbortWithStatusJSON(http.StatusForbidden, map[string]string{
			"error": "Access is denied",
		})
	}
}
