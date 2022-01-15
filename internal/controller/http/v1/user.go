package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	_ "newExp/docs"
	"newExp/internal/controller/http/v1/response"
	_ "newExp/internal/model"
)

func (h *Handler) initUser(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		signIn := user.Group("sign-in/")
		{
			signIn.POST("/", h.signIn)
		}
		//signUp := user.Group("sign-up/")
		//{
		//	signUp.POST("/", h.signUp)
		//}
		user.Use(h.auth).GET("/", h.getUser)
	}
}

// @Summary User
// @Security ApiKeyAuth
// @Tags user
// @Description get user
// @ID get-user
// @Produce json
// @Success 200 {object} model.User
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Router /api/v1/user [get]
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

// @Summary SignIn
// @Tags auth
// @Description login into system
// @ID login
// @Accept json
// @Produce json
// @Param input body signInRequest true "list info"
// @Success 200 {object} map[string]string
// @Failure 400 {object} response.ErrorResponse
// @Failure 403 {object} response.ErrorResponse
// @Router /api/v1/user/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInRequest
	if err := c.BindWith(&input, binding.JSON); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	bearer, err := h.service.Auth.SignIn(input.Username, input.Password)
	if err != nil {
		response.NewErrorResponse(c, http.StatusForbidden, err.Error())
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

// @Summary SignUp
// @Tags auth
// @Description create user
// @ID create-user
// @Accept json
// @Produce json
// @Param input body signUpRequest true "list info"
// @Success 201 {object} map[string]string
// @Failure 400 {object} response.ErrorResponse
// @Router /api/v1/user/sign-up [post]
//func (h *Handler) signUp(c *gin.Context) {
//	var input signUpRequest
//	if err := c.BindWith(&input, binding.JSON); err != nil {
//		response.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
//		return
//	}
//	bearer, err := h.service.Auth.CreateUser(&model.User{Username: input.Username, Password: input.Password})
//	if err != nil {
//		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
//		return
//	}
//
//	c.JSON(http.StatusCreated, map[string]string{
//		"bearerToken": bearer,
//	})
//}
