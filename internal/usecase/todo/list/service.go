package list

import (
	"newExp/internal/model/todo"
	repository "newExp/internal/repository/todo"
)

type Service struct {
	repo repository.List
}

func NewService(repo repository.List) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetList(id, userId uint64) (*todo.List, error) {
	result, err := s.repo.Get(id, userId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) SearchLists(userId uint64) ([]todo.List, error) {
	return s.repo.All(userId)
}

func (s *Service) CreateList(usedId uint64, list todo.InputListRequest) (uint64, error) {
	return s.repo.Create(usedId, list)
}

func (s *Service) UpdateList(list todo.UpdateItemInput, listId, userId uint64) error {
	return s.repo.Update(list, listId, userId)
}

func (s *Service) DeleteList(id uint64, userId uint64) error {
	return s.repo.Delete(id, userId)
}
func (s *Service) CheckAccessRight(listId uint64, userId uint64) (bool, error) {
	return s.repo.ListIsBelongToUser(listId, userId)
}
