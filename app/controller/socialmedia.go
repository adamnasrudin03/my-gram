package controller

import (
	"adamnasrudin03/my-gram/app/dto"
	"adamnasrudin03/my-gram/app/service"
	"adamnasrudin03/my-gram/pkg/helpers"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SocialMediaController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

type socialMediaHandler struct {
	Service *service.Services
}

func NewSocialMediaController(srv *service.Services) SocialMediaController {
	return &socialMediaHandler{
		Service: srv,
	}
}

// Create godoc
// @Summary Create SocialMedia
// @Description Create new SocialMedia
// @Tags Social Media
// @Accept json
// @Produce json
// @Param dto.SocialMediaCreateReq body dto.SocialMediaCreateReq true "Create SocialMedia"
// @Success 201 {object} entity.SocialMedia
// @Router /api/v1/social-media [post]
func (c *socialMediaHandler) Create(ctx *gin.Context) {
	var (
		input dto.SocialMediaCreateReq
	)

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)

		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(errors, http.StatusBadRequest, "error"))
		return
	}

	SocialMediaRes, err := c.Service.SocialMedia.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.APIResponse(err.Error(), http.StatusInternalServerError, "error"))
		return
	}

	ctx.JSON(http.StatusCreated, SocialMediaRes)
}

// Get All godoc
// @Summary Get All SocialMedia
// @Description Get All new SocialMedia
// @Tags Social Media
// @Accept json
// @Produce json
// @Param dto.ListParam body dto.ListParam true "Get All Social Media"
// @Success 201 {object} dto.SocialMediaListRes
// @Router /api/v1/social-media [GET]
func (c *socialMediaHandler) GetAll(ctx *gin.Context) {
	var (
		paramPage  uint64 = 1
		paramLimit uint64 = 10
		err        error
	)

	if ctx.Query("page") == "" {
		paramPage, err = strconv.ParseUint(ctx.Query("page"), 10, 32)
		if err != nil {
			err = errors.New("query param page invalid")
			ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
			return
		}
	}

	if ctx.Query("limit") != "" {
		paramLimit, err = strconv.ParseUint(ctx.Query("limit"), 10, 32)
		if err != nil {
			err = errors.New("query param limit invalid")
			ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
			return
		}
	}

	param := dto.ListParam{
		Page:  paramPage,
		Limit: paramLimit,
	}

	res, err := c.Service.SocialMedia.GetAll(ctx, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.APIResponse(err.Error(), http.StatusInternalServerError, "error"))
		return
	}

	ctx.JSON(http.StatusOK, res)
}
