package repository

import (
	"todolist/dto"
	"todolist/entity"

	"github.com/google/uuid"
)

type TodoRepository struct {
	Todo []*entity.Todo
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{[]*entity.Todo{}}
}

func (todoRepository *TodoRepository) Create(todo dto.TodoRequest) *entity.Todo {
	newTodo := &entity.Todo{
		Id:          uuid.New().String(),
		Title:       todo.Title,
		Description: todo.Description,
		DueDate:     todo.DueDate,
		Completed:   false,
	}

	todoRepository.Todo = append(todoRepository.Todo, newTodo)

	return newTodo
}

func (todoRepository *TodoRepository) FindAll() []*entity.Todo {
	return todoRepository.Todo
}

func (todoRepository *TodoRepository) Find(id string) *entity.Todo {
	for _, todo := range todoRepository.Todo {
		if todo.Id == id {
			return todo
		}
	}

	return nil
}

func (todoRepository *TodoRepository) Update(id string, newTodo *entity.Todo) *entity.Todo {
	todo := todoRepository.Find(id)

	if todo == nil {
		return nil
	}

	todo.Title = newTodo.Title
	todo.Description = newTodo.Description

	return todo
}

func (todoRepository *TodoRepository) SetCompleted(id string, completed bool) bool {
	todo := todoRepository.Find(id)

	if todo == nil {
		return false
	}

	todo.Completed = completed

	return true
}

func (todoRepository *TodoRepository) Delete(id string) bool {
	index := -1

	for i, todo := range todoRepository.Todo {
		if todo.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		return false
	}

	todoRepository.Todo = append(todoRepository.Todo[:index], todoRepository.Todo[index+1:]...)

	return true
}
