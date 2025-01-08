package todo

import (
	tContracts "todo-level-5/pkg/contract/todo"
	tRepo "todo-level-5/pkg/infrastructure/persistence/todo"
)

func ToCreateSpaceRes(todoM *tRepo.TodoModel) *tContracts.CreateTodoResponse {
	return &tContracts.CreateTodoResponse{
		Id:          todoM.ID,
		Title:       todoM.Title,
		Description: todoM.Description,
	}
}
