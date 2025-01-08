//go:build wireinject
// +build wireinject

package di

import (
	"context"
	"todo-level-5/config/db"

	h "todo-level-5/pkg/api/handlers"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func ProvideClient() *mongo.Client {
	return db.Connect(context.Background())
}

func ProvideTodoHandler() *h.TodoHandler {
	wire.Build(h.NewTodoHandler, ProvideClient)
	return nil
}

func InjectHandler() *h.Handlers {
	wire.Build(h.NewHandler, ProvideTodoHandler)
	return nil
}
