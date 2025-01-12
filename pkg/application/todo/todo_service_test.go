package todo

import (
	"context"
	"reflect"
	"testing"
	tContracts "todo-level-5/pkg/contract/todo"
	"todo-level-5/pkg/domain/persistence/mock"
	"todo-level-5/pkg/domain/todo_aggregate"

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
					Return(errRepositryError)
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
				f.tRepo.EXPECT().GetTodoByID(gomock.Any(), validTodoID).Return(&todoAgg, nil)
			},
			want:    ToGetByIDRes(&todoAgg),
			wantErr: false,
		},
		{
			id:   2,
			name: "GetTodoByID_WithInvalidID_ReturnsError",
			args: args{
				ctx: createTestContext(inValidTodoID),
			},
			beforeTest: func(f *fields) {
				f.tRepo.EXPECT().GetTodoByID(gomock.Any(), inValidTodoID).Return(nil, errInvalidTodoID) // assuming `someError` is a defined error
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
