package todo

import (
	"database/sql"
)

var (
	ItemTable = "items"
)

type Item struct {
	ID      uint64
	Title   string
	Text    string
	DueDate sql.NullString
	Checked bool
}

type UpdateItem struct {
	Title   *string `json:"title"`
	Text    *string `json:"text"`
	DueDate *string `json:"dueDate"`
	Checked *bool   `json:"checked"`
}
