package http

import "github.com/gin-gonic/gin"

func (h *Handler) cors(c *gin.Context) {
	c.Header("Origin", "*")
}
