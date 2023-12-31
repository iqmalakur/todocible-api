package service

import (
	"fmt"
	"testing"
	"todocible_api/dto"

	"github.com/stretchr/testify/assert"
)

var todoService TodoService

func TestMain(m *testing.M) {
	todoService = NewTodoService()

	m.Run()

	todoService.Close()
}

func TestSuccessCreate(t *testing.T) {
	todo, err := todoService.Create(dto.TodoRequest{
		Title:       "Coba",
		Description: "Hello World",
	})

	fmt.Println(todo.Id)
	fmt.Println(todo.Title)
	fmt.Println(todo.Description)
	fmt.Println(todo.DueDate)
	fmt.Println(todo.Completed)

	assert.Nil(t, err)
}

func TestCreateWithoutTitle(t *testing.T) {
	_, err := todoService.Create(dto.TodoRequest{})

	assert.NotNil(t, err)
}

func TestGetAll(t *testing.T) {
	todos, err := todoService.GetAll()

	for _, todo := range todos {
		fmt.Println("===================")
		fmt.Println(todo.Id)
		fmt.Println(todo.Title)
		fmt.Println(todo.Description)
		fmt.Println(todo.DueDate)
		fmt.Println(todo.Completed)
	}

	assert.Nil(t, err)
}

// func TestGet(t *testing.T) {
// 	todos := todoService.GetAll()
// 	todo, _ := todoService.Get(todos[0].Id)

// 	assert.Equal(t, "Todo 1", todo.Title)
// 	assert.Equal(t, "Todolist 1", todo.Description)
// 	assert.Equal(t, false, todo.Completed)
// }

// func TestUpdate(t *testing.T) {
// 	todos := todoService.GetAll()
// 	todo, _ := todoService.Get(todos[0].Id)
// 	assert.Equal(t, "Todo 1", todo.Title)
// 	assert.Equal(t, "Todolist 1", todo.Description)
// 	assert.Equal(t, false, todo.Completed)

// 	todo.Title = "Hello"
// 	todo.Description = "Hello World"
// 	todoService.Update(todo.Id, dto.TodoRequest{
// 		Title:       todo.Title,
// 		Description: todo.Description,
// 	})

// 	todo, _ = todoService.Get(todo.Id)
// 	assert.Equal(t, "Hello", todo.Title)
// 	assert.Equal(t, "Hello World", todo.Description)
// 	assert.Equal(t, false, todo.Completed)
// }

// func TestCompleted(t *testing.T) {
// 	todos := todoService.GetAll()
// 	todo, _ := todoService.Get(todos[0].Id)
// 	assert.Equal(t, false, todo.Completed)

// 	todoService.SetCompleted(todo.Id, true)

// 	todo, _ = todoService.Get(todo.Id)
// 	assert.Equal(t, true, todo.Completed)
// }

// func TestDelete(t *testing.T) {
// 	todos := todoService.GetAll()
// 	todo, _ := todoService.Get(todos[0].Id)
// 	assert.NotNil(t, todo)

// 	todoService.Delete(todo.Id)

// 	todo, _ = todoService.Get(todo.Id)
// 	assert.Nil(t, todo)
// }
