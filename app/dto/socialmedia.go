package dto

import "adamnasrudin03/my-gram/app/entity"

type SocialMediaCreateReq struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
}

type SocialMediaListRes struct {
	ResponseList
	Data []entity.SocialMedia `json:"data"`
}

type SocialMediaUpdateReq struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
}
