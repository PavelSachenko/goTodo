package todo

import (
	"database/sql"
)

var (
	ItemTable = "items"
)

type Item struct {
	ID      uint64  `json:"id"`
	Title   string  `json:"title"`
	Text    string  `json:"text"`
	DueDate *string `json:"dueDate"`
	Checked bool    `json:"checked"`
}

type UpdateItem struct {
	Title   *string `json:"title"`
	Text    *string `json:"text"`
	DueDate *string `json:"dueDate"`
	Checked *bool   `json:"checked"`
}

type InputItemRequest struct {
	Title   string         `form:"title" json:"title" binding:"required"`
	Text    string         `form:"text" json:"text" binding:"required"`
	DueDate sql.NullString `form:"due_date" json:"due_date"`
}
