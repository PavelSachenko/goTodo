package todo

import "github.com/jmoiron/sqlx"

type ItemMysql struct {
	db *sqlx.DB
}

func NewItemMysql(db *sqlx.DB) *ItemMysql {
	return &ItemMysql{
		db: db,
	}
}

func (i *ItemMysql) Get() {
}

func (i *ItemMysql) GetAllByListId() {
}

func (i *ItemMysql) Create() {
}

func (i *ItemMysql) Update() {
}

func (i *ItemMysql) Delete() {
}
