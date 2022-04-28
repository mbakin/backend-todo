package server

import (
	"backend_todo/handler"
	"backend_todo/repository"
	"backend_todo/service"
	"fmt"
	"net/http"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) StartServer(port string) error {
	todoRepository := repository.NewRepositoryTodo()
	todoService := service.NewTodoService(todoRepository)
	todoHandler := handler.NewHandlerTodo(todoService)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/todos", todoHandler.ServeHTTP)
	mux.HandleFunc("/api/v1/todos/deleteAll", todoHandler.DeleteAllHandler)

	fmt.Println("Server is working")
	fmt.Println("Server is running on port:", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	return err
}
