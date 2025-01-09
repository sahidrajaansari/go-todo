package handlers

import (
	"log"
	"net/http"
	tService "todo-level-5/pkg/application/todo"
	tContracts "todo-level-5/pkg/contract/todo"
	todoAgg "todo-level-5/pkg/domain/todo_aggregate"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoHandler struct {
	tService *tService.TodoService
	client   *mongo.Client
}

func NewTodoHandler(TodoService *tService.TodoService, client *mongo.Client) *TodoHandler {
	return &TodoHandler{
		tService: TodoService,
		client:   client,
	}
}

func todoCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("todoDB").Collection("todos")
}

func (th *TodoHandler) CreateTodo(ctx *gin.Context) {
	var requestBody *tContracts.CreateTodoRequest
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	todo, err := th.tService.Create(ctx, requestBody)
	if err != nil {
		ctx.JSON(402, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusAccepted, todo)
}

func (th *TodoHandler) GetTodoByID(ctx *gin.Context) {
	todoID := ctx.Param("id")
	log.Println("Finding TodoID of ", todoID)
	var todo todoAgg.Todo
	err := todoCollection(th.client).FindOne(ctx, bson.M{"_id": todoID}).Decode(&todo)
	if err != nil {
		ctx.JSON(402, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Get todo",
	})
}
