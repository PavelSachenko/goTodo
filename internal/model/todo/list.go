package todo

var (
	ListTable = "lists"
)

type List struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type InputListRequest struct {
	Title       string `form:"title" json:"title" binding:"required"`
	Description string `form:"description" json:"description"`
}

type UpdateItemInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
