package controller

import (
	"net/http"

	"github.com/adamnasrudin03/my-gram/app/dto"
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

// Register godoc
// @Summary Register User
// @Description Register new User
// @Tags json
// @Accept json
// @Produce json
// @Param dto.RegisterReq body dto.RegisterReq true "Register User"
// @Success 201 {object} entity.User
// @Router /api/v1/auth/register [post]
func (c *userController) Register(ctx *gin.Context) {
	var input dto.RegisterReq

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)

		response := helpers.APIResponse(err.Error(), http.StatusUnprocessableEntity, errors, nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response, err := c.Service.User.Register(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse("there was an error register new user", http.StatusBadRequest, err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.APIResponse("user registered", http.StatusCreated, "success", response))
}

// Login godoc
// @Summary Login User
// @Description Login User
// @Tags json
// @Accept json
// @Produce json
// @Param dto.LoginReq body dto.LoginReq true "Login User"
// @Success 200 {object} dto.LoginRes
// @Router /api/v1/auth/login [post]
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
