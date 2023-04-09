package dto

import "adamnasrudin03/my-gram/app/entity"

type PhotoCreateUpdateReq struct {
	Title    string `json:"title" validate:"required"`
	PhotoUrl string `json:"photo_url" validate:"required"`
	Caption  string `json:"caption"`
}

type PhotoListRes struct {
	ResponseList
	Data []entity.Photo `json:"data"`
}
