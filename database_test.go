package go_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:7qvt6t2738@tcp(localhost:3306)/go_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
