package db

import (
	"database/sql"
	"time"
)

func ConnectionDB() *sql.DB {
	db, err := sql.Open("mysql", "root@/perpustakaan")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
