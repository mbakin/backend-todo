package service

import (
	"backend_todo/model"
)

type ITodoService interface {
	GetTodos() []*model.Todo
	AddTodo(todoDescription string) *model.Todo
}

type TodoService struct {
}

func NewTodoService() ITodoService {
	return &TodoService{}
}
func (todoService TodoService) GetTodos() []*model.Todo {
	return nil
}
func (todoService TodoService) AddTodo(newTodo string) *model.Todo {
	return nil
}