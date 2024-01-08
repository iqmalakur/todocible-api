package service

import (
	"errors"
	"time"
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

func (s *TodoService) Create(todo dto.TodoRequest) (entity.Task, error) {
	if todo.Title == "" {
		return entity.Task{}, errors.New("'title' is not allowed to be empty")
	}

	newTodo, err := s.todoRepository.Create(todo)
	if err != nil {
		return entity.Task{}, database.ConnectionError
	}

	return newTodo, nil
}

func (s *TodoService) GetAll() ([]entity.Task, error) {
	todos, err := s.todoRepository.FindAll()
	if err != nil {
		return nil, database.ConnectionError
	}

	return todos, nil
}

func (s *TodoService) Get(id string) (entity.Task, error) {
	todo, err := s.todoRepository.Find(id)
	if err != nil {
		return entity.Task{}, err
	}

	return todo, nil
}

func (s *TodoService) Update(id string, body dto.TodoRequest) (entity.Task, error) {
	todo, err := s.todoRepository.Find(id)
	if err != nil {
		return entity.Task{}, err
	}

	// Title validation
	if body.Title == "" {
		body.Title = todo.Name
	} else {
		todo.Name = body.Title
	}

	// Description validation
	if body.Description == "" {
		body.Description = todo.Description
	} else {
		todo.Description = body.Description
	}

	// Due date validation
	var emptyDate time.Time
	if body.DueDate == emptyDate {
		body.DueDate = todo.DueDate
	} else {
		todo.DueDate = body.DueDate
	}

	err = s.todoRepository.Update(id, body)
	if err != nil {
		return entity.Task{}, err
	}

	return todo, nil
}

func (s *TodoService) SetCompleted(id string, completed bool) (entity.Task, error) {
	todo, err := s.todoRepository.Find(id)
	if err != nil {
		return entity.Task{}, err
	}

	success := s.todoRepository.SetCompleted(id, completed)

	if !success {
		return entity.Task{}, errors.New("failed to update todo status")
	}

	return todo, nil
}

func (s *TodoService) Delete(id string) (entity.Task, error) {
	todo, err := s.todoRepository.Find(id)
	if err != nil {
		return entity.Task{}, errors.New("todo with id " + id + " is not found")
	}

	success := s.todoRepository.Delete(id)

	if !success {
		return entity.Task{}, errors.New("failed to delete todo")
	}

	return todo, nil
}

func (s TodoService) Close() {
	defer s.todoRepository.Close()
}
