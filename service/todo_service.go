package service

import (
	"todolist/entity"
	"todolist/repository"
)

type TodoService struct {
	todoRepository *repository.TodoRepository
}

func (todoService *TodoService) Create(title, description string) *entity.Todo {
	todo := todoService.todoRepository.Create(title, description)
	return todo
}

func (todoService *TodoService) GetAll() []*entity.Todo {
	todos := todoService.todoRepository.FindAll()
	return todos
}

func (todoService *TodoService) Get(id int) *entity.Todo {
	todo := todoService.todoRepository.Find(id)
	return todo
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
