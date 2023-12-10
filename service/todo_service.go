package service

import (
	"errors"
	"strconv"
	"todolist/dto"
	"todolist/entity"
	"todolist/repository"
)

type TodoService struct {
	todoRepository *repository.TodoRepository
}

func NewTodoService() *TodoService {
	return &TodoService{repository.NewTodoRepository()}
}

func (todoService *TodoService) Create(todo dto.TodoRequest) (*entity.Todo, error) {
	if todo.Title == "" {
		return nil, errors.New("'title' is not allowed to be empty")
	}

	if todo.Description == "" {
		return nil, errors.New("'description' is not allowed to be empty")
	}

	newTodo := todoService.todoRepository.Create(todo.Title, todo.Description)
	return newTodo, nil
}

func (todoService *TodoService) GetAll() []*entity.Todo {
	todos := todoService.todoRepository.FindAll()
	return todos
}

func (todoService *TodoService) Get(id int) (*entity.Todo, error) {
	todo := todoService.todoRepository.Find(id)

	if todo == nil {
		return nil, errors.New("todo with id " + strconv.Itoa(id) + " is not found")
	}

	return todo, nil
}

func (todoService *TodoService) Update(id int, newTodo *entity.Todo) *entity.Todo {
	todo := todoService.todoRepository.Update(id, newTodo)
	return todo
}

func (todoService *TodoService) SetCompleted(id int, completed bool) bool {
	return todoService.todoRepository.SetCompleted(id, completed)
}

func (todoService *TodoService) Delete(id int) bool {
	return todoService.todoRepository.Delete(id)
}
