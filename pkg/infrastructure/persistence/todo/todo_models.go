package todo

import (
	"time"
	todoagg "todo-level-5/pkg/domain/todo_aggregate"
)

type TodoModel struct {
	ID          string           `bson:"_id"`
	Title       string           `bson:"title"`
	Description string           `bson:"description"`
	Status      string           `bson:"status"`
	Metadata    todoagg.MetaData `bson:",inline"`
}

func ToModelMetadata(md todoagg.MetaData) todoagg.MetaData {
	if md.CreatedAt.IsZero() {
		md.CreatedAt = time.Now()
	}
	return todoagg.MetaData{
		CreatedAt: md.CreatedAt,
		UpdatedAt: time.Now(),
	}
}

func ToSpaceModel(todo *todoagg.Todo) *TodoModel {
	return &TodoModel{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		Metadata:    ToModelMetadata(todo.MetaData),
	}
}
