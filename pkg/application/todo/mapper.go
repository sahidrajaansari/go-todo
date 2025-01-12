package todo

import (
	tContracts "todo-level-5/pkg/contract/todo"
	todoagg "todo-level-5/pkg/domain/todo_aggregate"
)

func fromCreateTodoRequest(todoID string, tsr *tContracts.CreateTodoRequest) *todoagg.Todo {
	return todoagg.NewTodo(todoID, tsr.Title, tsr.Description, tsr.Status)
}

func fromUpdateTodoRequest(todoID string, tsr *tContracts.UpdateTodoRequest) *todoagg.Todo {
	return todoagg.NewTodo(todoID, tsr.Title, tsr.Description, tsr.Status)
}

func ToCreateTodoRes(tAgg *todoagg.Todo) *tContracts.CreateTodoResponse {
	return &tContracts.CreateTodoResponse{
		ID:          tAgg.ID,
		Title:       tAgg.Title,
		Description: tAgg.Description,
	}
}

func ToGetByIDRes(tAgg *todoagg.Todo) *tContracts.GetTodoResponse {
	return &tContracts.GetTodoResponse{
		ID:          tAgg.ID,
		Title:       tAgg.Title,
		Description: tAgg.Description,
		Status:      tAgg.Status,
	}
}

func toUpateTodoRes(tAgg *todoagg.Todo) *tContracts.UpdateTodoResponse {
	return &tContracts.UpdateTodoResponse{
		ID:          tAgg.ID,
		Title:       tAgg.Title,
		Description: tAgg.Description,
		Status:      tAgg.Status,
		UpdatedAt:   tAgg.MetaData.UpdatedAt.String(),
	}
}
