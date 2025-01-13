package todo

import (
	"context"
	"reflect"
	"testing"
	tContracts "todo-level-5/pkg/contract/todo"
	"todo-level-5/pkg/domain/persistence/mock"
	"todo-level-5/pkg/domain/todo_aggregate"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

type fields struct {
	tRepo *mock.MockITodoRepo
}

// Haven't Used Want as it Can Never Be Wrong Created inside it Only
func TestTodoService_Create(t *testing.T) {
	type args struct {
		ctx context.Context
		tsr *tContracts.CreateTodoRequest
	}
	type test struct {
		id         int
		name       string
		args       args
		beforeTest func(f *fields)
		wantErr    bool
	}

	tests := []test{
		{
			id:   1,
			name: "Create_Todo_Success",
			args: args{
				ctx: context.Background(),
				tsr: &ValidCreateTodoRequest,
			},
			beforeTest: func(f *fields) {
				f.tRepo.EXPECT().
					Create(gomock.Any(), gomock.AssignableToTypeOf(&todo_aggregate.Todo{})).
					Return(nil)
			},
			wantErr: false,
		},
		{
			id:   2,
			name: "Create_Todo_Repository_Error",
			args: args{
				ctx: context.Background(),
				tsr: &ValidCreateTodoRequest,
			},
			beforeTest: func(f *fields) {
				f.tRepo.EXPECT().
					Create(gomock.Any(), gomock.AssignableToTypeOf(&todo_aggregate.Todo{})).
					Return(errRepositoryError)
			},
			wantErr: true,
		},
	}

	// Running tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Initialize the mock repository
			f := fields{
				tRepo: mock.NewMockITodoRepo(ctrl),
			}

			// Set up pre-test conditions
			if tt.beforeTest != nil {
				tt.beforeTest(&f)
			}

			// Initialize the service
			tServ := NewTodoService(f.tRepo)

			// Call the method under test
			_, err := tServ.Create(tt.args.ctx, tt.args.tsr)

			// Validate results
			if (err != nil) != tt.wantErr {
				t.Errorf("Test ID %v: Create() error = %v, wantErr %v", tt.id, err, tt.wantErr)
				return
			}
		})
	}
}
func TestTodoService_GetTodoByID(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	type test struct {
		id         int
		name       string
		args       args
		beforeTest func(f *fields)
		want       *tContracts.GetTodoResponse
		wantErr    bool
	}
	tests := []test{
		{
			id:   1,
			name: "GetTodoByID_WithValidID_ReturnsSuccess",
			args: args{
				ctx: createTestContext(validTodoID),
			},
			beforeTest: func(f *fields) {
				f.tRepo.EXPECT().GetTodoByID(gomock.Any(), validTodoID).Return(&todoAgg1, nil)
			},
			want:    ToGetByIDRes(&todoAgg1),
			wantErr: false,
		},
		{
			id:   2,
			name: "GetTodoByID_WithInvalidID_ReturnsError",
			args: args{
				ctx: createTestContext(nonExistentID),
			},
			beforeTest: func(f *fields) {
				f.tRepo.EXPECT().GetTodoByID(gomock.Any(), nonExistentID).Return(nil, errInvalidTodoID) // assuming `someError` is a defined error
			},
			want:    nil,
			wantErr: true,
		},
		{
			id:   3,
			name: "GetTodoByID_WithMissingTodoID_ReturnsError",
			args: args{
				ctx: context.Background(),
			},
			beforeTest: func(f *fields) {
				//Since error Occurs Earlier
			},
			want:    nil,
			wantErr: true,
		},
	}

	// Running the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				tRepo: mock.NewMockITodoRepo(ctrl),
			}

			if tt.beforeTest != nil {
				tt.beforeTest(&f)
			}
			tServ := NewTodoService(f.tRepo)
			got, err := tServ.GetTodoByID(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ID %v GetTodoByID() error = %v, wantErr %v", tt.id, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ID %v GetTodoByID() got = %v, want %v", tt.id, got, tt.want)
			}
		})
	}
}

func TestTodoService_GetTodos(t *testing.T) {
	ctx := CreateTestGinContext(createGinContextHelper1.Method, createGinContextHelper1.getUrl())
	type args struct {
		ctx *gin.Context
	}
	type test struct {
		id         int
		name       string
		args       args
		beforeTest func(f *fields)
		want       []tContracts.GetTodoResponse
		wantErr    bool
	}
	tests := []test{
		{
			id:   1,
			name: "Get_Todos_Success",
			args: args{
				ctx: ctx,
			},
			beforeTest: func(f *fields) {
				f.tRepo.EXPECT().GetTodos(gomock.Any(), createGinContextHelper1.getRawQuery()).
					Return([]*todo_aggregate.Todo{
						&todoAgg1,
						&todoAgg2,
					}, nil)
			},
			want: []tContracts.GetTodoResponse{
				*ToGetByIDRes(&todoAgg1),
				*ToGetByIDRes(&todoAgg2),
			},
			wantErr: false,
		},
		{
			id:   2,
			name: "Get_Todos_Repository_Error",
			args: args{
				ctx: ctx,
			},
			beforeTest: func(f *fields) {
				f.tRepo.EXPECT().GetTodos(ctx, createGinContextHelper1.getRawQuery()).
					Return(nil, errRepositoryError)
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				tRepo: mock.NewMockITodoRepo(ctrl),
			}

			if tt.beforeTest != nil {
				tt.beforeTest(&f)
			}
			tServ := NewTodoService(f.tRepo)
			got, err := tServ.GetTodos(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ID %v GetTodoByID() error = %v, wantErr %v", tt.id, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ID %v GetTodoByID() got = %v, want %v", tt.id, got, tt.want)
			}
		})
	}
}

func TestTodoService_UpdateTodoByID(t *testing.T) {
	ctx := createTestContext(validTodoID)
	type args struct {
		ctx context.Context
		tsr *tContracts.UpdateTodoRequest
	}
	type test struct {
		id         int
		name       string
		args       args
		beforeTest func(f *fields)
		want       tContracts.UpdateTodoResponse
		wantErr    bool
	}
	tests := []test{
		{
			id:   1,
			name: "UpdateTodoById_ALlFieldsUpdated_Success",
			args: args{
				ctx: ctx,
				tsr: &ValidUpdateTodoRequestAllFields,
			},
			beforeTest: func(f *fields) {
				f.tRepo.EXPECT().
					UpdateTodo(gomock.Any(), validTodoID, fromUpdateTodoRequest(validTodoID, &ValidUpdateTodoRequestAllFields)).
					Return(&todoAgg1, nil)
			},
			want:    *toUpateTodoRes(&todoAgg1),
			wantErr: false,
		},
		{
			id:   2,
			name: "UpdateTodoById_SingleFieldUpdated_Success",
			args: args{
				ctx: ctx,
				tsr: &ValidUpdateTodoRequestSingleField,
			},
			beforeTest: func(f *fields) {
				f.tRepo.EXPECT().
					UpdateTodo(gomock.Any(), validTodoID, fromUpdateTodoRequest(validTodoID, &ValidUpdateTodoRequestSingleField)).
					Return(&todoAgg1, nil)
			},
			want:    *toUpateTodoRes(&todoAgg1),
			wantErr: false,
		},
		{
			id:   3,
			name: "UpdateTodoByID_NonExistentID_ReturnsError",
			args: args{
				ctx: createTestContext(nonExistentID),
				tsr: &ValidUpdateTodoRequestAllFields,
			},
			beforeTest: func(f *fields) {
				f.tRepo.EXPECT().
					UpdateTodo(gomock.Any(), nonExistentID, fromUpdateTodoRequest(nonExistentID, &ValidUpdateTodoRequestAllFields)).
					Return(nil, errInvalidTodoID)
			},
			want:    tContracts.UpdateTodoResponse{},
			wantErr: true,
		},
		{
			id:   4,
			name: "UpdateTodoByID_WithMissingTodoID_ReturnsError",
			args: args{
				ctx: context.Background(),
				tsr: &ValidUpdateTodoRequestAllFields,
			},
			beforeTest: func(f *fields) {
			},
			want:    tContracts.UpdateTodoResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				tRepo: mock.NewMockITodoRepo(ctrl),
			}

			if tt.beforeTest != nil {
				tt.beforeTest(&f)
			}
			tServ := NewTodoService(f.tRepo)
			got, err := tServ.UpdateTodoByID(tt.args.ctx, tt.args.tsr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ID %v GetTodoByID() error = %v, wantErr %v", tt.id, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ID %v GetTodoByID() got = %v, want %v", tt.id, got, tt.want)
			}
		})
	}
}
