package model

var UserTable = "users"

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}
