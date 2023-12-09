package main

import (
	"fmt"
	"net/http"
	"todolist/controller"
)

func main() {
	todoController := controller.NewTodoController()

	http.HandleFunc("/", todoController.Index)

	fmt.Println("Server run on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
