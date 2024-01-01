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

func (s *TodoService) Create(todo dto.TodoRequest) (entity.Todo, error) {
	if todo.Title == "" {
		return entity.Todo{}, errors.New("'title' is not allowed to be empty")
	}

	newTodo, err := s.todoRepository.Create(todo)
	if err != nil {
		return entity.Todo{}, database.ConnectionError
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

func (s *TodoService) Get(id string) (entity.Todo, error) {
	todo, err := s.todoRepository.Find(id)
	if err != nil {
		return entity.Todo{}, err
	}

	return todo, nil
}

func (s *TodoService) Update(id string, body dto.TodoRequest) (entity.Todo, error) {
	todo, err := s.todoRepository.Find(id)
	if err != nil {
		return entity.Todo{}, err
	}

	// Title validation
	if body.Title == "" {
		body.Title = todo.Title
	} else {
		todo.Title = body.Title
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
		return entity.Todo{}, err
	}

	return todo, nil
}

func (s *TodoService) SetCompleted(id string, completed bool) (entity.Todo, error) {
	todo, err := s.todoRepository.Find(id)
	if err != nil {
		return entity.Todo{}, err
	}

	success := s.todoRepository.SetCompleted(id, completed)

	if !success {
		return entity.Todo{}, errors.New("failed to update todo status")
	}

	return todo, nil
}

func (s *TodoService) Delete(id string) (entity.Todo, error) {
	todo, _ := s.todoRepository.Find(id)

	// if todo == nil {
	// 	return nil, errors.New("todo with id " + id + " is not found")
	// }

	// success := s.todoRepository.Delete(id)

	// if !success {
	// 	return nil, errors.New("todo with id " + id + " is not found")
	// }

	return todo, nil
}

func (s TodoService) Close() {
	defer s.todoRepository.Close()
}
