package server

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(h *HttpServer) {
	api := h.engine.Group("/api/v1")
	ping := api.Group("/ping")
	todoGroup := api.Group("/todos")

	SetupTodoRoutes(todoGroup, h)
	SetupPingRoutes(ping)
}

func SetupTodoRoutes(rg *gin.RouterGroup, h *HttpServer) {
	// rg.GET("/", GetTodos)
	rg.POST("/", h.handlers.TodoHandler.CreateTodo)
	// rg.GET("/:id", func(ctx *gin.Context) {
	// 	id := ctx.Param("id")
	// 	fmt.Println(id)
	// 	ctx.JSON(200, gin.H{
	// 		"message": "Get todo",
	// 	})
	// })
	// rg.PUT("/:id", UpdateTodo)
	// rg.DELETE("/:id", DeleteTodo)
}

func SetupPingRoutes(rg *gin.RouterGroup) {
	rg.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
