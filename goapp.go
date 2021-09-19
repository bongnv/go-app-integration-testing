package goapp

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// LoadDataFromDB loads a string from DB to demonstrate a function working with DB.
func LoadDataFromDB() (string, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		return "", err
	}
	defer conn.Close(ctx)
	var name string
	row := conn.QueryRow(ctx, "select name from sample_table where id=1")
	if err := row.Scan(&name); err != nil {
		return "", err
	}
	return name, nil
}
