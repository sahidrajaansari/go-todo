package todo

import (
	"todo-level-5/pkg/contract/todo"

	"github.com/segmentio/ksuid"
)

type TodoModel struct {
	ID          string `bson:"_id"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Status      string `bson:"status"`
}

func ToSpaceModel(todo todo.CreateTodoRequest) *TodoModel {
	return &TodoModel{
		ID:          ksuid.New().String(),
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
	}
}
