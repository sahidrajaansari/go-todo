//go:build wireinject
// +build wireinject

package di

import (
	"context"
	"todo-level-5/config/db"

	h "todo-level-5/pkg/api/handlers"
	svcInter "todo-level-5/pkg/application/services"
	tApp "todo-level-5/pkg/application/todo"
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

func ProvideTodoService() *tApp.TodoService {
	wire.Build(tApp.NewTodoService, todoRepoSet)
	return nil
}

var todoServSet = wire.NewSet(
	ProvideTodoService,
	wire.Bind(new(svcInter.ITodoService), new(*tApp.TodoService)),
)

func ProvideTodoHandler() *h.TodoHandler {
	wire.Build(h.NewTodoHandler, todoServSet)
	return nil
}

func InjectHandler() *h.Handlers {
	wire.Build(h.NewHandler, ProvideTodoHandler)
	return nil
}
