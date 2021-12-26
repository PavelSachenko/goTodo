package user

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"newExp/internal/model"
)

type UserMysql struct {
	db *sqlx.DB
}

func NewUserMySql(db *sqlx.DB) *UserMysql {
	return &UserMysql{
		db: db,
	}
}

func (u *UserMysql) Get(id uint64) (*model.User, error) {
	query := "SELECT id, name, password_hash FROM " + model.UserTable + " where id = ? "
	stmt, err := u.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(id)
	user := &model.User{}
	err = row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserMysql) CreateUser(user *model.User) (uint64, error) {
	exist, err := u.rowExists("SELECT * FROM "+model.UserTable+" WHERE name = ? LIMIT 1", user.Username)
	if err != nil {
		return 0, err
	}

	if exist == true {
		return 0, errors.New("User already exists")
	}

	query := "INSERT INTO " + model.UserTable + " (name, password_hash) VALUES(?, ?)"
	row, err := u.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	result, err := row.Exec(user.Username, user.Password)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func (u *UserMysql) SignIn(username, password string) (*model.User, error) {
	query := "SELECT id, name, password_hash FROM " + model.UserTable + " where name = ? AND password_hash = ?"
	stmt, err := u.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(username, password)
	user := &model.User{}
	err = row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserMysql) rowExists(query string, args ...interface{}) (bool, error) {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := u.db.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists, nil
}
