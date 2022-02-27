package service_test

import (
	"backend_todo/mocks"
	"backend_todo/model"
	"backend_todo/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TodoService(t *testing.T) {
	repositoryReturn := map[string]*model.Todo{
		"Dummy": &model.Todo{
			ID:   1000,
			Todo: "Dummy",
		},
	}
	expectedRepository := make([]*model.Todo, 0)
	expectedRepository = append(expectedRepository, &model.Todo{
		ID:   1000,
		Todo: "Dummy",
	})

	repository := mocks.NewMockIRepositoryTodo(gomock.NewController(t))
	repository.EXPECT().GetTodos().Return(repositoryReturn).Times(1)

	service := service.NewTodoService(repository)
	actualRepository := service.GetTodos()

	assert.Equal(t, &expectedRepository, &actualRepository)
}
