package handler

import (
	model "github.com/RulerChen/NTUCS-CNAD/hw1/model"
)

type UserInterface interface {
	CreateUser(username string) error
	GetUser(username string) (model.User, error)
}
