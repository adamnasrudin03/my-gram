package service

import (
	"log"

	"github.com/adamnasrudin03/my-gram/app/dto"
	"github.com/adamnasrudin03/my-gram/app/entity"
	"github.com/adamnasrudin03/my-gram/app/repository"
	"github.com/adamnasrudin03/my-gram/pkg/helpers"
)

type UserService interface {
	Register(input entity.User) (res entity.User, err error)
	Login(input dto.LoginReq) (token string, err error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(UserRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: UserRepo,
	}
}

func (srv *userService) Register(input entity.User) (res entity.User, err error) {
	res, err = srv.userRepository.Register(input)
	if err != nil {
		log.Printf("[UserService-Register] error create data: %+v \n", err)
		return
	}

	return
}

func (srv *userService) Login(input dto.LoginReq) (token string, err error) {
	user, err := srv.userRepository.Login(input)
	if err != nil {
		log.Printf("[UserService-Login] error: %+v \n", err)
		return
	}

	token, err = helpers.GenerateToken(user.ID, user.Username, user.Email)
	if err != nil {
		log.Printf("[UserService-Login] error generate token: %+v \n", err)
		return
	}

	return
}
