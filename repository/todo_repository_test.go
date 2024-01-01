package repository

import (
	"fmt"
	"testing"
	"time"
	"todocible_api/database"
	"todocible_api/dto"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var todoRepository TodoRepository

func TestMain(m *testing.M) {
	godotenv.Load("../.env")

	todoRepository = NewTodoRepository(database.GetConnection())

	m.Run()

	todoRepository.Close()
}

func TestCreate(t *testing.T) {
	todo, err := todoRepository.Create(dto.TodoRequest{
		Title:       "Coba",
		Description: "Hello World",
		DueDate:     time.Now(),
	})

	fmt.Println(todo.Id)
	fmt.Println(todo.Title)
	fmt.Println(todo.Description)
	fmt.Println(todo.DueDate)
	fmt.Println(todo.Completed)

	assert.Nil(t, err)
}

func TestFindAll(t *testing.T) {
	todos, err := todoRepository.FindAll()

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

func TestFindWithValidId(t *testing.T) {
	newTodo, err := todoRepository.Create(dto.TodoRequest{
		Title:       "Coba",
		Description: "Hello World",
		DueDate:     time.Now(),
	})

	assert.Nil(t, err)

	todo, err := todoRepository.Find(newTodo.Id)

	assert.Nil(t, err)

	fmt.Println("===================")
	fmt.Println(todo.Id)
	fmt.Println(todo.Title)
	fmt.Println(todo.Description)
	fmt.Println(todo.DueDate)
	fmt.Println(todo.Completed)
}

func TestFindWithInvalidId(t *testing.T) {
	_, err := todoRepository.Find(uuid.New().String())

	assert.NotNil(t, err)

	fmt.Println(err.Error())
}

func TestUpdate(t *testing.T) {
	newTodo, err := todoRepository.Create(dto.TodoRequest{
		Title:       "Coba",
		Description: "Hello World",
		DueDate:     time.Now(),
	})

	assert.Nil(t, err)
	assert.Equal(t, "Coba", newTodo.Title)

	err = todoRepository.Update(newTodo.Id, dto.TodoRequest{
		Title: "YO",
	})

	assert.Nil(t, err)

	todo, err := todoRepository.Find(newTodo.Id)

	assert.Nil(t, err)
	assert.Equal(t, "YO", todo.Title)
}

func TestUpdateWithInvalidId(t *testing.T) {
	err := todoRepository.Update(uuid.New().String(), dto.TodoRequest{
		Title: "YO",
	})

	assert.NotNil(t, err)
}

// func TestCompleted(t *testing.T) {
// 	todos := todoRepository.FindAll()
// 	todo := todoRepository.Find(todos[0].Id)
// 	expectedTodo := &entity.Todo{
// 		Title:       "Todo 1",
// 		Description: "Todolist 1",
// 		Completed:   false,
// 	}

// 	checkTodo(t, expectedTodo, todo)
// 	todoRepository.SetCompleted(todo.Id, true)

// 	todo = todoRepository.Find(todos[0].Id)
// 	expectedTodo = &entity.Todo{
// 		Title:       "Todo 1",
// 		Description: "Todolist 1",
// 		Completed:   true,
// 	}
// 	checkTodo(t, expectedTodo, todo)
// }

// func TestDelete(t *testing.T) {
// 	todos := todoRepository.FindAll()
// 	todo := todoRepository.Find(todos[0].Id)
// 	assert.NotNil(t, todo)

// 	todoRepository.Delete(todo.Id)

// 	todo = todoRepository.Find(todo.Id)
// 	assert.Nil(t, todo)
// }
