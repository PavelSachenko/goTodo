package todo

var (
	UserListTable = "users_lists"
)

type UserList struct {
	ID     uint64 `json:"id"`
	UserId uint64 `json:"user_id"`
	ListId uint64 `json:"list_id"`
}
