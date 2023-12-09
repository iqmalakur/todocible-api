package main

import (
	"fmt"
	"net/http"
	"todolist/router"
)

func main() {
	http.HandleFunc("/todos/", router.TodoRouter)

	port := "8000"

	fmt.Println("Server run on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
