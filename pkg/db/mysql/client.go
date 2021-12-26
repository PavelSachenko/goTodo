package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

func NewMySqlConnection(dbname string) *sqlx.DB {
	db, err := sqlx.Open("mysql", fmt.Sprintf("pavel:pass@tcp(127.0.0.1:3306)/%s", dbname))
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
