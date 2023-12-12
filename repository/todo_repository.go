package repository

import "todolist/entity"

type TodoRepository struct {
	Todo    []*entity.Todo
	counter int
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{Todo: []*entity.Todo{}, counter: 0}
}

func (todoRepository *TodoRepository) Create(title, description string) *entity.Todo {
	todo := &entity.Todo{
		Id:          todoRepository.counter,
		Title:       title,
		Description: description,
		Completed:   false,
	}

	todoRepository.Todo = append(todoRepository.Todo, todo)
	todoRepository.counter++

	return todo
}

func (todoRepository *TodoRepository) FindAll() []*entity.Todo {
	return todoRepository.Todo
}

func (todoRepository *TodoRepository) Find(id int) *entity.Todo {
	for _, todo := range todoRepository.Todo {
		if todo.Id == id {
			return todo
		}
	}

	return nil
}

func (todoRepository *TodoRepository) Update(id int, newTodo *entity.Todo) *entity.Todo {
	todo := todoRepository.Find(id)

	if todo == nil {
		return nil
	}

	todo.Title = newTodo.Title
	todo.Description = newTodo.Description

	return todo
}

func (todoRepository *TodoRepository) SetCompleted(id int, completed bool) bool {
	todo := todoRepository.Find(id)

	if todo == nil {
		return false
	}

	todo.Completed = completed

	return true
}

func (todoRepository *TodoRepository) Delete(id int) bool {
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
