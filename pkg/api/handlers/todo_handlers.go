package handlers

import (
	"net/http"
	tService "todo-level-5/pkg/application/todo"
	tContracts "todo-level-5/pkg/contract/todo"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	tService *tService.TodoService
}

func NewTodoHandler(TodoService *tService.TodoService) *TodoHandler {
	return &TodoHandler{
		tService: TodoService,
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
	ctx.Set("todoID", ctx.Param("id"))

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

func (th *TodoHandler) UpdateTodoByID(ctx *gin.Context) {
	var requestBody *tContracts.UpdateTodoRequest
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	requestBody.SetDefaultValues()

	todo, err := th.tService.UpdateTodoByID(ctx, requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve todo item",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Todo Updated successfully",
		"todo":    todo,
	})
}

func (th *TodoHandler) GetTodos(ctx *gin.Context) {
	// Get the todos from the service layer
	todos, err := th.tService.GetTodos(ctx)
	if err != nil {
		// Handle the error and return an appropriate response
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve todos",
			"details": err.Error(),
		})
		return
	}

	// Return the list of todos with a success message
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Todos retrieved successfully",
		"todos":   todos, // Changed "todo" to "todos"
	})
}

func (th *TodoHandler) DeleteTodo(ctx *gin.Context) {
	err := th.tService.DeleteTodo(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "todo not found",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "Todo Had been Deleted",
	})
}
