package repository_test

import (
	"backend_todo/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_GetAll(t *testing.T) {

	repository := repository.NewRepositoryTodo()

	lengthOfWallet := len(repository.GetTodos())
	assert.Equal(t, 0, lengthOfWallet)
}
