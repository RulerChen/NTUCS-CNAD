package service

import (
	db "github.com/RulerChen/NTUCS-CNAD/hw1/infra"
	model "github.com/RulerChen/NTUCS-CNAD/hw1/model"
)

type UserService interface {
	CreateUser(username string) error
	GetUser(username string) (model.User, error)
}

type UserServiceImpl struct {
	DB db.DB
}

func NewUserService(database db.DB) UserService {
	return &UserServiceImpl{DB: database}
}

func (us *UserServiceImpl) CreateUser(username string) error {
	return us.DB.CreateUser(username)
}

func (us *UserServiceImpl) GetUser(username string) (model.User, error) {
	return us.DB.GetUser(username)
}
