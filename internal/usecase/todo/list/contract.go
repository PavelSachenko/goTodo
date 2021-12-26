package list

import "newExp/internal/model/todo"

type List interface {
	GetList(id uint64) (*todo.List, error)
	SearchLists() ([]*todo.List, error)
	CreateList(*todo.List) error
	UpdateList(id uint64) error
	DeleteList(id uint64) error
}
