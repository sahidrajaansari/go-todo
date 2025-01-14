package services

import (
	"context"
	tService "todo-level-5/pkg/application/todo"
	tContracts "todo-level-5/pkg/contract/todo"

	"github.com/gin-gonic/gin"
)

var _ ITodoService = (*tService.TodoService)(nil)

type ITodoService interface {
	Create(ctx context.Context, tsr *tContracts.CreateTodoRequest) (*tContracts.CreateTodoResponse, error)
	GetTodoByID(ctx context.Context) (*tContracts.GetTodoResponse, error)
	GetTodos(ctx *gin.Context) ([]tContracts.GetTodoResponse, error)
	UpdateTodoByID(ctx context.Context, tsr *tContracts.UpdateTodoRequest) (tContracts.UpdateTodoResponse, error)
	DeleteTodo(ctx context.Context) error
}
