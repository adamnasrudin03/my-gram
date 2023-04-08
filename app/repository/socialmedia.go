package repository

import (
	"log"

	"adamnasrudin03/my-gram/app/entity"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Create(input entity.SocialMedia) (res entity.SocialMedia, err error)
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
