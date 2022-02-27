package handler_test

import (
	"backend_todo/handler"
	"backend_todo/mocks"
	"backend_todo/model"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Handler_GetTodos(t *testing.T) {
	expectedRepository := make([]*model.Todo, 0)
	expectedRepository = append(expectedRepository, &model.Todo{
		ID:   1000,
		Todo: "Dummy",
	})

	service := mocks.NewMockITodoService(gomock.NewController(t))
	service.EXPECT().GetTodos().Return(expectedRepository).Times(1)

	handler := handler.NewHandlerTodo(service)
	r := httptest.NewRequest(http.MethodGet, "/api/v1/todos", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	actual := w.Body.String()
	json.Unmarshal(w.Body.Bytes(), &actual)

	assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("content-type"))
}
