package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	userId = "userId"
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
