package todo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"newExp/internal/model/todo"
	"strings"
)

type ItemMysql struct {
	db *sqlx.DB
}

func NewItemMysql(db *sqlx.DB) *ItemMysql {
	return &ItemMysql{
		db: db,
	}
}

func (i *ItemMysql) Get(id uint64) (*todo.Item, error) {
	item := &todo.Item{}
	query := "SELECT id,title,text,due_date,checked FROM " + todo.ItemTable + " WHERE id = ?"
	err := i.db.QueryRow(query, id).Scan(&item.ID, &item.Title, &item.Text, &item.DueDate, &item.Checked)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (i *ItemMysql) GetAllByListId(listId uint64) ([]todo.Item, error) {
	query := "SELECT items.* FROM " + todo.ItemTable + " INNER JOIN " + todo.ListItemTable + " li ON li.item_id = items.id " +
		"AND li.list_id = ?"
	rows, err := i.db.Query(query, listId)
	if err != nil {
		return nil, err
	}
	var items []todo.Item
	for rows.Next() {
		item := todo.Item{}
		err := rows.Scan(&item.ID, &item.Title, &item.Text, &item.DueDate, &item.Checked)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (i *ItemMysql) Create(listId uint64, item *todo.Item) (uint64, error) {
	ctx := context.Background()
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.ExecContext(ctx, "INSERT INTO "+todo.ItemTable+" (title, text, due_date) VALUES(?,?,?)",
		item.Title, item.Text, item.DueDate)
	if err != nil {
		return 0, err
	}
	itemId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	result, err = tx.ExecContext(ctx, "INSERT INTO "+todo.ListItemTable+" (list_id, item_id) VALUES(?, ?)",
		listId, itemId)
	if err != nil {
		return 0, err
	}

	return uint64(itemId), tx.Commit()
}

func (i *ItemMysql) Update(updateItem todo.UpdateItem, itemId uint64) error {
	var updatedQuery []string
	var args []interface{}
	if updateItem.Title != nil {
		updatedQuery = append(updatedQuery, " title = ?")
		args = append(args, updateItem.Title)
	}
	if updateItem.Text != nil {
		updatedQuery = append(updatedQuery, " text = ?")
		args = append(args, updateItem.Text)
	}
	if updateItem.Checked != nil {
		updatedQuery = append(updatedQuery, " checked = ?")
		args = append(args, updateItem.Checked)
	}
	if updateItem.DueDate != nil {
		updatedQuery = append(updatedQuery, " due_date = ?")
		args = append(args, updateItem.DueDate)
	}
	args = append(args, itemId)
	query := "UPDATE " + todo.ItemTable + " SET " + strings.Join(updatedQuery, ",") + " WHERE id = ?"
	_, err := i.db.Exec(query, args...)
	return err
}

func (i *ItemMysql) Delete(id uint64) error {
	query := "DELETE  FROM " + todo.ItemTable + " WHERE id = ?"
	_, err := i.db.Exec(query, id)
	return err
}
