package todo

import (
	"context"
	"fmt"
	tContracts "todo-level-5/pkg/contract/todo"
	iPersist "todo-level-5/pkg/domain/persistence"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

type TodoService struct {
	tRepo iPersist.ITodoRepo
}

func NewTodoService(tRepo iPersist.ITodoRepo) *TodoService {
	return &TodoService{
		tRepo: tRepo,
	}
}

func createNewTodoID() string {
	return ksuid.New().String()
}

func (ts *TodoService) Create(ctx context.Context, tsr *tContracts.CreateTodoRequest) (*tContracts.CreateTodoResponse, error) {
	todoID := createNewTodoID()
	todo := fromCreateTodoRequest(todoID, tsr)

	err := ts.tRepo.Create(ctx, todo)
	if err != nil {
		return nil, err
	}

	return ToCreateTodoRes(todo), nil
}

func (ts *TodoService) GetTodoByID(ctx context.Context) (*tContracts.GetTodoResponse, error) {
	todoID, ok := ctx.Value("todoID").(string)
	if !ok {
		return nil, fmt.Errorf("todoID not found in context")
	}

	todo, err := ts.tRepo.GetTodoByID(ctx, todoID)
	if err != nil {
		return nil, err
	}

	return ToGetByIDRes(todo), nil
}

func (ts *TodoService) GetTodos(ctx *gin.Context) ([]tContracts.GetTodoResponse, error) {
	query := ctx.Request.URL.RawQuery

	todos, err := ts.tRepo.GetTodos(ctx, query)
	if err != nil {
		return nil, err
	}
	var allTodos []tContracts.GetTodoResponse

	for _, todo := range todos {
		allTodos = append(allTodos, (*ToGetByIDRes(todo)))
	}

	return allTodos, nil
}

func (ts *TodoService) UpdateTodoByID(ctx context.Context, tsr *tContracts.UpdateTodoRequest) (tContracts.UpdateTodoResponse, error) {
	todoID, ok := ctx.Value("todoID").(string)
	if !ok {
		return tContracts.UpdateTodoResponse{}, fmt.Errorf("todoID not found in context")
	}

	todoAgg := fromUpdateTodoRequest(todoID, tsr)

	todo, err := ts.tRepo.UpdateTodo(ctx, todoID, todoAgg)
	if err != nil {
		return tContracts.UpdateTodoResponse{}, err
	}
	return (*toUpateTodoRes(todo)), nil
}

func (ts *TodoService) DeleteTodo(ctx context.Context) error {
	todoID, ok := ctx.Value("todoID").(string)
	if !ok {
		return fmt.Errorf("todoID not found in context")
	}

	err := ts.tRepo.DeleteTodo(ctx, todoID)
	if err != nil {
		return err
	}

	return nil
}
