package todo

import (
	"context"
	"todo-level-5/pkg/contract/todo"
	tRepo "todo-level-5/pkg/infrastructure/persistence/todo"

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

func (ts *TodoService) Create(ctx context.Context, tsr *todo.CreateTodoRequest) (*todo.CreateTodoResponse, error) {
	todoID := createNewTodoID()
	todo := FromSpaceTodoRequest(todoID, tsr)

	err := ts.tRepo.Create(ctx, todo)
	if err != nil {
		return nil, err
	}

	return ToCreateSpaceRes(todo), nil
}
