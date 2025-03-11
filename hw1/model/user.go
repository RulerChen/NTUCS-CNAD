package model

type User struct {
	Username string
}

type UserInterface interface {
	CreateUser(username string) error
	GetUser(username string) (User, error)
}
