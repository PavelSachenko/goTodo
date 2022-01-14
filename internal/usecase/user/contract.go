package user

import "newExp/internal/model"

//go:generate mockgen -source=contract.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user *model.User) (string, error)
	SignIn(username, password string) (string, error)
	ParseToke(token string) (uint64, error)
	User
}

type User interface {
	GetUser(id uint64) (*model.User, error)
}
