package router

import (
	"fmt"
	"net/http"
	"strings"
	"todolist/controller"
)

var todoController = controller.NewTodoController()

func TodoRouter(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path[len("/todos/"):], "/")

	todoId := ""
	action := ""

	if len(params) > 0 {
		todoId = params[0]
	}

	if len(params) > 1 {
		action = params[1]
	}

	fmt.Println(r.Method, r.URL.Path)

	switch {
	case todoId == "" && r.Method == "GET":
		todoController.Index(w, r)
	case todoId == "" && r.Method == "POST":
		todoController.Create(w, r)
	case r.Method == "GET":
		todoController.Show(w, r)
	case r.Method == "PUT":
		switch action {
		case "done":
			fallthrough
		case "undone":
			todoController.SetDone(w, r)
		default:
			todoController.Update(w, r)
		}
	case r.Method == "DELETE":
		todoController.Delete(w, r)
	}
}
