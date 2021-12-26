package todo

import (
	"github.com/jmoiron/sqlx"
	"newExp/internal/model/todo"
)

type ListRepo struct {
	db *sqlx.DB
}

func NewListMysql(db *sqlx.DB) *ListRepo {
	return &ListRepo{
		db: db,
	}
}

func (l *ListRepo) Get(id uint64) (*todo.List, error) {
	return nil, nil
}

func (l *ListRepo) All() ([]*todo.List, error) {
	return nil, nil
}

func (l *ListRepo) Create(*todo.List) error {
	return nil
}

func (l *ListRepo) Update(id uint64) error {
	return nil
}

func (l *ListRepo) Delete(id uint64) error {
	return nil
}
