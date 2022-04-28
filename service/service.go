package service

import (
	"backend_todo/model"
	"backend_todo/repository"
	"fmt"
)

type ITodoService interface {
	GetTodos() []*model.Todo
	AddTodo(todoDescription string) *model.Todo
	DeleteAllTodos() []*model.Todo
}

type TodoService struct {
	todoRepository repository.IRepositoryTodo
}

func NewTodoService(todoRepository repository.IRepositoryTodo) ITodoService {
	return &TodoService{todoRepository: todoRepository}
}

func (todoService TodoService) GetTodos() []*model.Todo {
	getTodos := todoService.todoRepository.GetTodos()
	getTodoList := make([]*model.Todo, 0)

	for _, getTodo := range getTodos {
		getTodoList = append(getTodoList, getTodo)
	}

	return getTodoList
}

func (todoService TodoService) AddTodo(newTodo string) *model.Todo {
	AllTodos := todoService.todoRepository.GetTodos()
	TodoByDescription, existTodo := AllTodos[newTodo]
	if existTodo == true {
		return TodoByDescription
	}
	todoId := len(AllTodos) + 1
	todo := model.Todo{
		ID:   todoId,
		Todo: newTodo,
	}
	return todoService.todoRepository.AddTodo(todo)
}

func (todoService TodoService) DeleteAllTodos() []*model.Todo {

	getTodos := todoService.todoRepository.GetTodos()

	fmt.Println("test:", getTodos)

	deleteAllTodos := todoService.todoRepository.DeleteAllTodos()
	getTodoList := make([]*model.Todo, 0)

	for _, getTodo := range deleteAllTodos {
		getTodoList = append(getTodoList, getTodo)
	}

	return getTodoList

}
