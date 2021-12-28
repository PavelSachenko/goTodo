package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
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
