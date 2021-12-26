package user

import "newExp/internal/model"

type Authorization interface {
	CreateUser(user *model.User) (string, error)
	SignIn(username, password string) (string, error)
	ParseToke(token string) (uint64, error)
}

type User interface {
	GetUser(id uint64) (*model.User, error)
}
