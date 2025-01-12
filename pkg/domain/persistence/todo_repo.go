package persistence

import (
	"context"
	todoAgg "todo-level-5/pkg/domain/todo_aggregate"
	tPersist "todo-level-5/pkg/infrastructure/persistence/todo"
)

var _ ITodoRepo = (*tPersist.TodoRepo)(nil)

type ITodoRepo interface {
	Create(ctx context.Context, todoAgg *todoAgg.Todo) error
	GetTodoByID(ctx context.Context, todoID string) (*todoAgg.Todo, error)
	GetTodos(ctx context.Context, query string) ([]*todoAgg.Todo, error)
	UpdateTodo(ctx context.Context, todoID string, updatedTodoAgg *todoAgg.Todo) (*todoAgg.Todo, error)
	DeleteTodo(ctx context.Context, todoID string) error
}
