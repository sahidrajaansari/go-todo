package todo

import (
	"fmt"
	"time"
	todoAgg "todo-level-5/pkg/domain/todo_aggregate"

	"go.mongodb.org/mongo-driver/bson"
)

func getUpdatedFields(updatedTodoAgg *todoAgg.Todo, updatedFields *bson.M) error {
	// Only update fields that are non-empty
	if updatedTodoAgg.Title != "" {
		(*updatedFields)["title"] = updatedTodoAgg.Title
	}
	if updatedTodoAgg.Description != "" {
		(*updatedFields)["description"] = updatedTodoAgg.Description
	}
	if updatedTodoAgg.Status != "" {
		(*updatedFields)["status"] = updatedTodoAgg.Status
	}
	if len((*updatedFields)) == 0 {
		return fmt.Errorf("no fields have been updated")
	}

	(*updatedFields)["updatedat"] = time.Now()
	return nil
}

// Function to generate a sample Todo object
func CreateSampleTodo(id, title, description, status string, createdAt, updatedAt time.Time) todoAgg.Todo {
	return todoAgg.Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      status,
		MetaData: todoAgg.MetaData{
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}
}
