package controller

import (
	"net/http"
	"text/template"
)

type TodoController struct{}

func (todoController *TodoController) Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./view/index.html")
	t.Execute(w, nil)
}
