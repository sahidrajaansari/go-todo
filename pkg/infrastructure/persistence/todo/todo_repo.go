package todo

import (
	"context"
	"errors"
	"fmt"
	"log"
	todoAgg "todo-level-5/pkg/domain/todo_aggregate"

	"github.com/ajclopez/mgs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	todo := ToTodoModel(todoAgg)

	_, err := todoCollection(tr.client).InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}

	log.Println("Created a Todo with id ", todo.ID)
	return nil
}

func (tr *TodoRepo) GetTodoByID(ctx context.Context, todoID string) (*todoAgg.Todo, error) {
	log.Println("Fetching todo with id ", todoID)

	var todo TodoModel
	// Todo Without MGS
	// err := todoCollection(tr.client).FindOne(ctx, bson.M{"_id": todoID}).Decode(&todo)
	// if err != nil {
	// 	if err == mongo.ErrNoDocuments {
	// 		return nil, errors.New("todo not found")
	// 	}
	// 	return nil, err
	// }

	//TODO WIth MGS
	opts := mgs.FindOption()
	query := fmt.Sprintf("_id=%s", todoID)
	result, err := mgs.MongoGoSearch(query, opts)
	if err != nil {
		log.Println(ctx, fmt.Sprintf("Invalid query params: %v", query), err)
		return nil, errors.New("invalid query parameters ")
	}

	findOpts := options.FindOne()
	findOpts.SetProjection(result.Projection)

	if err := todoCollection(tr.client).FindOne(ctx, result.Filter, findOpts).Decode(&todo); err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("todo not found")
			return nil, errors.New("todo not found")
		}
		return nil, err
	}

	return todo.toDomain(), nil

}
