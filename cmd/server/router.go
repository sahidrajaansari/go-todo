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

func SetupPingRoutes(rg *gin.RouterGroup) {
	rg.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
