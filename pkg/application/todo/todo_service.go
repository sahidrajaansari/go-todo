package todo

import (
	"context"
	tContracts "todo-level-5/pkg/contract/todo"
	tRepo "todo-level-5/pkg/infrastructure/persistence/todo"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

type TodoService struct {
	tRepo *tRepo.TodoRepo
}

func NewTodoService(todoRepo *tRepo.TodoRepo) *TodoService {
	return &TodoService{
		tRepo: todoRepo,
	}
}

func createNewTodoID() string {
	return ksuid.New().String()
}

func (ts *TodoService) Create(ctx context.Context, tsr *tContracts.CreateTodoRequest) (*tContracts.CreateTodoResponse, error) {
	todoID := createNewTodoID()
	todo := FromSpaceTodoRequest(todoID, tsr)

	err := ts.tRepo.Create(ctx, todo)
	if err != nil {
		return nil, err
	}

	return ToCreateSpaceRes(todo), nil
}

func (ts *TodoService) GetTodoByID(ctx *gin.Context) (*tContracts.GetTodoResponse, error) {
	todoID := ctx.Param(":id")

	todo, err := ts.tRepo.GetTodoByID(ctx, todoID)
	if err != nil {
		return nil, err
	}

	return ToGetByIDRes(todo), nil
}
