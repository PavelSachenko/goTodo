package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func NewMySqlConnection(username, password, host, port, database string) *sqlx.DB {
	fmt.Printf("%s:%s@tcp(%s:%s)/%s\n", username, password, host, port, database)
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database))
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func SetNowYmdIhs() sql.NullString {
	time := time.Now()
	return sql.NullString{String: time.Format("2006-01-02 15:01:05"), Valid: true}
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
