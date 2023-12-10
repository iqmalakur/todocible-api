package service

import (
	"testing"
	"todolist/dto"
	"todolist/repository"

	"github.com/stretchr/testify/assert"
)

var todoService = TodoService{&repository.TodoRepository{}}

func TestMain(m *testing.M) {
	todoService.Create(dto.TodoRequest{
		Title:       "Todo 1",
		Description: "Todolist 1",
	})
	todoService.Create(dto.TodoRequest{
		Title:       "Todo 2",
		Description: "Todolist 2",
	})
	todoService.Create(dto.TodoRequest{
		Title:       "Todo 3",
		Description: "Todolist 3",
	})

	m.Run()
}

func TestCreate(t *testing.T) {
	todo, _ := todoService.Create(dto.TodoRequest{
		Title:       "Coba",
		Description: "Hello World",
	})

	assert.Equal(t, "Coba", todo.Title)
	assert.Equal(t, "Hello World", todo.Description)
	assert.Equal(t, false, todo.Completed)
}

func TestGetAll(t *testing.T) {
	todos := todoService.GetAll()

	assert.Equal(t, "Todo 1", todos[0].Title)
	assert.Equal(t, "Todolist 1", todos[0].Description)
	assert.Equal(t, false, todos[0].Completed)
}

func TestGet(t *testing.T) {
	todo, _ := todoService.Get(0)

	assert.Equal(t, "Todo 1", todo.Title)
	assert.Equal(t, "Todolist 1", todo.Description)
	assert.Equal(t, false, todo.Completed)
}

func TestUpdate(t *testing.T) {
	todo, _ := todoService.Get(0)
	assert.Equal(t, "Todo 1", todo.Title)
	assert.Equal(t, "Todolist 1", todo.Description)
	assert.Equal(t, false, todo.Completed)

	todo.Title = "Hello"
	todo.Description = "Hello World"
	todoService.Update(0, dto.TodoRequest{
		Title:       todo.Title,
		Description: todo.Description,
	})

	todo, _ = todoService.Get(0)
	assert.Equal(t, "Hello", todo.Title)
	assert.Equal(t, "Hello World", todo.Description)
	assert.Equal(t, false, todo.Completed)
}

func TestCompleted(t *testing.T) {
	todo, _ := todoService.Get(0)
	assert.Equal(t, false, todo.Completed)

	todoService.SetCompleted(0, true)

	todo, _ = todoService.Get(0)
	assert.Equal(t, true, todo.Completed)
}

func TestDelete(t *testing.T) {
	todo, _ := todoService.Get(0)
	assert.NotNil(t, todo)

	todoService.Delete(0)

	todo, _ = todoService.Get(0)
	assert.Nil(t, todo)
}
