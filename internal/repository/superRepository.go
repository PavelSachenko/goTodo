package repository

import (
	"github.com/jmoiron/sqlx"
	"newExp/internal/repository/todo"
)

type SuperRepository struct {
	List todo.List
	Item todo.Item
}

func NewSuperRepository(db *sqlx.DB) *SuperRepository {
	return &SuperRepository{
		List: todo.NewListMysql(db),
		Item: todo.NewItemMysql(db),
	}
}
