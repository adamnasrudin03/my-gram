package service

import (
	"log"

	"adamnasrudin03/my-gram/app/dto"
	"adamnasrudin03/my-gram/app/entity"
	"adamnasrudin03/my-gram/app/repository"
	"adamnasrudin03/my-gram/pkg/helpers"
)

type UserService interface {
	Register(input dto.RegisterReq) (res entity.User, err error)
	Login(input dto.LoginReq) (res dto.LoginRes, err error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(UserRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: UserRepo,
	}
}

func (srv *userService) Register(input dto.RegisterReq) (res entity.User, err error) {
	user := entity.User{
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
		Age:      input.Age,
	}

	res, err = srv.userRepository.Register(user)
	if err != nil {
		log.Printf("[UserService-Register] error create data: %+v \n", err)
		return
	}

	res.Password = ""

	return
}

func (srv *userService) Login(input dto.LoginReq) (res dto.LoginRes, err error) {
	user, err := srv.userRepository.Login(input)
	if err != nil {
		log.Printf("[UserService-Login] error: %+v \n", err)
		return
	}

	res.Token, err = helpers.GenerateToken(user.ID, user.Username, user.Email)
	if err != nil {
		log.Printf("[UserService-Login] error generate token: %+v \n", err)
		return
	}

	return
}
