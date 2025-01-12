package todo

import (
	"context"
	"fmt"
	"reflect"
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

func TestTodoRepo_GetTodoByID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	// defer mt.Close() // Close the mock client after test

	// Define arguments structure for the test cases
	type args struct {
		ctx    context.Context
		todoID string
	}
	// Define the structure for each test case
	type test struct {
		id         int
		name       string
		beforeTest func(m *mtest.T) // Function to set up mock responses
		args       args
		want       *todoAgg.Todo // Expected result
		wantErr    bool          // Expected error status
	}
	ctx := context.Background()

	// Define all test cases
	tests := []test{
		{
			id:   1,
			name: "Get Todo By ID - Success", // Successful test case
			beforeTest: func(m *mtest.T) {
				m.AddMockResponses(
					mtest.CreateCursorResponse(1, "todoDB.todos", mtest.FirstBatch,
						ToTodoModel(&todoAgg.TodoAgg, true).ToBsonD()),
				)
			},
			args: args{
				ctx:    ctx,
				todoID: "valid", // Valid Todo ID
			},
			want:    &todoAgg.TodoAgg,
			wantErr: false,
		},
		{
			id:   2,
			name: "Get Todo by ID Not Found - Failure", // Todo not found scenario
			beforeTest: func(m *mtest.T) {
				m.AddMockResponses(mtest.CreateCursorResponse(0, "todoDB.todos", mtest.FirstBatch))
			},
			args: args{
				ctx:    ctx,
				todoID: "nonExistentTodoID", // Invalid Todo ID
			},
			want:    nil,
			wantErr: true,
		},
		{
			id:   3,
			name: "Get Todo by ID BSON Error - Failure", // BSON decoding error scenario
			beforeTest: func(m *mtest.T) {
				m.AddMockResponses(mtest.CreateCursorResponse(0, "todoDB.todos", mtest.FirstBatch, bson.D{
					{Key: "_id", Value: 1}, // Incorrect BSON format (simulated error)
				}))
			},
			args: args{
				ctx:    ctx,
				todoID: "valid", // Valid Todo ID
			},
			want:    nil,
			wantErr: true,
		},
	}

	// Loop through all test cases
	for _, tt := range tests {
		mt.Run(tt.name, func(mt *mtest.T) {
			// Set up the mock responses for the current test case
			if tt.beforeTest != nil {
				tt.beforeTest(mt)
			}

			// Create the TodoRepo instance
			todoRepo := &TodoRepo{client: mt.Client}

			// Call the method being tested
			result, err := todoRepo.GetTodoByID(tt.args.ctx, tt.args.todoID)

			// Check for errors and ensure they match the expected outcome
			if (err != nil) != tt.wantErr {
				t.Errorf("Test ID %d - GetTodoByID() error = %v, wantErr = %v", tt.id, err, tt.wantErr)
			}

			// If there is an expected result, compare it with the actual result
			if tt.want != nil && !reflect.DeepEqual(result, tt.want) {
				t.Errorf("Test ID %d - GetTodoByID() = %v, expected = %v", tt.id, result, tt.want)
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
