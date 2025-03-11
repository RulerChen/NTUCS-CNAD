package service

import (
	model "github.com/RulerChen/NTUCS-CNAD/hw1/model"
)

type UserService struct {
	userService model.UserInterface
}

func NewUserService(userService model.UserInterface) *UserService {
	return &UserService{
		userService: userService,
	}
}

func (us *UserService) CreateUser(username string) error {
	return us.userService.CreateUser(username)
}

func (us *UserService) GetUser(username string) (model.User, error) {
	return us.userService.GetUser(username)
}
