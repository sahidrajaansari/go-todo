package server

import (
	"github.com/gin-gonic/gin"
)

func SetupTodoRoutes(rg *gin.RouterGroup, h *HttpServer) {
	rg.GET("/", h.handlers.TodoHandler.GetTodos)
	rg.POST("/", h.handlers.TodoHandler.CreateTodo)
	rg.GET("/:id", h.handlers.TodoHandler.GetTodoByID)
	// rg.PUT("/:id", UpdateTodo)
	// rg.DELETE("/:id", DeleteTodo)
}
