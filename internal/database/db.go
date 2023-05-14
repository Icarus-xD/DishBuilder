package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Init(driver string, url string) *sql.DB {
	db, err := sql.Open(driver, url)

	if err != nil {
		log.Fatalln(err)
	}

	return db
}