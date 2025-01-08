package todo

import (
	tContracts "todo-level-5/pkg/contract/todo"
	todoagg "todo-level-5/pkg/domain/todo_aggregate"
)

func FromSpaceTodoRequest(todoId string, tsr *tContracts.CreateTodoRequest) *todoagg.Todo {
	return todoagg.NewTodo(todoId, tsr.Title, tsr.Description, tsr.Status)
}

func ToCreateSpaceRes(tAgg *todoagg.Todo) *tContracts.CreateTodoResponse {
	return &tContracts.CreateTodoResponse{
		Id:          tAgg.ID,
		Title:       tAgg.Title,
		Description: tAgg.Description,
	}
}
