package usecase

import (
	"newExp/internal/repository"
	"newExp/internal/usecase/todo/list"
	"newExp/internal/usecase/user"
)

type SuperService struct {
	List *list.Service
	Auth *user.Service
}

func NewSuperService(repository *repository.SuperRepository) *SuperService {
	return &SuperService{
		List: list.NewService(repository.List),
		Auth: user.NewService(repository.Auth, repository.User),
	}
}
