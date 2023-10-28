package controller

import (
	"net/http"

	"github.com/Zeta201/identity-server/model"
	"github.com/Zeta201/identity-server/service"
	"github.com/Zeta201/identity-server/util"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func NewUserController(s service.UserService) UserController {
	return UserController{
		service: s,
	}
}

func (controller *UserController) SignUp(ctx *gin.Context) {
	var user model.User
	ctx.ShouldBindJSON(&user)

	if user.Username == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Username is required")
	}

	if user.Password == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Password is required")
	}

	err := controller.service.Save(user)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to save user")
		return
	}

	util.SuccessJSON(ctx, http.StatusCreated, "Signup successful")

}

func (controller *UserController) Login(ctx *gin.Context) {
	var user model.User
	ctx.ShouldBindJSON(&user)

	if user.Username == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Username is required")
	}

	if user.Password == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Password is required")
	}

	token, err := controller.service.LogIn(user)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	// response := foundUser.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "User found",
		Data:    &token,
	})
}
