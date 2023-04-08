package service

import (
	"adamnasrudin03/my-gram/app/dto"
	"adamnasrudin03/my-gram/app/entity"
	"adamnasrudin03/my-gram/app/repository"
	"log"
)

type SocialMediaService interface {
	Create(input dto.SocialMediaCreateReq) (res entity.SocialMedia, err error)
}

type socialMediaSrv struct {
	SocialMediaRepository repository.SocialMediaRepository
}

func NewSocialMediaService(SocialMediaRepo repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaSrv{
		SocialMediaRepository: SocialMediaRepo,
	}
}

func (srv *socialMediaSrv) Create(input dto.SocialMediaCreateReq) (res entity.SocialMedia, err error) {
	socialMedia := entity.SocialMedia{
		UserID:         input.UserID,
		Name:           input.Name,
		SocialMediaUrl: input.SocialMediaUrl,
	}

	res, err = srv.SocialMediaRepository.Create(socialMedia)
	if err != nil {
		log.Printf("[SocialMediaService-Create] error create data: %+v \n", err)
		return
	}

	return
}
