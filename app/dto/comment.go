package dto

import "adamnasrudin03/my-gram/app/entity"

type CommentCreateUpdateReq struct {
	PhotoID uint64 `json:"photo_id"`
	Message string `json:"message" validate:"required"`
}

type CommentListRes struct {
	ResponseList
	Data []entity.Comment `json:"data"`
}
