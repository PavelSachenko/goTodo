package todo

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"newExp/internal/model/todo"
)

type ListRepo struct {
	db *sqlx.DB
}

func NewListMysql(db *sqlx.DB) *ListRepo {
	return &ListRepo{
		db: db,
	}
}

func (l *ListRepo) Get(id uint64, userId uint64) (*todo.List, error) {
	query := "SELECT lists.id, lists.title, lists.description FROM " + todo.ListTable + " LEFT JOIN " +
		todo.UserListTable + " ul ON ul.list_id = lists.id WHERE ul.user_id = ? AND lists.id = ?"
	rows, err := l.db.Query(query, userId, id)
	if err != nil {
		return nil, err
	}
	var list todo.List
	for rows.Next() {
		err = rows.Scan(&list.ID, &list.Title, &list.Description)
	}
	return &list, err
}

func (l *ListRepo) All(userId uint64) ([]todo.List, error) {
	query := "SELECT lists.id, lists.title, lists.description FROM " + todo.ListTable + " LEFT JOIN " +
		todo.UserListTable + " ul ON ul.list_id = lists.id where ul.user_id = ? "
	rows, err := l.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	var lists []todo.List
	for rows.Next() {
		list := todo.List{}
		err := rows.Scan(&list.ID, &list.Title, &list.Description)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return lists, nil
}

func (l *ListRepo) Create(userId uint64, list todo.InputListRequest) (uint64, error) {
	ctx := context.Background()
	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.ExecContext(ctx, "INSERT INTO "+todo.ListTable+" (title, description) VALUES(?, ?)",
		list.Title, list.Description)
	if err != nil {
		return 0, err
	}
	listId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	result, err = tx.ExecContext(ctx, "INSERT INTO "+todo.UserListTable+" (user_id, list_id) VALUES(?, ?)",
		userId, listId)
	if err != nil {
		return 0, err
	}

	return uint64(listId), tx.Commit()
}

func (l *ListRepo) Update(list todo.UpdateItemInput, listId, userId uint64) error {
	query := "UPDATE " + todo.ListTable + " LEFT JOIN " + todo.UserListTable + " ul ON ul.list_id = lists.id  SET lists.title = ?, lists.description = ? WHERE ul.user_id = ? AND lists.id = ?"
	_, err := l.db.Exec(query, list.Title, list.Description, userId, listId)
	return err
}

func (l *ListRepo) Delete(id uint64, userId uint64) error {
	query := "DELETE " + todo.ListTable + " FROM " + todo.ListTable + " LEFT JOIN " +
		todo.UserListTable + " ul ON ul.list_id = lists.id WHERE ul.user_id = ? AND lists.id = ?"
	_, err := l.db.Exec(query, userId, id)
	return err
}
func (l *ListRepo) ListIsBelongToUser(listId uint64, userId uint64) (bool, error) {
	query := "SELECT lists.id FROM " + todo.ListTable + " INNER JOIN " + todo.UserListTable + " ul on lists.id = ul.list_id AND ul.user_id = ? AND ul.list_id = ?"
	var exist int
	err := l.db.QueryRow(query, userId, listId).Scan(&exist)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return true, err
}
