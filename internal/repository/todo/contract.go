package todo

import "newExp/internal/model/todo"

type Item interface {
	Get(id uint64) (*todo.Item, error)
	GetAllByListId(listId uint64) ([]todo.Item, error)
	Create(listId uint64, item *todo.InputItemRequest) (uint64, error)
	Update(updateItem todo.UpdateItem, itemId uint64) error
	Delete(id uint64) error
}

type List interface {
	Get(id uint64, userId uint64) (*todo.List, error)
	All(userId uint64) ([]todo.List, error)
	Create(uint64, todo.InputListRequest) (uint64, error)
	Update(list todo.UpdateItemInput, listId, userId uint64) error
	Delete(id uint64, userId uint64) error
	ListIsBelongToUser(listId uint64, userId uint64) (bool, error)
}
