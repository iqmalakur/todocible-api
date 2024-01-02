package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	host := os.Getenv("dbhost")
	if host == "" {
		host = "localhost"
	}

	user := os.Getenv("dbuser")
	if user == "" {
		user = "postgres"
	}

	pass := os.Getenv("dbpass")
	if pass == "" {
		pass = "postgres"
	}

	dbname := os.Getenv("dbname")
	if dbname == "" {
		dbname = "todocible"
	}

	sslmode := os.Getenv("sslmode")
	if sslmode == "" {
		sslmode = "disable"
	}

	conn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", user, pass, host, dbname, sslmode)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(time.Hour)

	return db
}
