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

func TestTodoRepo_UpdateTodo(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	type args struct {
		ctx            context.Context
		todoID         string
		updatedTodoAgg *todoAgg.Todo
	}
	type test struct {
		id         int
		name       string
		beforeTest func(m *mtest.T) // Function to set up mock responses
		args       args
		want       *todoAgg.Todo // Expected result
		wantErr    bool          // Expected error status
	}
	ctx := context.Background()

	tests := []test{
		{
			id:   1,
			name: "Update Todo Update Everything - Success",
			beforeTest: func(m *mtest.T) {
				m.AddMockResponses(mtest.CreateSuccessResponse(bson.E{
					Key: "value", Value: ToTodoModel(&todoAgg.TodoAgg, true).ToBsonD(),
				}))
			},
			args: args{
				ctx:    ctx,
				todoID: "valid",
				updatedTodoAgg: &todoAgg.Todo{
					Title:       "Changeing The Title",
					Description: "Change Description",
					Status:      "Change Status",
				},
			},
			want: &todoAgg.Todo{
				ID:          "valid",
				Title:       "Changeing The Title",
				Description: "Change Description",
				Status:      "Change Status",
			},
			wantErr: false,
		},
		{
			id:   2,
			name: "Update Todo Update Single Field - Success",
			beforeTest: func(m *mtest.T) {
				m.AddMockResponses(mtest.CreateSuccessResponse(bson.E{
					Key: "value", Value: ToTodoModel(&todoAgg.TodoAgg, true).ToBsonD(),
				}))
			},
			args: args{
				ctx:    ctx,
				todoID: "valid",
				updatedTodoAgg: &todoAgg.Todo{
					Title:       "Changing The Title",
					Description: "",
					Status:      "",
				},
			},
			want: &todoAgg.Todo{
				ID:          "valid",
				Title:       "Changing The Title",
				Description: "Change Description",
				Status:      "Change Status",
			},
			wantErr: false,
		},
		{
			id:   3,
			name: "Document Not found - Failure",
			beforeTest: func(m *mtest.T) {
				m.AddMockResponses(mtest.CreateSuccessResponse())
			},
			args: args{
				ctx:    ctx,
				todoID: "valid",
				updatedTodoAgg: &todoAgg.Todo{
					Title:       "Changing The Title",
					Description: "",
					Status:      "",
				},
			},
			want: &todoAgg.Todo{
				ID:          "valid",
				Title:       "Changing The Title",
				Description: "Change Description",
				Status:      "Change Status",
			},
			wantErr: true,
		},
		{
			id:   4,
			name: "Nothing to Update Error - Failure",
			beforeTest: func(m *mtest.T) {
				m.AddMockResponses(mtest.CreateSuccessResponse())
			},
			args: args{
				ctx:    ctx,
				todoID: "valid",
				updatedTodoAgg: &todoAgg.Todo{
					Title:       "",
					Description: "",
					Status:      "",
				},
			},
			want: &todoAgg.Todo{
				ID:          "valid",
				Title:       "Changing The Title",
				Description: "Change Description",
				Status:      "Change Status",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		mt.Run(tt.name, func(mt *mtest.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(mt)
			}
			todoRepo := &TodoRepo{client: mt.Client}
			_, err := todoRepo.UpdateTodo(tt.args.ctx, tt.args.todoID, tt.args.updatedTodoAgg)
			// Check for errors and ensure they match the expected outcome
			if (err != nil) != tt.wantErr {
				t.Errorf("Test ID %d - UpdateTodo() error = %v, wantErr = %v", tt.id, err, tt.wantErr)
			}

			// if tt.want != nil && !reflect.DeepEqual(result, tt.want) {
			// 	t.Errorf("Test ID %d - UpdateTodo() = %v, expected = %v", tt.id, result, tt.want)
			// }
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

func TestTodoRepo_GetTodos(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	// defer mt.Close() // Ensure the mock client is closed after the test

	type args struct {
		ctx   context.Context
		query string
	}
	type test struct {
		id         int
		name       string
		beforeTest func(m *mtest.T) // Function to mock the database response
		args       args
		want       []*todoAgg.Todo // Expected list of todos
		wantErr    bool            // Whether an error is expected
	}

	ctx := context.Background()

	tests := []test{
		{
			id:   1,
			name: "Get Todos - Success",
			beforeTest: func(m *mtest.T) {
				first := mtest.CreateCursorResponse(1, "todoDB.todos", mtest.FirstBatch, ToTodoModel(&todoAgg.TodoAgg, true).ToBsonD())
				nextBatch := mtest.CreateCursorResponse(1, "todoDB.todos", mtest.NextBatch, ToTodoModel(&todoAgg.TodoAgg, true).ToBsonD())
				killCursor := mtest.CreateCursorResponse(0, "todoDB.todos", mtest.NextBatch)
				m.AddMockResponses(first, nextBatch, killCursor)
			},
			args: args{
				ctx:   ctx,
				query: "status=pending",
			},
			want: []*todoAgg.Todo{
				&todoAgg.TodoAgg, &todoAgg.TodoAgg,
			},
			wantErr: false,
		},
		{
			id:   2,
			name: "Get Todos - Empty Result",
			beforeTest: func(m *mtest.T) {
				m.AddMockResponses(mtest.CreateCursorResponse(0, "todoDB.todos", mtest.FirstBatch))
			},
			args: args{
				ctx:   ctx,
				query: "status=completed",
			},
			want:    []*todoAgg.Todo{}, // Expecting no todos
			wantErr: false,
		},
		{
			id:   3,
			name: "Get Todos - Database Error",
			beforeTest: func(m *mtest.T) {
				// Simulate database error
				m.AddMockResponses(
					mtest.CreateCommandErrorResponse(mtest.CommandError{
						Code:    11000,
						Message: "Database error",
					}),
				)
			},
			args: args{
				ctx:   ctx,
				query: "status=pending",
			},
			want:    nil, // Expecting no result due to error
			wantErr: true,
		},
		{
			id:   4,
			name: "Invalid Query Format - Failure",
			beforeTest: func(m *mtest.T) {
				// No mock responses needed as this should fail before a database call
			},
			args: args{
				ctx:   ctx,
				query: "status=invalid%query",
			},
			want:    nil,  // Expecting no result due to query error
			wantErr: true, // Error expected
		},
		{
			id:   5,
			name: "Bson-Error -Failure",
			beforeTest: func(m *mtest.T) {
				first := mtest.CreateCursorResponse(1, "todoDB.todos", mtest.FirstBatch, correctMongoData1)
				nextBatch1 := mtest.CreateCursorResponse(1, "todoDB.todos", mtest.NextBatch, errorMongoData)
				nextBatch2 := mtest.CreateCursorResponse(1, "todoDB.todos", mtest.NextBatch, correctMongoData2)
				killCursor := mtest.CreateCursorResponse(0, "todoDB.todos", mtest.NextBatch)
				m.AddMockResponses(first, nextBatch1, nextBatch2, killCursor)
			},
			args: args{
				ctx:   ctx,
				query: "correct query",
			},
			want:    nil,   // Expecting no result due to query error
			wantErr: false, // Error expected
		},
	}

	for _, tt := range tests {
		mt.Run(tt.name, func(mt *mtest.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(mt)
			}

			// Initialize the repository
			todoRepo := NewTodoRepo(mt.Client)

			// Call the GetTodos method
			_, err := todoRepo.GetTodos(tt.args.ctx, tt.args.query)

			// Check error status
			if (err != nil) != tt.wantErr {
				t.Errorf("Test ID %d - GetTodos() error = %v, wantErr = %v", tt.id, err, tt.wantErr)
			}

			// // Check the result if expected
			// if tt.want != nil && !reflect.DeepEqual(result, tt.want) {
			// 	t.Errorf("Test ID %d - GetTodos() = %v, expected = %v", tt.id, result, tt.want)
			// }
		})
	}
}
