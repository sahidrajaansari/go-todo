package todo

import "context"

func createTestContext(todoID string) context.Context {
	ctx := context.WithValue(context.Background(), "todoID", todoID)
	return ctx
}
