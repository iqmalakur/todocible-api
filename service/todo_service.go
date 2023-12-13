package service

import (
	"errors"
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

func (todoService *TodoService) Get(id string) (*entity.Todo, error) {
	todo := todoService.todoRepository.Find(id)

	if todo == nil {
		return nil, errors.New("todo with id " + id + " is not found")
	}

	return todo, nil
}

func (todoService *TodoService) Update(id string, body dto.TodoRequest) (*entity.Todo, error) {
	if body.Title == "" {
		return nil, errors.New("'title' is not allowed to be empty")
	}

	if body.Description == "" {
		return nil, errors.New("'description' is not allowed to be empty")
	}

	todo := todoService.todoRepository.Update(id, &entity.Todo{
		Title:       body.Title,
		Description: body.Description,
	})

	if todo == nil {
		return nil, errors.New("todo with id " + id + " is not found")
	}

	return todo, nil
}

func (todoService *TodoService) SetCompleted(id string, completed bool) (*entity.Todo, error) {
	todo := todoService.todoRepository.Find(id)

	if todo == nil {
		return nil, errors.New("todo with id " + id + " is not found")
	}

	success := todoService.todoRepository.SetCompleted(id, completed)

	if !success {
		return nil, errors.New("todo with id " + id + " is not found")
	}

	return todo, nil
}

func (todoService *TodoService) Delete(id string) (*entity.Todo, error) {
	todo := todoService.todoRepository.Find(id)

	if todo == nil {
		return nil, errors.New("todo with id " + id + " is not found")
	}

	success := todoService.todoRepository.Delete(id)

	if !success {
		return nil, errors.New("todo with id " + id + " is not found")
	}

	return todo, nil
}
