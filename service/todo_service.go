package service

import (
	"errors"
	"todocible_api/database"
	"todocible_api/dto"
	"todocible_api/entity"
	"todocible_api/repository"
)

type TodoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService() TodoService {
	return TodoService{repository.NewTodoRepository(database.GetConnection())}
}

func (s *TodoService) Create(todo dto.TodoRequest) (*entity.Todo, error) {
	if todo.Title == "" {
		return nil, errors.New("'title' is not allowed to be empty")
	}

	newTodo, err := s.todoRepository.Create(todo)
	if err != nil {
		return nil, database.ConnectionError
	}

	return newTodo, nil
}

func (s *TodoService) GetAll() ([]entity.Todo, error) {
	todos, err := s.todoRepository.FindAll()
	if err != nil {
		return nil, database.ConnectionError
	}

	return todos, nil
}

func (s *TodoService) Get(id string) (*entity.Todo, error) {
	todo := s.todoRepository.Find(id)

	if todo == nil {
		return nil, errors.New("todo with id " + id + " is not found")
	}

	return todo, nil
}

func (s *TodoService) Update(id string, body dto.TodoRequest) (*entity.Todo, error) {
	if body.Title == "" {
		return nil, errors.New("'title' is not allowed to be empty")
	}

	if body.Description == "" {
		return nil, errors.New("'description' is not allowed to be empty")
	}

	todo := s.todoRepository.Update(id, body)

	if todo == nil {
		return nil, errors.New("todo with id " + id + " is not found")
	}

	return todo, nil
}

func (s *TodoService) SetCompleted(id string, completed bool) (*entity.Todo, error) {
	todo := s.todoRepository.Find(id)

	if todo == nil {
		return nil, errors.New("todo with id " + id + " is not found")
	}

	success := s.todoRepository.SetCompleted(id, completed)

	if !success {
		return nil, errors.New("todo with id " + id + " is not found")
	}

	return todo, nil
}

func (s *TodoService) Delete(id string) (*entity.Todo, error) {
	todo := s.todoRepository.Find(id)

	if todo == nil {
		return nil, errors.New("todo with id " + id + " is not found")
	}

	success := s.todoRepository.Delete(id)

	if !success {
		return nil, errors.New("todo with id " + id + " is not found")
	}

	return todo, nil
}

func (s TodoService) Close() {
	defer s.todoRepository.Close()
}
