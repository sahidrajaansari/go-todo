package todo

import (
	"context"
	"errors"
	"log"
	todoAgg "todo-level-5/pkg/domain/todo_aggregate"

	"go.mongodb.org/mongo-driver/bson"
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

func (tr *TodoRepo) Create(ctx context.Context, todoAgg *todoAgg.Todo) error {
	todo := ToSpaceModel(todoAgg)

	_, err := todoCollection(tr.client).InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}

	log.Println("Created a Todo with id ", todo.ID)
	return nil
}

func (tr *TodoRepo) GetTodoByID(ctx context.Context, todoID string) (*todoAgg.Todo, error) {
	var todo *todoAgg.Todo
	collection := todoCollection(tr.client)
	err := collection.FindOne(ctx, bson.M{"_id": todoID}).Decode(&todo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("todo not found")
		}
		return nil, err
	}
	return todo, nil
}
