package todo

import (
	"time"
	todoAgg "todo-level-5/pkg/domain/todo_aggregate"
)

type TodoModel struct {
	ID          string           `bson:"_id" validate:"required"`
	Title       string           `bson:"title" validate:"required"`
	Description string           `bson:"description"`
	Status      string           `bson:"status" validate:"required"`
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

func (tm *TodoModel) toDomain() *todoAgg.Todo {
	return &todoAgg.Todo{
		ID:          tm.ID,
		Title:       tm.Title,
		Description: tm.Description,
		Status:      tm.Status,
	}
}
