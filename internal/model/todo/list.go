package todo

import "time"

type List struct {
	ID          uint64
	UserID      uint64
	Title       string
	Description string
	DateCreated time.Time
}
