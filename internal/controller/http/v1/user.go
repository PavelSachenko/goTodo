package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"newExp/internal/controller/http/v1/response"
	"newExp/internal/model"
)

func (h *Handler) initUser(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		signIn := user.Group("sign-in/")
		{
			signIn.POST("/", h.signIn)
		}
		signUp := user.Group("sign-up/")
		{
			signUp.POST("/", h.signUp)
		}
		user.Use(h.auth).GET("/", h.getUser)
	}
}

func (h *Handler) getUser(c *gin.Context) {
	id := getUserId(c)
	user, err := h.service.Auth.GetUser(id)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, user)
}

type signInRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInRequest
	if err := c.BindWith(&input, binding.FormMultipart); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	bearer, err := h.service.Auth.SignIn(input.Username, input.Password)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"bearerToken": bearer,
	})
}

type signUpRequest struct {
	Username string `form:"username"  json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input signUpRequest
	if err := c.BindWith(&input, binding.FormMultipart); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	bearer, err := h.service.Auth.CreateUser(&model.User{Username: input.Username, Password: input.Password})
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"bearerToken": bearer,
	})
}
