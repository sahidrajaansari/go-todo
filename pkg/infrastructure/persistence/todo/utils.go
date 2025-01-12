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

	(*updatedFields)["updatedAt"] = time.Now()
	return nil
}
