package repository

import (
	"testing"
	"time"
	"todocible_api/database"
	"todocible_api/dto"
	"todocible_api/entity"

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

func checkTodo(t *testing.T, expected *entity.Todo, actual *entity.Todo) {
	assert.Equal(t, expected.Title, actual.Title)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Completed, actual.Completed)
}

func TestCreate(t *testing.T) {
	_, err := todoRepository.Create(dto.TodoRequest{
		Title:       "Coba",
		Description: "Hello World",
		DueDate:     time.Now(),
	})

	assert.Nil(t, err)
}

func TestFindAll(t *testing.T) {
	todos := todoRepository.FindAll()

	expectedTodo := &entity.Todo{
		Title:       "Todo 1",
		Description: "Todolist 1",
		Completed:   false,
	}

	assert.NotEqual(t, 0, len(todos))
	checkTodo(t, expectedTodo, todos[0])
}

func TestFind(t *testing.T) {
	todos := todoRepository.FindAll()
	todo := todoRepository.Find(todos[1].Id)

	expectedTodo := &entity.Todo{
		Title:       "Todo 2",
		Description: "Todolist 2",
		Completed:   false,
	}

	checkTodo(t, expectedTodo, todo)
}

func TestUpdate(t *testing.T) {
	todos := todoRepository.FindAll()
	todo := todoRepository.Find(todos[1].Id)
	expectedTodo := &entity.Todo{
		Title:       "Todo 2",
		Description: "Todolist 2",
		Completed:   false,
	}

	checkTodo(t, expectedTodo, todo)
	todoRepository.Update(todo.Id, dto.TodoRequest{Title: "Hello", Description: "World"})

	todo = todoRepository.Find(todos[1].Id)
	expectedTodo = &entity.Todo{
		Title:       "Hello",
		Description: "World",
		Completed:   false,
	}
	checkTodo(t, expectedTodo, todo)
}

func TestCompleted(t *testing.T) {
	todos := todoRepository.FindAll()
	todo := todoRepository.Find(todos[0].Id)
	expectedTodo := &entity.Todo{
		Title:       "Todo 1",
		Description: "Todolist 1",
		Completed:   false,
	}

	checkTodo(t, expectedTodo, todo)
	todoRepository.SetCompleted(todo.Id, true)

	todo = todoRepository.Find(todos[0].Id)
	expectedTodo = &entity.Todo{
		Title:       "Todo 1",
		Description: "Todolist 1",
		Completed:   true,
	}
	checkTodo(t, expectedTodo, todo)
}

func TestDelete(t *testing.T) {
	todos := todoRepository.FindAll()
	todo := todoRepository.Find(todos[0].Id)
	assert.NotNil(t, todo)

	todoRepository.Delete(todo.Id)

	todo = todoRepository.Find(todo.Id)
	assert.Nil(t, todo)
}
