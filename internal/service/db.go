package service

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Database *sql.DB

func InitializeDatabase() error {
	var err error
	Database, err = sql.Open("postgres", "host=db port=5432 user=postgres password=123 dbname=go_db sslmode=disable")
	if err != nil {
		return err
	}
	err = Database.Ping()
	if err != nil {
		return err
	}
	return nil
}
