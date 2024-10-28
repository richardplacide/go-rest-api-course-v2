package db

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase() (*Database, error) {
	//db, err := sqlx.Open("sqlite3", ":memory:")
	db, err := sqlx.Connect("sqlite3", "./db.sqlite")
	if err != nil {
		return nil, err
	}
	return &Database{Client: db}, nil
}

// health check Ping the database
func (d *Database) Ping(ctx context.Context) error {
	return d.Client.DB.PingContext(ctx)
}
