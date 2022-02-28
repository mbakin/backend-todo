package handler

import (
	"backend_todo/model"
	"backend_todo/service"
	"encoding/json"
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

func (h *HandlerTodo) AddTodo(w http.ResponseWriter, r *http.Request) {
	todoDescription := model.Todo{}

	err := json.NewDecoder(r.Body).Decode(&todoDescription)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo := h.TodoService.AddTodo(todoDescription.Todo)

	jsonBytes, err := json.Marshal(todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBytes)
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
}

func (h *HandlerTodo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("content-type", "application/json; charset=utf-8")
	switch {
	case r.Method == http.MethodGet:
		h.GetTodos(w, r)
		return
	case r.Method == http.MethodPost:
		h.AddTodo(w, r)
		return
	default:
		return
	}

}

func NewHandlerTodo(todoService service.ITodoService) IHandlerTodo {
	return &HandlerTodo{TodoService: todoService}
}
