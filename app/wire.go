package app

import (
	"adamnasrudin03/my-gram/app/repository"
	"adamnasrudin03/my-gram/app/service"

	"gorm.io/gorm"
)

func WiringRepository(db *gorm.DB) *repository.Repositories {
	return &repository.Repositories{
		User:        repository.NewUserRepository(db),
		SocialMedia: repository.NewSocialMediaRepository(db),
		Comment:     repository.NewCommentRepository(db),
	}
}

func WiringService(repo *repository.Repositories) *service.Services {
	return &service.Services{
		User:        service.NewUserService(repo.User),
		SocialMedia: service.NewSocialMediaService(repo.SocialMedia),
		Comment:     service.NewCommentService(repo.Comment),
	}
}
