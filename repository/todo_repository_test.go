package repository

import (
	"testing"
	"todolist/entity"

	"github.com/stretchr/testify/assert"
)

var todoRepository TodoRepository

func TestMain(m *testing.M) {
	todoRepository.Create("Todo 1", "Todolist 1")
	todoRepository.Create("Todo 2", "Todolist 2")
	todoRepository.Create("Todo 3", "Todolist 3")
	todoRepository.Create("Todo 4", "Todolist 4")
	todoRepository.Create("Todo 5", "Todolist 5")

	m.Run()
}

func checkTodo(t *testing.T, expected *entity.Todo, actual *entity.Todo) {
	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.Title, actual.Title)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Completed, actual.Completed)
}

func TestCreate(t *testing.T) {
	success := todoRepository.Create("Coba", "Hello World")

	if success {
		todo := todoRepository.Find(len(todoRepository.Todo) - 1)

		expectedTodo := &entity.Todo{
			Id:          len(todoRepository.Todo) - 1,
			Title:       "Coba",
			Description: "Hello World",
			Completed:   false,
		}

		checkTodo(t, expectedTodo, todo)
	}
}

func TestFindAll(t *testing.T) {
	todos := todoRepository.FindAll()

	expectedTodo := &entity.Todo{
		Id:          0,
		Title:       "Todo 1",
		Description: "Todolist 1",
		Completed:   false,
	}

	assert.NotEqual(t, 0, len(todos))
	checkTodo(t, expectedTodo, todos[0])
}

func TestFind(t *testing.T) {
	todo := todoRepository.Find(1)

	expectedTodo := &entity.Todo{
		Id:          1,
		Title:       "Todo 2",
		Description: "Todolist 2",
		Completed:   false,
	}

	checkTodo(t, expectedTodo, todo)
}

func TestUpdate(t *testing.T) {
	todo := todoRepository.Find(1)
	expectedTodo := &entity.Todo{
		Id:          1,
		Title:       "Todo 2",
		Description: "Todolist 2",
		Completed:   false,
	}

	checkTodo(t, expectedTodo, todo)
	todoRepository.Update(1, entity.Todo{Title: "Hello", Description: "World"})

	todo = todoRepository.Find(1)
	expectedTodo = &entity.Todo{
		Id:          1,
		Title:       "Hello",
		Description: "World",
		Completed:   false,
	}
	checkTodo(t, expectedTodo, todo)
}

func TestCompleted(t *testing.T) {
	todo := todoRepository.Find(0)
	expectedTodo := &entity.Todo{
		Id:          0,
		Title:       "Todo 1",
		Description: "Todolist 1",
		Completed:   false,
	}

	checkTodo(t, expectedTodo, todo)
	todoRepository.SetCompleted(0, true)

	todo = todoRepository.Find(0)
	expectedTodo = &entity.Todo{
		Id:          0,
		Title:       "Todo 1",
		Description: "Todolist 1",
		Completed:   true,
	}
	checkTodo(t, expectedTodo, todo)
}

func TestDelete(t *testing.T) {
	todo := todoRepository.Find(0)
	assert.NotNil(t, todo)

	todoRepository.Delete(0)

	todo = todoRepository.Find(0)
	assert.Nil(t, todo)
}
