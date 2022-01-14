package usecase

import (
	"newExp/internal/config"
	"newExp/internal/repository"
	"newExp/internal/usecase/todo/item"
	"newExp/internal/usecase/todo/list"
	"newExp/internal/usecase/user"
)

type SuperService struct {
	List   list.List
	Auth   user.Authorization
	Item   item.Item
	Config *config.Config
}

func NewSuperService(repository *repository.SuperRepository, cnf *config.Config) *SuperService {
	return &SuperService{
		Config: cnf,
		List:   list.NewService(repository.List),
		Item:   item.NewService(repository.Item),
		Auth:   user.NewService(repository.Auth, repository.User),
	}
}
