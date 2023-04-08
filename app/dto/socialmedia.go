package dto

import "adamnasrudin03/my-gram/app/entity"

type SocialMediaCreateReq struct {
	UserID         uint64 `json:"user_id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

type SocialMediaListRes struct {
	ResponseList
	Data []entity.SocialMedia `json:"data"`
}

type SocialMediaUpdateReq struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}
