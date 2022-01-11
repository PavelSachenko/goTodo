package response

type ItemResponse struct {
	ID      uint64 `json:"id"`
	Title   string `json:"title"`
	Text    string `json:"text"`
	DueDate string `json:"dueDate"`
	Checked bool   `json:"checked"`
}
