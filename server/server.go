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

func (s *Server) StartServer(port int) error {
	todoRepository := repository.NewRepositoryTodo()
	todoService := service.NewTodoService(todoRepository)
	todoHandler := handler.NewHandlerTodo(todoService)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/todos", todoHandler.ServeHTTP)

	fmt.Println("Server is working")
	fmt.Println("Server is running on port:", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	return err
}
