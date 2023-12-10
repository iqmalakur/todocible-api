package controller

import (
	"encoding/json"
	"net/http"
	"todolist/dto"
	"todolist/service"
)

type TodoController struct {
	service *service.TodoService
}

func NewTodoController() *TodoController {
	return &TodoController{service.NewTodoService()}
}

func (todoController *TodoController) Index(w http.ResponseWriter, r *http.Request) {
	todos := todoController.service.GetAll()
	json.NewEncoder(w).Encode(dto.TodoResponse{
		Success: true,
		Message: "success get all data",
		Data:    todos,
	})
}

func (todoController *TodoController) Create(w http.ResponseWriter, r *http.Request) {
	var body dto.TodoRequest
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.TodoResponse{
			Success: false,
			Message: "bad request",
			Data:    nil,
		})
		return
	}

	todo, err := todoController.service.Create(body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.TodoResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto.TodoResponse{
		Success: true,
		Message: "success create new todo",
		Data:    todo,
	})
}
