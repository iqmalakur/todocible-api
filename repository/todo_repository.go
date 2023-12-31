package repository

import (
	"context"
	"database/sql"
	"todocible_api/dto"
	"todocible_api/entity"

	"github.com/google/uuid"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return TodoRepository{db}
}

func (r *TodoRepository) Create(newTodo dto.TodoRequest) (*entity.Todo, error) {
	todo := &entity.Todo{
		Id:          uuid.New().String(),
		Title:       newTodo.Title,
		Description: newTodo.Description,
		DueDate:     newTodo.DueDate,
		Completed:   false,
	}

	ctx := context.Background()
	query := "INSERT INTO todos (id, title, description, due_date, completed) VALUES ($1, $2, $3, $4, $5)"

	_, err := r.db.ExecContext(ctx, query, todo.Id, todo.Title, todo.Description, todo.DueDate, todo.Completed)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *TodoRepository) FindAll() []*entity.Todo {
	return []*entity.Todo{}
}

func (r *TodoRepository) Find(id string) *entity.Todo {
	// for _, todo := range r.Todo {
	// 	if todo.Id == id {
	// 		return todo
	// 	}
	// }

	return nil
}

func (r *TodoRepository) Update(id string, newTodo dto.TodoRequest) *entity.Todo {
	todo := r.Find(id)

	if todo == nil {
		return nil
	}

	todo.Title = newTodo.Title
	todo.Description = newTodo.Description
	todo.DueDate = newTodo.DueDate

	return todo
}

func (r *TodoRepository) SetCompleted(id string, completed bool) bool {
	todo := r.Find(id)

	if todo == nil {
		return false
	}

	todo.Completed = completed

	return true
}

func (r *TodoRepository) Delete(id string) bool {
	// index := -1

	// for i, todo := range r.Todo {
	// 	if todo.Id == id {
	// 		index = i
	// 		break
	// 	}
	// }

	// if index == -1 {
	// 	return false
	// }

	// r.Todo = append(r.Todo[:index], r.Todo[index+1:]...)

	return true
}

func (r *TodoRepository) Close() {
	defer r.db.Close()
}
