package todo

var (
	ListTable = "lists"
)

type List struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateItemInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
