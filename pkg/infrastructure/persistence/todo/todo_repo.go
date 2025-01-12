package todo

import (
	"context"
	"errors"
	"fmt"
	"log"
	todoAgg "todo-level-5/pkg/domain/todo_aggregate"

	"github.com/ajclopez/mgs"
	"go.mongodb.org/mongo-driver/bson"
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
	todo := ToTodoModel(todoAgg, false)

	_, err := todoCollection(tr.client).InsertOne(ctx, todo)
	if err != nil {
		return err
	}

	// log.Println("Created a Todo with id ", todo.ID)
	return nil
}

func (tr *TodoRepo) GetTodoByID(ctx context.Context, todoID string) (*todoAgg.Todo, error) {
	// log.Println("Fetching todo with id ", todoID)

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
	result, _ := mgs.MongoGoSearch(query, opts)
	// if err != nil {
	// 	log.Println(ctx, fmt.Sprintf("Invalid query params: %v", query), err)
	// 	return nil, errors.New("invalid query parameters ")
	// }

	findOpts := options.FindOne()
	findOpts.SetProjection(result.Projection)

	if err := todoCollection(tr.client).FindOne(ctx, result.Filter, findOpts).Decode(&todo); err != nil {
		return nil, err
	}

	return todo.toDomain(), nil

}

func (tr *TodoRepo) GetTodos(ctx context.Context) ([]*todoAgg.Todo, error) {
	var todos []*todoAgg.Todo
	log.Println("Fetching all the todos")

	cursor, err := todoCollection(tr.client).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var todo TodoModel
		if err := cursor.Decode(&todo); err != nil {
			log.Println("Error decoding todo:", err)
			continue
		}
		todos = append(todos, todo.toDomain())
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

func (tr *TodoRepo) UpdateTodo(ctx context.Context, todoID string, updatedTodoAgg *todoAgg.Todo) (*todoAgg.Todo, error) {
	var todo TodoModel
	updatedFields := bson.M{}
	err := getUpdatedFields(updatedTodoAgg, &updatedFields)
	if err != nil {
		return nil, err
	}
	updateQuery := bson.M{
		"$set": updatedFields,
	}
	log.Println("Int Repo; ", updateQuery)

	err = todoCollection(tr.client).FindOneAndUpdate(
		ctx,
		bson.M{"_id": todoID},
		updateQuery,
		options.FindOneAndUpdate().SetReturnDocument(options.After), // Return the updated document
	).Decode(&todo)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no Todo found with ID: %s", todoID)
		}
		return nil, fmt.Errorf("failed to update Todo: %v", err)
	}

	return todo.toDomain(), nil // Return the updated Todo
}

func (tr *TodoRepo) DeleteTodo(ctx context.Context, todoID string) error {
	result, err := todoCollection(tr.client).DeleteOne(ctx, bson.M{
		"_id": todoID,
	})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("document not found")
	}
	return nil
}
