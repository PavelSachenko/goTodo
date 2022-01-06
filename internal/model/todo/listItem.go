package todo

var (
	ListItemTable = "lists_items"
)

type ListItem struct {
	ID     uint64 `json:"id"`
	ListId uint64 `json:"list_id"`
	ItemId uint64 `json:"item_id"`
}
