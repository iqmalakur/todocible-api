package database

import (
	"context"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func TestConnection(t *testing.T) {
	godotenv.Load("../.env")

	db := GetConnection()
	defer db.Close()

	query := "SELECT title, description FROM todos LIMIT 10"
	rows, err := db.QueryContext(context.Background(), query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println("OK")
	fmt.Println(db)

	var title, description string

	for rows.Next() {
		err = rows.Scan(&title, &description)
		if err != nil {
			panic(err)
		}

		fmt.Println("=====================")
		fmt.Println(title)
		fmt.Println(description)
	}
}
