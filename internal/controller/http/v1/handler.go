package v1

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("v1")
	{
		v1.GET("test", func(context *gin.Context) {
			_, err := context.Writer.Write([]byte("Hello world"))
			if err != nil {
				return
			}
		})
	}
}
