package controller

import (
	"adamnasrudin03/my-gram/app/dto"
	"adamnasrudin03/my-gram/app/service"
	"adamnasrudin03/my-gram/pkg/helpers"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SocialMediaController interface {
	CreateSocialMedia(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	UpdateSocialMedia(ctx *gin.Context)
	DeleteSocialMedia(ctx *gin.Context)
}

type socialMediaHandler struct {
	Service *service.Services
}

func NewSocialMediaController(srv *service.Services) SocialMediaController {
	return &socialMediaHandler{
		Service: srv,
	}
}

// CreateSocialMedia godoc
// @Summary CreateSocialMedia
// @Description Create new SocialMedia
// @Tags Social Media
// @Accept json
// @Produce json
// @Param dto.SocialMediaCreateReq body dto.SocialMediaCreateReq true "Create SocialMedia"
// @Success 201 {object} entity.SocialMedia
// @Router /api/v1/social-media [post]
func (c *socialMediaHandler) CreateSocialMedia(ctx *gin.Context) {
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
// @Param page query uint64 false "Pagination Get All Social Media"
// @Param limit query uint64 false "Pagination Get All Social Media"
// @Success 201 {object} dto.SocialMediaListRes
// @Router /api/v1/social-media [get]
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

// GetOne godoc
// @Summary GetOne SocialMedia
// @Description GetOne SocialMedia
// @Tags Social Media
// @Accept json
// @Produce json
// @Param id path uint64 true "Social Media ID"
// @Success 201 {object} entity.SocialMedia
// @Router /api/v1/social-media/{{id}} [get]
func (c *socialMediaHandler) GetOne(ctx *gin.Context) {
	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	res, err := c.Service.SocialMedia.GetByID(ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, helpers.APIResponse(err.Error(), http.StatusNotFound, "error"))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.APIResponse(err.Error(), http.StatusInternalServerError, "error"))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// UpdateSocialMedia godoc
// @Summary UpdateSocialMedia
// @Description Update  SocialMedia
// @Tags Social Media
// @Accept json
// @Produce json
// @Param dto.SocialMediaUpdateReq body dto.SocialMediaUpdateReq true "Update SocialMedia"
// @Param id path uint64 true "Social Media ID"
// @Success 200 {object} entity.SocialMedia
// @Router /api/v1/social-media/{id} [put]
func (c *socialMediaHandler) UpdateSocialMedia(ctx *gin.Context) {
	var (
		input dto.SocialMediaUpdateReq
	)

	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	SocialMediaRes, httpStatus, err := c.Service.SocialMedia.UpdateByID(ID, input)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, SocialMediaRes)
}

// DeleteSocialMedia godoc
// @Summary DeleteSocialMedia
// @Description Delete SocialMedia by ID
// @Tags Social Media
// @Accept json
// @Produce json
// @Param id path uint64 true "Social Media ID"
// @Success 200 {object} helpers.ResponseDefault
// @Router /api/v1/social-media/{id} [delete]
func (c *socialMediaHandler) DeleteSocialMedia(ctx *gin.Context) {
	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	httpStatus, err := c.Service.SocialMedia.DeleteByID(ID)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, helpers.APIResponse("deleted", httpStatus, "success"))
}
