package todo

import (
	"context"
	"todo-level-5/pkg/contract/todo"
	tRepo "todo-level-5/pkg/infrastructure/persistence/todo"

	"go.mongodb.org/mongo-driver/mongo"
)

type TodoService struct {
	client *mongo.Client
}

func NewTodoService(client *mongo.Client) *TodoService {
	return &TodoService{
		client: client,
	}
}

func collections(client *mongo.Client) *mongo.Collection {
	return client.Database("todoDB").Collection("todos")
}

func (ts *TodoService) Create(ctx context.Context, todo todo.CreateTodoRequest) (*todo.CreateTodoResponse, error) {
	todoM := tRepo.ToSpaceModel(todo)

	_, err := collections(ts.client).InsertOne(context.Background(), todoM)
	if err != nil {
		return nil, err
	}

	return ToCreateSpaceRes(todoM), nil
}
