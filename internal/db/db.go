package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// OpenDbConnection opens a single Postgres connection, panicking if it cannot
// be established. The caller is responsible for closing it.
func OpenDbConnection(ctx context.Context, connStr string) *pgx.Conn {
	db, err := pgx.Connect(ctx, connStr)
	if err != nil {
		panic(err)
	}

	return db
}
