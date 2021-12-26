package todo

import "newExp/internal/model/todo"

type Item interface {
	Get()
	GetAllByListId()
	Create()
	Update()
	Delete()
}

type List interface {
	Get(id uint64) (*todo.List, error)
	All() ([]*todo.List, error)
	Create(*todo.List) error
	Update(id uint64) error
	Delete(id uint64) error
}
