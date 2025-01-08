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
