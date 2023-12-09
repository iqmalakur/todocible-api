package controller

import (
	"encoding/json"
	"net/http"
	"todolist/service"
)

type TodoController struct {
	service *service.TodoService
}

func NewTodoController() *TodoController {
	return &TodoController{service.NewTodoService()}
}

func (todoController *TodoController) Index(w http.ResponseWriter) {
	todoService := service.NewTodoService()
	res := todoService.GetAll()
	json.NewEncoder(w).Encode(res)
}

// func (todoController *TodoController) create(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()

// 	if err != nil {
// 		fmt.Println("Error :", err)
// 	}

// 	todo := todoController.service.Create(
// 		r.FormValue("title"),
// 		r.FormValue("description"),
// 	)

// 	if todo == nil {
// 		fmt.Println("Error : cannot create a new todo")
// 	}

// 	http.Redirect(w, r, "/", http.StatusFound)
// }
