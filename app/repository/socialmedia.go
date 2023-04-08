package repository

import (
	"log"

	"adamnasrudin03/my-gram/app/dto"
	"adamnasrudin03/my-gram/app/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SocialMediaRepository interface {
	Create(input entity.SocialMedia) (res entity.SocialMedia, err error)
	GetAll(ctx *gin.Context, queryparam dto.ListParam) (result []entity.SocialMedia, total uint64, err error)
}

type socialMediaRepo struct {
	DB *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepo{
		DB: db,
	}
}

func (repo *socialMediaRepo) Create(input entity.SocialMedia) (res entity.SocialMedia, err error) {
	if err = repo.DB.Create(&input).Error; err != nil {
		log.Printf("[SocialMediaRepository-Create] error Create new data: %+v \n", err)
		return input, err
	}

	return input, err
}

func (repo *socialMediaRepo) GetAll(ctx *gin.Context, queryparam dto.ListParam) (result []entity.SocialMedia, total uint64, err error) {
	offset := queryparam.Limit * (queryparam.Page - 1)
	var totaldata int64 = 0

	query := repo.DB.WithContext(ctx)

	err = query.Model(&entity.SocialMedia{}).Count(&totaldata).Error
	if err != nil {
		log.Printf("[SocialMediaRepository-GetAll] error count total data: %+v \n", err)
		return
	}

	total = uint64(totaldata)

	err = query.Offset(int(offset)).Limit(int(queryparam.Limit)).Preload(clause.Associations).Find(&result).Error
	if err != nil {
		log.Printf("[SocialMediaRepository-GetAll] error get data: %+v \n", err)
		return
	}

	return
}
