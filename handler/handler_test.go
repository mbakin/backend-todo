package handler_test

import (
	"backend_todo/handler"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTodos(t *testing.T) {

	todoHandler := handler.NewHandlerTodo()

	t.Run("only GET method allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/todos", http.NoBody)
		res := httptest.NewRecorder()

		todoHandler.GetTodos(res, req)

		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})

	t.Run("only POST method allowed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/v1/todos", http.NoBody)
		res := httptest.NewRecorder()

		todoHandler.AddTodo(res, req)

		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
}
