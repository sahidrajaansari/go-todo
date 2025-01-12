//go:build wireinject
// +build wireinject

package di

import (
	"context"
	"todo-level-5/config/db"

	h "todo-level-5/pkg/api/handlers"
	tService "todo-level-5/pkg/application/todo"
	iPersist "todo-level-5/pkg/domain/persistence"
	tRepo "todo-level-5/pkg/infrastructure/persistence/todo"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func ProvideClient() *mongo.Client {
	return db.Connect(context.Background())
}

func ProvideTodoRepo() *tRepo.TodoRepo {
	wire.Build(tRepo.NewTodoRepo, ProvideClient)
	return nil
}

var todoRepoSet = wire.NewSet(
	ProvideTodoRepo,
	wire.Bind(new(iPersist.ITodoRepo), new(*tRepo.TodoRepo)),
)

func ProvideTodoService() *tService.TodoService {
	wire.Build(tService.NewTodoService, todoRepoSet)
	return nil
}

func ProvideTodoHandler() *h.TodoHandler {
	wire.Build(h.NewTodoHandler, ProvideTodoService)
	return nil
}

func InjectHandler() *h.Handlers {
	wire.Build(h.NewHandler, ProvideTodoHandler)
	return nil
}
