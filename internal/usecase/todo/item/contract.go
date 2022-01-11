package item

import "newExp/internal/model/todo"

type Item interface {
	CreateItem(listId uint64, item *todo.InputItemRequest) (uint64, error)
	DeleteItem(id uint64) error
	UpdateItem(updateItem todo.UpdateItem, itemId uint64) error
	GetAllFromList(listId uint64) ([]todo.Item, error)
	GetById(id uint64) (*todo.Item, error)
}
