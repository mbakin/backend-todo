package handler

import (
	"backend_todo/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type IHandlerTodo interface {
	AddTodo(w http.ResponseWriter, r *http.Request)
	GetTodos(w http.ResponseWriter, r *http.Request)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type HandlerTodo struct {
	TodoService service.ITodoService
}

func (h HandlerTodo) AddTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddTodo")
}

func (h *HandlerTodo) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := h.TodoService.GetTodos()

	jsonBytes, err := json.Marshal(todos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
	fmt.Println("get request received")
}

func (h *HandlerTodo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	switch {
	case r.Method == http.MethodGet:
		h.GetTodos(w, r)
		return
	case r.Method == http.MethodPost:
		h.AddTodo(w, r)
		fmt.Println("post request received")
		return
	case r.Method == http.MethodDelete:
		fmt.Println("delete request received ")
		return
	default:
		fmt.Println("Bad request")
		return
	}

	fmt.Println("request received")
}

func NewHandlerTodo(todoService service.ITodoService) IHandlerTodo {
	return &HandlerTodo{TodoService: todoService}
}
