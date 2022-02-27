package service_test

import (
	"backend_todo/mocks"
	"backend_todo/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TodoService(t *testing.T) {

	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockTodoRepository := mocks.NewMockIRepositoryTodo(mockController)

	todoService := service.NewTodoService(mockTodoRepository)
	getTodos := todoService.GetTodos()

	assert.Nil(t, getTodos)
}
