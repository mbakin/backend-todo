package repository

import "backend_todo/model"

type IRepositoryTodo interface {
	GetTodos() map[string]*model.Todo
	AddTodo(todo model.Todo) *model.Todo
}

type TodoRepository struct {
	Todo map[string]*model.Todo
}

func (r *TodoRepository) GetTodos() map[string]*model.Todo {
	allList := r.Todo
	return allList
}

func (r *TodoRepository) AddTodo(todo model.Todo) *model.Todo {
	r.Todo[todo.Todo] = &model.Todo{
		ID:   todo.ID,
		Todo: todo.Todo,
	}

	return r.Todo[todo.Todo]
}

func NewRepositoryTodo() IRepositoryTodo {
	return &TodoRepository{Todo: map[string]*model.Todo{}}
}
