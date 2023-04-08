package controller

import (
	"adamnasrudin03/my-gram/app/dto"
	"adamnasrudin03/my-gram/app/service"
	"adamnasrudin03/my-gram/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SocialMediaController interface {
	Create(ctx *gin.Context)
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
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	ctx.JSON(http.StatusCreated, SocialMediaRes)
}
