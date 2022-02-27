package service

import (
	"backend_todo/model"
	"backend_todo/repository"
)

type ITodoService interface {
	GetTodos() []*model.Todo
	AddTodo(todoDescription string) *model.Todo
}

type TodoService struct {
	todoRepository repository.IRepositoryTodo
}

func NewTodoService(todoRepository repository.IRepositoryTodo) ITodoService {
	return &TodoService{todoRepository: todoRepository}
}

func (todoService TodoService) GetTodos() []*model.Todo {
}

func (todoService TodoService) AddTodo(newTodo string) *model.Todo {
	return nil
}
