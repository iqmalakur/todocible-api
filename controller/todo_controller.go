package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"todocible_api/database"
	"todocible_api/dto"
	"todocible_api/service"
)

type TodoController struct {
	service service.TodoService
	writer  http.ResponseWriter
	request *http.Request
}

func NewTodoController(writer http.ResponseWriter, request *http.Request) TodoController {
	return TodoController{
		service: service.NewTodoService(),
		writer:  writer,
		request: request,
	}
}

func (c *TodoController) Index() {
	defer c.service.Close()

	todos := c.service.GetAll()

	c.writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.writer).Encode(dto.TodoResponse{
		Success: true,
		Message: "success get all todos",
		Data:    todos,
	})
}

func (c *TodoController) Create() {
	defer c.service.Close()

	var body dto.TodoRequest

	err := json.NewDecoder(c.request.Body).Decode(&body)
	if err != nil {
		c.writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.writer).Encode(dto.TodoResponse{
			Success: false,
			Message: "bad request",
			Data:    nil,
		})
		return
	}

	todo, err := c.service.Create(body)
	if err != nil {
		if errors.Is(err, database.ConnectionError) {
			c.writer.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(c.writer).Encode(dto.TodoResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		c.writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.writer).Encode(dto.TodoResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(c.writer).Encode(dto.TodoResponse{
		Success: true,
		Message: "success create new todo",
		Data:    todo,
	})
}

func (c *TodoController) Show() {
	todoId := c.request.URL.Path[len("/todos/"):]

	todo, err := c.service.Get(todoId)

	if err != nil {
		c.writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(c.writer).Encode(dto.TodoResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.writer).Encode(dto.TodoResponse{
		Success: true,
		Message: "success get todo",
		Data:    todo,
	})
}

func (c *TodoController) Update() {
	todoId := c.request.URL.Path[len("/todos/"):]

	var body dto.TodoRequest
	err := json.NewDecoder(c.request.Body).Decode(&body)

	if err != nil {
		c.writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.writer).Encode(dto.TodoResponse{
			Success: false,
			Message: "bad request",
			Data:    nil,
		})
		return
	}

	todo, err := c.service.Update(todoId, body)

	if err != nil {
		if strings.HasSuffix(err.Error(), "not found") {
			c.writer.WriteHeader(http.StatusNotFound)
		} else {
			c.writer.WriteHeader(http.StatusBadRequest)
		}

		json.NewEncoder(c.writer).Encode(dto.TodoResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.writer).Encode(dto.TodoResponse{
		Success: true,
		Message: "success update todo",
		Data:    todo,
	})
}

func (c *TodoController) Delete() {
	todoId := c.request.URL.Path[len("/todos/"):]
	todo, err := c.service.Delete(todoId)

	if err != nil {
		c.writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(c.writer).Encode(dto.TodoResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.writer).Encode(dto.TodoResponse{
		Success: true,
		Message: "success delete todo",
		Data:    todo,
	})
}

func (c *TodoController) SetDone() {
	params := strings.Split(c.request.URL.Path[len("/todos/"):], "/")
	todoId := params[0]
	action := params[1]

	status := false

	if action == "done" {
		status = true
	}

	todo, err := c.service.SetCompleted(todoId, status)

	if err != nil {
		c.writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(c.writer).Encode(dto.TodoResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.writer).Encode(dto.TodoResponse{
		Success: true,
		Message: "success set " + action + " todo",
		Data:    todo,
	})
}
