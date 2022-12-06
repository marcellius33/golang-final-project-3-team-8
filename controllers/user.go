package controllers

import (
	"github.com/dgrijalva/jwt-go"
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
	user := params.UserUpdateRequest{}
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}

	userData, _ := c.MustGet("userData").(jwt.MapClaims)
	updatedUser, err := u.service.Update(uint(userData["id"].(float64)), user)
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessResponse(updatedUser, "Update Success"))
}

func (u *UserController) UserDeleteController(c *gin.Context) {
	userData, _ := c.MustGet("userData").(jwt.MapClaims)
	err := u.service.Delete(uint(userData["id"].(float64)))
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.DeleteSuccess("Your account has been successfully deleted"))
}
