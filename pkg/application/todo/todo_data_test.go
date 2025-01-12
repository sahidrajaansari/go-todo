package todo

import (
	"errors"
	"todo-level-5/pkg/domain/todo_aggregate"
)

var todoAgg = todo_aggregate.TodoAgg

var validTodoID = todoAgg.ID
var inValidTodoID = "121212311"

var errInvalidTodoID = errors.New("invalid todo ID")
