package user

import "newExp/internal/model"

type Authorization interface {
	CreateUser(user *model.User) (uint64, error)
	SignIn(username, password string) (*model.User, error)
}

type User interface {
	Get(Id uint64) (*model.User, error)
}
