package service_test

import (
	"backend_todo/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TodoService(t *testing.T) {
	todoService := service.NewTodoService()
	getTodos := todoService.GetTodos()

	assert.Nil(t, getTodos)
}
