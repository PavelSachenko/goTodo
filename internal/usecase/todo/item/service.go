package item

import (
	"newExp/internal/model/todo"
	repository "newExp/internal/repository/todo"
)

type Service struct {
	repo repository.Item
}

func NewService(repo repository.Item) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateItem(listId uint64, item *todo.Item) (uint64, error) {
	return s.repo.Create(listId, item)
}

func (s *Service) DeleteItem(id uint64) error {
	return s.repo.Delete(id)
}

func (s *Service) UpdateItem(updateItem todo.UpdateItem, itemId uint64) error {
	return s.repo.Update(updateItem, itemId)
}

func (s *Service) GetAllFromList(listId uint64) ([]todo.Item, error) {
	return s.repo.GetAllByListId(listId)
}

func (s *Service) GetById(id uint64) (*todo.Item, error) {
	return s.repo.Get(id)
}
