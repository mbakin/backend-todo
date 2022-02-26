package main

import (
	"backend_todo/handler"
	"fmt"
	"net/http"
)

func main() {
	todoHandler := handler.NewHandlerTodo()
	http.HandleFunc("/api/v1/todos", todoHandler.GetTodos)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
