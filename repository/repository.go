package repository

import "backend_todo/model"

type IRepositoryTodo interface {
	GetTodos() map[string]*model.Todo
	AddTodo(todo model.Todo) *model.Todo
}

type TodoRepository struct {
}

func (t *TodoRepository) GetTodos() map[string]*model.Todo {
	return map[string]*model.Todo{}
}

func (t *TodoRepository) AddTodo(todo model.Todo) *model.Todo {
	return nil
}

func NewRepositoryTodo() IRepositoryTodo {
	return &TodoRepository{}
}
