package go_database

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnetion()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(id, name) VALUES ('Herman', 'dev')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Insert Data Success!")
}

func TestQuerySql(t *testing.T) {
	db := GetConnetion()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	fmt.Println("Select Data Success!")
}
