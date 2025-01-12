package todo

import "context"

func getContext() context.Context {
	ctx := context.WithValue(context.Background(), "todoID", "Valid")
	return ctx
}
