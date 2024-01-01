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

	todos, err := c.service.GetAll()
	if err != nil {
		c.writer.WriteHeader(http.StatusInternalServerError)
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

func (c *TodoController) Show(id string) {
	defer c.service.Close()

	todo, err := c.service.Get(id)

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

func (c *TodoController) Update(id string) {
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

	todo, err := c.service.Update(id, body)

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

func (c *TodoController) Delete(id string) {
	defer c.service.Close()

	todo, err := c.service.Delete(id)
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

func (c *TodoController) SetDone(id, action string) {
	defer c.service.Close()

	status := false
	if action == "done" {
		status = true
	}

	todo, err := c.service.SetCompleted(id, status)
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
		Message: "success update todo status",
		Data:    todo,
	})
}
