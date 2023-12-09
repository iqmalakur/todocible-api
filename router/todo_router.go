package router

import (
	"net/http"
	"todolist/controller"
)

func TodoRouter(w http.ResponseWriter, r *http.Request) {
	todoId := r.URL.Path[len("/todos/"):]
	todoController := controller.NewTodoController()

	switch {
	case todoId == "" && r.Method == "GET":
		todoController.Index(w)
	}
}
