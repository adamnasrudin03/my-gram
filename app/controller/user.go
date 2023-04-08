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
// @Success 201 {object} dto.RegisterRes
// @Router /api/v1/auth/register [post]
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

	userRes, err := c.Service.User.Register(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	ctx.JSON(http.StatusCreated, userRes)
}

// Login godoc
// @Summary Login User
// @Description Login User
// @Tags Auth
// @Accept json
// @Produce json
// @Param dto.LoginReq body dto.LoginReq true "Login User"
// @Success 200 {object} dto.LoginRes
// @Router /api/v1/auth/login [post]
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

	loginRes, err := c.Service.User.Login(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.APIResponse(err.Error(), http.StatusInternalServerError, "error"))
		return
	}

	ctx.JSON(http.StatusOK, loginRes)
}
