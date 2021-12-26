package usecase

import (
	"newExp/internal/repository"
	"newExp/internal/usecase/todo/list"
)

type SuperService struct {
	List *list.Service
}

func NewSuperService(repository *repository.SuperRepository) *SuperService {
	return &SuperService{
		List: list.NewService(repository.List),
	}
}
