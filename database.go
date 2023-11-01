package go_database

import (
	"database/sql"
	"time"
)

func GetConnetion() *sql.DB {
	db, err := sql.Open("mysql", "root:7qvt6t2738@tcp(localhost:3306)/go_db")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(60 * time.Minute)
	return db
}
