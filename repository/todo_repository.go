package repository

import (
	"context"
	"todocible_api/database"
	"todocible_api/dto"
	"todocible_api/entity"

	"github.com/google/uuid"
)

type TodoRepository struct {
	Todo []*entity.Todo
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{[]*entity.Todo{}}
}

func (repo *TodoRepository) Create(newTodo dto.TodoRequest) (*entity.Todo, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	todo := &entity.Todo{
		Id:          uuid.New().String(),
		Title:       newTodo.Title,
		Description: newTodo.Description,
		DueDate:     newTodo.DueDate,
		Completed:   false,
	}

	ctx := context.Background()
	query := "INSERT INTO todos (id, title, description, due_date, completed) VALUES ($1, $2, $3, $4, $5)"

	_, err = db.ExecContext(ctx, query, todo.Id, todo.Title, todo.Description, todo.DueDate, todo.Completed)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (repo *TodoRepository) FindAll() []*entity.Todo {
	return repo.Todo
}

func (repo *TodoRepository) Find(id string) *entity.Todo {
	for _, todo := range repo.Todo {
		if todo.Id == id {
			return todo
		}
	}

	return nil
}

func (repo *TodoRepository) Update(id string, newTodo dto.TodoRequest) *entity.Todo {
	todo := repo.Find(id)

	if todo == nil {
		return nil
	}

	todo.Title = newTodo.Title
	todo.Description = newTodo.Description
	todo.DueDate = newTodo.DueDate

	return todo
}

func (repo *TodoRepository) SetCompleted(id string, completed bool) bool {
	todo := repo.Find(id)

	if todo == nil {
		return false
	}

	todo.Completed = completed

	return true
}

func (repo *TodoRepository) Delete(id string) bool {
	index := -1

	for i, todo := range repo.Todo {
		if todo.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		return false
	}

	repo.Todo = append(repo.Todo[:index], repo.Todo[index+1:]...)

	return true
}
