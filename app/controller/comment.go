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

type CommentController interface {
	CreateComment(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}

type commentHandler struct {
	Service *service.Services
}

func NewCommentController(srv *service.Services) CommentController {
	return &commentHandler{
		Service: srv,
	}
}

// CreateComment godoc
// @Summary CreateComment
// @Description Create new data Comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param dto.CommentCreateUpdateReq body dto.CommentCreateUpdateReq true "Create Comment"
// @Success 201 {object} dto.CommentCreateUpdateResponse
// @Router /comments [post]
// @Security BearerAuth
func (c *commentHandler) CreateComment(ctx *gin.Context) {
	var (
		input dto.CommentCreateUpdateReq
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

	Comment := entity.Comment{
		UserID:  uint64(userData["id"].(float64)),
		PhotoID: input.PhotoID,
		Message: input.Message,
	}

	CommentRes, httpStatus, err := c.Service.Comment.Create(Comment)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, CommentRes)
}

// GetAll godoc
// @Summary GetAll
// @Description Get All Comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param page query uint64 true "Pagination Get All Comment"
// @Param limit query uint64 true "Pagination Get All Comment"
// @Success 200 {object} dto.CommentListRes
// @Router /comments [get]
// @Security BearerAuth
func (c *commentHandler) GetAll(ctx *gin.Context) {
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

	res, httpStatus, err := c.Service.Comment.GetAll(ctx, param)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, res)
}

// GetOne godoc
// @Summary GetOne
// @Description GetOne Comment by ID
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path uint64 true "Comment ID"
// @Success 200 {object} entity.Comment
// @Router /comments/{id} [get]
// @Security BearerAuth
func (c *commentHandler) GetOne(ctx *gin.Context) {
	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	res, httpStatus, err := c.Service.Comment.GetByID(ID)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}
	ctx.JSON(httpStatus, res)
}

// UpdateComment godoc
// @Summary UpdateComment
// @Description Update Comment by ID
// @Tags Comment
// @Accept json
// @Produce json
// @Param dto.CommentCreateUpdateReq body dto.CommentCreateUpdateReq true "Update Comment"
// @Param id path uint64 true "Comment ID"
// @Success 200 {object} dto.CommentCreateUpdateResponse
// @Router /comments/{id} [put]
// @Security BearerAuth
func (c *commentHandler) UpdateComment(ctx *gin.Context) {
	var (
		input dto.CommentCreateUpdateReq
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

	CommentRes, httpStatus, err := c.Service.Comment.UpdateByID(ID, input)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, CommentRes)
}

// DeleteComment godoc
// @Summary DeleteComment
// @Description Delete Comment by ID
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path uint64 true "Comment ID"
// @Success 200 {object} helpers.ResponseDefault
// @Router /comments/{id} [delete]
// @Security BearerAuth
func (c *commentHandler) DeleteComment(ctx *gin.Context) {
	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), http.StatusBadRequest, "error"))
		return
	}

	httpStatus, err := c.Service.Comment.DeleteByID(ID)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, "error"))
		return
	}

	ctx.JSON(httpStatus, helpers.APIResponse("Deleted", httpStatus, "false"))
}
