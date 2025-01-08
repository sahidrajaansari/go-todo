package handlers

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoHandler struct {
	client *mongo.Client
}

func NewTodoHandler(client *mongo.Client) *TodoHandler {
	return &TodoHandler{
		client: client,
	}
}

func collections(client *mongo.Client) *mongo.Collection {
	return client.Database("todoDB").Collection("todos")
}

func (th *TodoHandler) CreateTodo(ctx *gin.Context) {
	type Todo struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}
	var todo Todo
	if err := ctx.BindJSON(&todo); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Println(todo)
	type todoModel struct {
		ID          string `bson:"_id"`
		Title       string `bson:"title"`
		Description string `bson:"description"`
		Status      string `bson:"status"`
	}
	todoM := todoModel{
		ID:          "1", // Replace with a generated unique ID, e.g., ksuid.New().String()
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
	}
	rTodo, err := collections(th.client).InsertOne(context.Background(), todoM)
	if err != nil {
		ctx.JSON(402, gin.H{
			"error": "DB Connection Error",
		})
		return
	}
	log.Println(rTodo)
	ctx.JSON(200, gin.H{
		"message": "Create todo",
	})
}
