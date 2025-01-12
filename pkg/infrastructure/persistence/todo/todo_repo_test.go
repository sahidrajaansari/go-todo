package todo

import (
	"context"
	"fmt"
	"testing"
	todoAgg "todo-level-5/pkg/domain/todo_aggregate"

	"go.mongodb.org/mongo-driver/bson"
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
	// defer mt.Close() // Ensures resources are properly released
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
	}
	tests := []test{
		{
			id:   1,
			name: "Delete Todo - Success",
			beforeTest: func(m *mtest.T) {
				m.AddMockResponses(mtest.CreateSuccessResponse(bson.E{Key: "n", Value: 1}))
			},
			args: args{
				ctx:    ctx,
				todoID: "validTodoID", // Replace with a valid static ID
			},
			wantErr: false, // No error want
		},
		{
			id:   2,
			name: "Delete Todo with non-Existent TodoID - Failure",
			beforeTest: func(m *mtest.T) {
				m.AddMockResponses(mtest.CreateSuccessResponse(bson.E{Key: "n", Value: 0}))
			},
			args: args{
				ctx:    ctx,
				todoID: "nonExistentTodoID",
			},
			wantErr: true,
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
		})
	}
}

// func TestTodoRepo_GetTodoByID(t *testing.T) {
// 	ctx := context.Background()
// 	type args struct {
// 		ctx    context.Context
// 		todoID string
// 	}
// 	type test struct {
// 		id         int
// 		name       string
// 		tr         *TodoRepo
// 		beforeTest func(m *mtest.T)
// 		args       args
// 		want       *todoAgg.Todo
// 		wantErr    bool
// 	}
// 	tests := []test{
// 		{
// 			id:   1,
// 			name: "Get Todo - Success",
// 			beforeTest: func(m *mtest.T) {
// 				m.AddMockResponses(
// 					mtest.CreateCursorResponse(1, "test.todo", mtest.FirstBatch, bson.D{
// 						{Key: "_id", Value: "validTodoID"},
// 						{Key: "title", Value: "Sample Todo"},
// 						{Key: "completed", Value: "Hello"},
// 					}),
// 				)
// 			},
// 			args: args{
// 				ctx:    ctx,
// 				todoID: "validTodoID",
// 			},
// 			wantErr: false,
// 			want:    &todoAgg.Todo{ID: "validTodoID", Title: "Sample Todo", Status: "Compeleted"},
// 		},
// 		// {
// 		// 	id:   2,
// 		// 	name: "Get Todo - Document Not Found",
// 		// 	beforeTest: func(m *mtest.T) {
// 		// 		m.AddMockResponses(mtest.CreateCursorResponse(0, "test.todo", mtest.FirstBatch))
// 		// 	},
// 		// 	args: args{
// 		// 		ctx:    ctx,
// 		// 		todoID: "nonExistentTodoID",
// 		// 	},
// 		// 	wantErr: true,
// 		// 	want:    nil,
// 		// },
// 		// {
// 		// 	id:   3,
// 		// 	name: "Get Todo - Invalid Query Parameters",
// 		// 	beforeTest: func(m *mtest.T) {
// 		// 		m.AddMockResponses(bson.D{
// 		// 			{Key: "ok", Value: 0},
// 		// 			{Key: "errmsg", Value: "invalid query"},
// 		// 		})
// 		// 	},
// 		// 	args: args{
// 		// 		ctx:    ctx,
// 		// 		todoID: "invalidTodoID",
// 		// 	},
// 		// 	wantErr: true,
// 		// 	want:    nil,
// 		// },
// 		// {
// 		// 	id:   4,
// 		// 	name: "Get Todo - BSON Decoding Error",
// 		// 	beforeTest: func(m *mtest.T) {
// 		// 		m.AddMockResponses(mtest.CreateCursorResponse(1, "test.todo", mtest.FirstBatch, bson.D{
// 		// 			{Key: "_id", Value: bson.A{1, 2, 3}}, // Invalid BSON structure
// 		// 		}))
// 		// 	},
// 		// 	args: args{
// 		// 		ctx:    ctx,
// 		// 		todoID: "invalidBsonTodoID",
// 		// 	},
// 		// 	wantErr: true,
// 		// 	want:    nil,
// 		// },
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.tr.GetTodoByID(tt.args.ctx, tt.args.todoID)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("TodoRepo.GetTodoByID() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("TodoRepo.GetTodoByID() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
