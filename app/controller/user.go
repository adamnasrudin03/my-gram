package controller

import (
	"net/http"

	"github.com/adamnasrudin03/my-gram/app/dto"
	"github.com/adamnasrudin03/my-gram/app/entity"
	"github.com/adamnasrudin03/my-gram/app/service"
	"github.com/adamnasrudin03/my-gram/pkg/helpers"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type userController struct {
	Service *service.Services
}

func NewUserController(srv *service.Services) UserController {
	return &userController{
		Service: srv,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var input entity.User

	//Validation input user
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)

		response := helpers.APIResponse(err.Error(), http.StatusUnprocessableEntity, errors, nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = c.Service.User.Register(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse("there was an error register new user", http.StatusBadRequest, err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.APIResponse("user registered", http.StatusCreated, "success", nil))
}

func (c *userController) Login(ctx *gin.Context) {
	var (
		input    dto.LoginReq
		response dto.LoginRes
	)

	//Validation input user
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error", nil))
		return
	}

	response.Token, err = c.Service.User.Login(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.APIResponse(err.Error(), http.StatusInternalServerError, "error", nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.APIResponse("login success", http.StatusOK, "success", response))
}
