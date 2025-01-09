package handlers

import (
	"net/http"
	tService "todo-level-5/pkg/application/todo"
	tContracts "todo-level-5/pkg/contract/todo"

	"github.com/gin-gonic/gin"
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

// func todoCollection(client *mongo.Client) *mongo.Collection {
// 	return client.Database("todoDB").Collection("todos")
// }

func (th *TodoHandler) CreateTodo(ctx *gin.Context) {
	var requestBody *tContracts.CreateTodoRequest
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	todo, err := th.tService.Create(ctx, requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to Create an Item",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, todo)
}

func (th *TodoHandler) GetTodoByID(ctx *gin.Context) {
	todo, err := th.tService.GetTodoByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve todo item",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Todo retrieved successfully",
		"todo":    todo,
	})

}
