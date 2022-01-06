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
	Get(id uint64, userId uint64) (*todo.List, error)
	All(userId uint64) ([]todo.List, error)
	Create(uint64, *todo.List) (uint64, error)
	Update(list todo.UpdateItemInput, listId, userId uint64) error
	Delete(id uint64, userId uint64) error
}
