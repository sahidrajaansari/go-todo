package todo

import (
	"context"
	"log"
	todoagg "todo-level-5/pkg/domain/todo_aggregate"

	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepo struct {
	client *mongo.Client
}

func NewTodoRepo(client *mongo.Client) *TodoRepo {
	return &TodoRepo{
		client: client,
	}
}

func todoCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("todoDB").Collection("todos")
}

func (tr *TodoRepo) Create(ctx context.Context, todoAgg *todoagg.Todo) error {
	todo := ToSpaceModel(todoAgg)

	_, err := todoCollection(tr.client).InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}

	log.Println("Created a Todo with id ", todo.ID)
	return nil
}
