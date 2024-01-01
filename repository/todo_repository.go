package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"todocible_api/database"
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

func (r *TodoRepository) Create(newTodo dto.TodoRequest) (entity.Todo, error) {
	todo := entity.Todo{
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
		return entity.Todo{}, err
	}

	return todo, nil
}

func (r *TodoRepository) FindAll() ([]entity.Todo, error) {
	ctx := context.Background()
	query := "SELECT id, title, description, due_date, completed FROM todos LIMIT 100"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []entity.Todo{}

	for rows.Next() {
		todo := entity.Todo{}

		rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.DueDate, &todo.Completed)

		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *TodoRepository) Find(id string) (entity.Todo, error) {
	ctx := context.Background()
	query := "SELECT id, title, description, due_date, completed FROM todos WHERE id = $1 LIMIT 1"

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		fmt.Println(err)
		return entity.Todo{}, database.ConnectionError
	}
	defer rows.Close()

	if rows.Next() {
		todo := entity.Todo{}

		rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.DueDate, &todo.Completed)
		return todo, nil
	}

	return entity.Todo{}, errors.New("todo with id " + id + " is not found")
}

func (r *TodoRepository) Update(id string, todo dto.TodoRequest) error {
	_, err := r.Find(id)
	if err != nil {
		return err
	}

	ctx := context.Background()
	query := "UPDATE todos SET title = $1, description = $2, due_date = $3 WHERE id = $4"

	_, err = r.db.ExecContext(ctx, query, todo.Title, todo.Description, todo.DueDate, id)

	return err
}

func (r *TodoRepository) SetCompleted(id string, completed bool) bool {
	ctx := context.Background()
	query := "UPDATE todos SET completed = $1 WHERE id = $2"

	_, err := r.db.ExecContext(ctx, query, completed, id)

	return err == nil
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
