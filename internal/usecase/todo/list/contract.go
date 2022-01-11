package list

import "newExp/internal/model/todo"

type List interface {
	GetList(id uint64, userId uint64) (*todo.List, error)
	SearchLists(userId uint64) ([]todo.List, error)
	CreateList(uint64, *todo.List) (uint64, error)
	UpdateList(list todo.UpdateItemInput, listId, userId uint64) error
	DeleteList(id uint64, userId uint64) error
	CheckAccessRight(listId uint64, userId uint64) (bool, error)
}
