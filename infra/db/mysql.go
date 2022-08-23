package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func GetDatabaseConnection() *sql.DB {
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/clean_arch")
	if err != nil {
		panic(err)
	}
	return db
}

func CloseDatabaseConnection() {
	db.Close()
}
