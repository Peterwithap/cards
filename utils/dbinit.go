package utils

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func dbinit() *sql.DB {
	cfg := mysql.Config{
		User:   "fela",
		Passwd: "africa70",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "iconEditor",
	}
	// Get a database handle.
	var err error
	var db *sql.DB
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db
}
