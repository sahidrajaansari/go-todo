package todo

import (
	"errors"
	tContracts "todo-level-5/pkg/contract/todo"
	"todo-level-5/pkg/domain/todo_aggregate"
)

var todoAgg = todo_aggregate.TodoAgg

var validTodoID = todoAgg.ID
var inValidTodoID = "121212311"

var errInvalidTodoID = errors.New("invalid todo ID")
var errRepositryError = errors.New("database error: unable to insert new todo")

// Prototype test data for CreateTodoRequest
var ValidCreateTodoRequest = tContracts.CreateTodoRequest{
	Title:       "Learn GoLang",
	Description: "Complete the GoLang tutorials on concurrency and testing.",
	Status:      "pending",
}
