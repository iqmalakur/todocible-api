package main

import (
	"fmt"
	"net/http"
	"todocible_api/router"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	godotenv.Load()

	http.HandleFunc("/todos/", router.TodoRouter)
	http.HandleFunc("/", router.NotFoundHandler)

	port := "8000"

	fmt.Println("Server run on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
