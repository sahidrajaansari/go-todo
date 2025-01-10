package todo

import (
	"context"
	"fmt"
	"testing"
	todoAgg "todo-level-5/pkg/domain/todo_aggregate"

	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestTodoRepo_Create(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	// defer mt.Close() // Ensures the mock client is properly closed after tests
	ctx := context.Background()

	type args struct {
		ctx     context.Context
		todoAgg *todoAgg.Todo
	}
	type test struct {
		id         int
		name       string
		beforeTest func(m *mtest.T)
		args       args
		wantErr    bool
	}
	tests := []test{
		{
			id:   1,
			name: "Create Todo - success",
			beforeTest: func(m *mtest.T) {
				m.AddMockResponses(mtest.CreateSuccessResponse())
			},
			args: args{
				ctx:     ctx,
				todoAgg: &todoAgg.TodoAgg,
			},
			wantErr: false,
		},
		{
			id:   2,
			name: "Create Todo with duplicate key - failure",
			beforeTest: func(m *mtest.T) {
				m.AddMockResponses(
					mtest.CreateWriteErrorsResponse(mtest.WriteError{
						Index:   0,
						Code:    11000,
						Message: fmt.Sprintf(`E11000 duplicate key error collection: todoDB.todos index: _id_ dup key: { _id: "%v" }`, todoAgg.TodoAgg.ID),
					}),
				)
			},
			args: args{
				ctx:     ctx,
				todoAgg: &todoAgg.TodoAgg,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		mt.Run(tt.name, func(mt *mtest.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(mt)
			}
			todoRepo := NewTodoRepo(mt.Client)

			err := todoRepo.Create(tt.args.ctx, tt.args.todoAgg)
			if (err != nil) != tt.wantErr {
				mt.Errorf("ID: %v Create() error = %v, wantErr = %v", tt.id, err, tt.wantErr)
			}
		})
	}
}

func TestTodoRepo_DeleteTodo(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	// defer mt.Close()
	ctx := context.Background()

	type args struct {
		ctx    context.Context
		todoID string
	}
	type test struct {
		id         int
		name       string
		beforeTest func(m *mtest.T)
		args       args
		wantErr    bool
		errMessage string
	}
	tests := []test{
		{
			id:   1,
			name: "Delete Todo - success",
			beforeTest: func(m *mtest.T) {
				m.AddMockResponses(mtest.CreateSuccessResponse())
			},
			args: args{
				ctx:    ctx,
				todoID: "2rNWzlOcXbhQAkRRk77StyDYAqf",
			},
			wantErr:    false,
			errMessage: "",
		},
		{
			id:   2,
			name: "Delete Todo - invalid query parameters",
			beforeTest: func(m *mtest.T) {
				m.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
					Code:    2,
					Message: "Invalid query parameters",
				}))
			},
			args: args{
				ctx:    ctx,
				todoID: "invalid_id",
			},
			wantErr:    true,
			errMessage: "invalid query parameters",
		},
		{
			id:   3,
			name: "Delete Todo - todo not found",
			beforeTest: func(m *mtest.T) {
				m.AddMockResponses(mtest.CreateCursorResponse(0, "todoDB.todos", mtest.FirstBatch))
			},
			args: args{
				ctx:    ctx,
				todoID: "nonexistent_id",
			},
			wantErr:    true,
			errMessage: "todo not found",
		},
	}

	for _, tt := range tests {
		mt.Run(tt.name, func(mt *mtest.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(mt)
			}
			todoRepo := &TodoRepo{
				client: mt.Client,
			}

			err := todoRepo.DeleteTodo(tt.args.ctx, tt.args.todoID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Test %d: DeleteTodo() error = %v, wantErr = %v", tt.id, err, tt.wantErr)
			}
			if tt.wantErr && err != nil && err.Error() != tt.errMessage {
				t.Errorf("Test %d: DeleteTodo() error message = %v, expected = %v", tt.id, err.Error(), tt.errMessage)
			}
		})
	}
}
