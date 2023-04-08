package service

import (
	"adamnasrudin03/my-gram/app/dto"
	"adamnasrudin03/my-gram/app/entity"
	"adamnasrudin03/my-gram/app/repository"
	"log"
	"math"

	"github.com/gin-gonic/gin"
)

type SocialMediaService interface {
	Create(input dto.SocialMediaCreateReq) (res entity.SocialMedia, err error)
	GetAll(ctx *gin.Context, queryparam dto.ListParam) (result dto.SocialMediaListRes, err error)
	GetByID(ID uint64) (result entity.SocialMedia, err error)
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

func (srv *socialMediaSrv) GetAll(ctx *gin.Context, queryparam dto.ListParam) (result dto.SocialMediaListRes, err error) {
	result.Limit = queryparam.Limit
	result.Page = queryparam.Page

	result.Data, result.Total, err = srv.SocialMediaRepository.GetAll(ctx, queryparam)
	if err != nil {
		log.Printf("[SocialMediaService-GetAll] error get data repo: %+v \n", err)
		return result, err
	}

	result.LastPage = uint64(math.Ceil(float64(result.Total) / float64(queryparam.Limit)))

	return result, nil
}

func (srv *socialMediaSrv) GetByID(ID uint64) (result entity.SocialMedia, err error) {
	result, err = srv.SocialMediaRepository.GetByID(ID)
	if err != nil {
		log.Printf("[SocialMediaService-GetByID] error get data repo: %+v \n", err)
		return result, err
	}

	return result, nil
}
