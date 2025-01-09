package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupTodoRoutes(rg *gin.RouterGroup, h *HttpServer) {
	// rg.GET("/", GetTodos)
	rg.POST("/", h.handlers.TodoHandler.CreateTodo)
	rg.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		fmt.Println(id)
		ctx.JSON(200, gin.H{
			"message": "Get todo",
		})
	})
	// rg.PUT("/:id", UpdateTodo)
	// rg.DELETE("/:id", DeleteTodo)
}
