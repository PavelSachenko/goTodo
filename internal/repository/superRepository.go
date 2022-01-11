package repository

import (
	"github.com/jmoiron/sqlx"
	"newExp/internal/repository/todo"
	"newExp/internal/repository/user"
)

type SuperRepository struct {
	List todo.List
	Item todo.Item
	Auth user.Authorization
	User user.User
}

func NewSuperRepository(db *sqlx.DB) *SuperRepository {
	return &SuperRepository{
		List: todo.NewListMysql(db),
		Item: todo.NewItemMysql(db),
		Auth: user.NewUserMySql(db),
		User: user.NewUserMySql(db),
	}
}
