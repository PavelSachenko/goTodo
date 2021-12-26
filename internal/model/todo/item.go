package todo

import "time"

type Items struct {
	ID          uint64
	ListId      uint64
	Text        string
	DueDate     time.Time
	Checked     bool
	DateCreated time.Time
}
