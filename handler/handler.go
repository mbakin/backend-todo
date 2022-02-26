package handler

import (
	"fmt"
	"net/http"
)

type IHandlerTodo interface {
	AddTodo(w http.ResponseWriter, r *http.Request)
	GetTodos(w http.ResponseWriter, r *http.Request)
}

type HandlerTodo struct{}

func (h HandlerTodo) AddTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddTodo")
}

func (h HandlerTodo) GetTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func NewHandlerTodo() IHandlerTodo {
	return &HandlerTodo{}
}
