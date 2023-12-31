package database

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func GetConnection() (*sql.DB, error) {
	conn := "postgres://postgres:postgres@localhost/todocible?sslmode=disable"

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}
