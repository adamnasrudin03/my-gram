package controller

import (
	"net/http"

	"adamnasrudin03/my-gram/app/dto"
	"adamnasrudin03/my-gram/app/service"
	"adamnasrudin03/my-gram/pkg/helpers"

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

// Register godoc
// @Summary Register User
// @Description Register new User
// @Tags Auth
// @Accept json
// @Produce json
// @Param dto.RegisterReq body dto.RegisterReq true "Register User"
// @Success 201 {object} entity.User
// @Router /auth/register [post]
func (c *userController) Register(ctx *gin.Context) {
	var (
		input dto.RegisterReq
	)

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)

		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(errors, http.StatusBadRequest, "error"))
		return
	}

	userRes, statusHttp, err := c.Service.User.Register(input)
	if err != nil {
		ctx.JSON(statusHttp, helpers.APIResponse(err.Error(), statusHttp, "error"))
		return
	}

	ctx.JSON(statusHttp, userRes)
}

// Login godoc
// @Summary Login User
// @Description Login User
// @Tags Auth
// @Accept json
// @Produce json
// @Param dto.LoginReq body dto.LoginReq true "Login User"
// @Success 200 {object} dto.LoginRes
// @Router /auth/login [post]
func (c *userController) Login(ctx *gin.Context) {
	var (
		input dto.LoginReq
	)

	//Validation input user
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	loginRes, statusHttp, err := c.Service.User.Login(input)
	if err != nil {
		ctx.JSON(statusHttp, helpers.APIResponse(err.Error(), statusHttp, "error"))
		return
	}

	ctx.JSON(statusHttp, loginRes)
}
