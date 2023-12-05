package main

import (
	"net/http"
	"todolist/controller"
)

func main() {
	todoController := controller.New()

	http.HandleFunc("/", todoController.Index)

	http.ListenAndServe(":8000", nil)
}
