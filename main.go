package main

import (
	"net/http"
	"todolist/controller"
)

func main() {
	todoController := controller.NewTodoController()

	http.HandleFunc("/", todoController.Index)

	http.ListenAndServe(":8000", nil)
}
