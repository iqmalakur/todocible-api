package controller

import (
	"encoding/json"
	"net/http"
	"strings"
	"todocible_api/dto"
	"todocible_api/service"
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
		Message: "success get all todos",
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

func (todoController *TodoController) Show(w http.ResponseWriter, r *http.Request) {
	todoId := r.URL.Path[len("/todos/"):]

	todo, err := todoController.service.Get(todoId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(dto.TodoResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.TodoResponse{
		Success: true,
		Message: "success get todo",
		Data:    todo,
	})
}

func (todoController *TodoController) Update(w http.ResponseWriter, r *http.Request) {
	todoId := r.URL.Path[len("/todos/"):]

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

	todo, err := todoController.service.Update(todoId, body)

	if err != nil {
		if strings.HasSuffix(err.Error(), "not found") {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}

		json.NewEncoder(w).Encode(dto.TodoResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.TodoResponse{
		Success: true,
		Message: "success update todo",
		Data:    todo,
	})
}

func (todoController *TodoController) Delete(w http.ResponseWriter, r *http.Request) {
	todoId := r.URL.Path[len("/todos/"):]
	todo, err := todoController.service.Delete(todoId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(dto.TodoResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.TodoResponse{
		Success: true,
		Message: "success delete todo",
		Data:    todo,
	})
}

func (todoController *TodoController) SetDone(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path[len("/todos/"):], "/")
	todoId := params[0]
	action := params[1]

	status := false

	if action == "done" {
		status = true
	}

	todo, err := todoController.service.SetCompleted(todoId, status)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(dto.TodoResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.TodoResponse{
		Success: true,
		Message: "success set " + action + " todo",
		Data:    todo,
	})
}
