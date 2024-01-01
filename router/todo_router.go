package router

import (
	"fmt"
	"net/http"
	"strings"
	"todocible_api/controller"
)

func TodoRouter(w http.ResponseWriter, r *http.Request) {
	controller := controller.NewTodoController(w, r)

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

	HeaderConfig(w)

	switch {
	case todoId == "" && action == "" && r.Method == "GET":
		controller.Index()
	case todoId == "" && action == "" && r.Method == "POST":
		controller.Create()
	case todoId != "" && action == "" && r.Method == "GET":
		controller.Show(todoId)
	case todoId != "" && action == "" && r.Method == "DELETE":
		controller.Delete(todoId)
	case todoId != "" && r.Method == "PUT":
		switch action {
		case "done":
			fallthrough
		case "undone":
			controller.SetDone(todoId, action)
		default:
			controller.Update(todoId)
		}
	default:
		NotFoundHandler(w, r)
	}
}
