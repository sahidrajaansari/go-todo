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

func ToModelMetadata(md todoAgg.MetaData, test bool) todoAgg.MetaData {
	if test {
		return todoAgg.MetaData{
			CreatedAt: md.CreatedAt,
			UpdatedAt: md.UpdatedAt,
		}
	}
	if md.CreatedAt.IsZero() {
		md.CreatedAt = time.Now()
	}
	return todoAgg.MetaData{
		CreatedAt: md.CreatedAt,
		UpdatedAt: time.Now(),
	}
}

func ToTodoModel(todo *todoAgg.Todo, test bool) *TodoModel {
	return &TodoModel{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		Metadata:    ToModelMetadata(todo.MetaData, test),
	}
}
func (tm *TodoModel) ToBsonD() bson.D {
	return bson.D{
		{Key: "_id", Value: tm.ID},
		{Key: "title", Value: tm.Title},
		{Key: "description", Value: tm.Description},
		{Key: "status", Value: tm.Status},
		{Key: "createdat", Value: tm.Metadata.CreatedAt},
		{Key: "updatedat", Value: tm.Metadata.UpdatedAt},
	}
}

func (tm *TodoModel) toDomain() *todoAgg.Todo {
	return &todoAgg.Todo{
		ID:          tm.ID,
		Title:       tm.Title,
		Description: tm.Description,
		Status:      tm.Status,
		MetaData:    tm.Metadata, // include MetaData here
	}
}
