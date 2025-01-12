package todo

import (
	"context"
	"log"
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
	todo := fromCreateTodoRequest(todoID, tsr)

	err := ts.tRepo.Create(ctx, todo)
	if err != nil {
		return nil, err
	}

	return ToCreateTodoRes(todo), nil
}

func (ts *TodoService) GetTodoByID(ctx *gin.Context) (*tContracts.GetTodoResponse, error) {
	todoID := ctx.Param("id")

	todo, err := ts.tRepo.GetTodoByID(ctx, todoID)
	if err != nil {
		return nil, err
	}

	return ToGetByIDRes(todo), nil
}

func (ts *TodoService) GetTodos(ctx *gin.Context) ([]*tContracts.GetTodoResponse, error) {
	query := ctx.Request.URL.RawQuery
	log.Println(query)

	todos, err := ts.tRepo.GetTodos(ctx, query)
	if err != nil {
		return nil, err
	}
	var allTodos []*tContracts.GetTodoResponse

	for _, todo := range todos {
		allTodos = append(allTodos, ToGetByIDRes(todo))
	}

	return allTodos, nil
}

func (ts *TodoService) UpdateTodoByID(ctx *gin.Context, tsr *tContracts.UpdateTodoRequest) (*tContracts.UpdateTodoResponse, error) {
	todoID := ctx.Param("id")
	todoAgg := fromUpdateTodoRequest(todoID, tsr)

	todo, err := ts.tRepo.UpdateTodo(ctx, todoID, todoAgg)
	if err != nil {
		return nil, err
	}
	return toUpateTodoRes(todo), nil
}

func (ts *TodoService) DeleteTodo(ctx *gin.Context) error {
	todoID := ctx.Param("id")

	err := ts.tRepo.DeleteTodo(ctx, todoID)
	if err != nil {
		return err
	}

	return nil
}
