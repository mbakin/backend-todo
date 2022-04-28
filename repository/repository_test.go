package repository_test

import (
	"backend_todo/model"
	"backend_todo/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_GetAll(t *testing.T) {

	repository := repository.NewRepositoryTodo()

	lengthOfWallet := len(repository.GetTodos())
	assert.Equal(t, 0, lengthOfWallet)
}

func TestRepository_AddTodo(t *testing.T) {
	todo := model.Todo{
		ID:   11,
		Todo: "test",
	}
	repository := repository.NewRepositoryTodo()

	initialLengthOfTodo := len(repository.GetTodos())
	repository.AddTodo(todo)
	assert.Equal(t, initialLengthOfTodo+1, len(repository.GetTodos()))
}

func TestRpository_DeleteAll(t *testing.T) {

	repository := repository.NewRepositoryTodo()

	repository.GetTodos()
	initialLengthOfTodo := len(repository.GetTodos())
	repository.DeleteAllTodos()
	assert.Equal(t, 0, initialLengthOfTodo)
}
