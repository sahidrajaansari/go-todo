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

func TestTodoService_GetTodoByID(t *testing.T) {
	ctx := getContext()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		id         int
		name       string
		args       args
		beforeTest func(f *fields)
		want       *tContracts.GetTodoResponse
		wantErr    bool
	}{
		{
			name: "Get Todo By Id -Success",
			args: args{
				ctx: ctx,
			},
			beforeTest: func(f *fields) {
				f.tRepo.EXPECT().GetTodoByID(ctx, "Valid").Return(&todo_aggregate.TodoAgg, nil)
			},
			want:    ToGetByIDRes(&todo_aggregate.TodoAgg),
			wantErr: false,
		},
		{
			name: "Error - Todo Not Found",
			args: args{
				ctx: ctx,
			},
			beforeTest: func(f *fields) {
				f.tRepo.EXPECT().GetTodoByID(ctx, "Invalid").Return(nil, "Error") // assuming `someError` is a defined error
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

			// Initialize mock repo and setup test data
			f := fields{
				tRepo: mock.NewMockITodoRepo(ctrl),
			}

			if tt.beforeTest != nil {
				tt.beforeTest(&f)
			}

			// Create the service and call the method
			tServ := NewTodoService(f.tRepo)
			got, err := tServ.GetTodoByID(tt.args.ctx)

			// Check for errors
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTodoByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check if the result matches the expected output
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTodoByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
