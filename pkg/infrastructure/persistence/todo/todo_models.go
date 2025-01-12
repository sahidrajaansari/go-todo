package todo

import (
	"time"
	todoAgg "todo-level-5/pkg/domain/todo_aggregate"

	"go.mongodb.org/mongo-driver/bson"
)

type TodoModel struct {
	ID          string           `bson:"_id"`
	Title       string           `bson:"title"`
	Description string           `bson:"description"`
	Status      string           `bson:"status"`
	Metadata    todoAgg.MetaData `bson:",inline"`
}

func ToModelMetadata(md todoAgg.MetaData) todoAgg.MetaData {
	if md.CreatedAt.IsZero() {
		md.CreatedAt = time.Now()
	}
	return todoAgg.MetaData{
		CreatedAt: md.CreatedAt,
		UpdatedAt: time.Now(),
	}
}

func ToTodoModel(todo *todoAgg.Todo) *TodoModel {
	return &TodoModel{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		Metadata:    ToModelMetadata(todo.MetaData),
	}
}
func (tm *TodoModel) ToBsonD() bson.D {
	return bson.D{
		{Key: "_id", Value: tm.ID},
		{Key: "title", Value: tm.Title},
		{Key: "description", Value: tm.Description},
		{Key: "status", Value: tm.Status},
	}
}

func (tm *TodoModel) toDomain() *todoAgg.Todo {
	return &todoAgg.Todo{
		ID:          tm.ID,
		Title:       tm.Title,
		Description: tm.Description,
		Status:      tm.Status,
	}
}
