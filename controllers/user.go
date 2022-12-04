package controllers

import (
	"github.com/gin-gonic/gin"
	"kanbanboard/helpers"
	"kanbanboard/params"
	"kanbanboard/services"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (u *UserController) UserRegisterController(c *gin.Context) {
	userRequest := params.UserRegisterRequest{}
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}

	createUser, err := u.service.Register(userRequest)
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessCreateResponse(createUser, "Register Success"))
}

func (u *UserController) UserLoginController(c *gin.Context) {
	loginRequest := params.UserLoginRequest{}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}

	token, err := u.service.Login(loginRequest)
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessResponse(token, "Login Success"))
}

func (u *UserController) UserUpdateController(c *gin.Context) {

}

func (u *UserController) UserDeleteController(c *gin.Context) {

}
