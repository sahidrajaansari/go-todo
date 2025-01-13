package todo

import (
	"context"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func createTestContext(todoID string) context.Context {
	ctx := context.WithValue(context.Background(), "todoID", todoID)
	return ctx
}

func CreateTestGinContext(method, url string) *gin.Context {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Create a test request
	req := httptest.NewRequest(method, url, nil)

	// Initialize the Gin context with the test request
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	return ctx
}
