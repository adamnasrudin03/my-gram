package controller

import (
	"adamnasrudin03/my-gram/app/dto"
	"adamnasrudin03/my-gram/app/entity"
	"adamnasrudin03/my-gram/app/service"
	"adamnasrudin03/my-gram/pkg/helpers"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type PhotoController interface {
	CreatePhoto(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
}

type photoHandler struct {
	Service *service.Services
}

func NewPhotoController(srv *service.Services) PhotoController {
	return &photoHandler{
		Service: srv,
	}
}

// CreatePhoto godoc
// @Summary CreatePhoto
// @Description Create new data Photo
// @Tags Photo
// @Accept json
// @Produce json
// @Param dto.PhotoCreateUpdateReq body dto.PhotoCreateUpdateReq true "Create Photo"
// @Success 201 {object} entity.Photo
// @Router /photos [post]
// @Security BearerAuth
func (c *photoHandler) CreatePhoto(ctx *gin.Context) {
	var (
		input dto.PhotoCreateUpdateReq
	)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	validate := validator.New()

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	err = validate.Struct(input)
	if err != nil {
		errors := helpers.FormatValidationError(err)

		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(errors, http.StatusBadRequest, "error"))
		return
	}

	Photo := entity.Photo{
		UserID:   uint64(userData["id"].(float64)),
		Title:    input.Title,
		PhotoUrl: input.PhotoUrl,
		Caption:  input.Caption,
	}

	PhotoRes, httpStatus, err := c.Service.Photo.Create(Photo)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, PhotoRes)
}

// GetAll godoc
// @Summary GetAll
// @Description Get All Photo
// @Tags Photo
// @Accept json
// @Produce json
// @Param page query uint64 true "Pagination Get All Photo"
// @Param limit query uint64 true "Pagination Get All Photo"
// @Success 200 {object} dto.PhotoListRes
// @Router /photos [get]
// @Security BearerAuth
func (c *photoHandler) GetAll(ctx *gin.Context) {
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

	res, httpStatus, err := c.Service.Photo.GetAll(ctx, param)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, res)
}

// GetOne godoc
// @Summary GetOne
// @Description GetOne Photo by ID
// @Tags Photo
// @Accept json
// @Produce json
// @Param id path uint64 true "Photo ID"
// @Success 200 {object} entity.Photo
// @Router /photos/{id} [get]
// @Security BearerAuth
func (c *photoHandler) GetOne(ctx *gin.Context) {
	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	res, httpStatus, err := c.Service.Photo.GetByID(ID)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}
	ctx.JSON(httpStatus, res)
}

// UpdatePhoto godoc
// @Summary UpdatePhoto
// @Description Update Photo by ID
// @Tags Photo
// @Accept json
// @Produce json
// @Param dto.PhotoCreateUpdateReq body dto.PhotoCreateUpdateReq true "Update Photo"
// @Param id path uint64 true "Photo ID"
// @Success 200 {object} entity.Photo
// @Router /photos/{id} [put]
// @Security BearerAuth
func (c *photoHandler) UpdatePhoto(ctx *gin.Context) {
	var (
		input dto.PhotoCreateUpdateReq
	)

	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	validate := validator.New()
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	err = validate.Struct(input)
	if err != nil {
		errors := helpers.FormatValidationError(err)

		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(errors, http.StatusBadRequest, "error"))
		return
	}

	PhotoRes, httpStatus, err := c.Service.Photo.UpdateByID(ID, input)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, PhotoRes)
}

// DeletePhoto godoc
// @Summary DeletePhoto
// @Description Delete Photo by ID
// @Tags Photo
// @Accept json
// @Produce json
// @Param id path uint64 true "Photo ID"
// @Success 200 {object} helpers.ResponseDefault
// @Router /photos/{id} [delete]
// @Security BearerAuth
func (c *photoHandler) DeletePhoto(ctx *gin.Context) {
	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	httpStatus, err := c.Service.Photo.DeleteByID(ID)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, helpers.APIResponse("deleted", httpStatus, "success"))
}
