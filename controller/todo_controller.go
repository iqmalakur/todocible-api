package controller

import (
	"fmt"
	"net/http"
	"text/template"
	"todolist/service"
)

type TodoController struct {
	service *service.TodoService
}

func NewTodoController() *TodoController {
	return &TodoController{service.NewTodoService()}
}

func (todoController *TodoController) Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		todoController.show(w, r)
	case "POST":
		todoController.create(w, r)
	}

}

func (todoController *TodoController) show(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./view/index.html")
	t.Execute(w, todoController.service.GetAll())
}

func (todoController *TodoController) create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Println("Error :", err)
	}

	todo := todoController.service.Create(
		r.FormValue("title"),
		r.FormValue("description"),
	)

	if todo == nil {
		fmt.Println("Error : cannot create a new todo")
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
