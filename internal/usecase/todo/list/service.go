package list

import (
	"newExp/internal/app"
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

func (s *Service) GetList(id uint64) (*todo.List, error) {
	result, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, app.ErrNotFound
	}
	return result, nil
}

func (s *Service) SearchLists() ([]*todo.List, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) CreateList(list *todo.List) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateList(id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) DeleteList(id uint64) error {
	//TODO implement me
	panic("implement me")
}
