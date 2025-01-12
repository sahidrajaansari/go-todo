package todo

import (
	"context"
	"reflect"
	"testing"
	tContracts "todo-level-5/pkg/contract/todo"
	"todo-level-5/pkg/domain/persistence/mock"

	"github.com/golang/mock/gomock"
)

type fields struct {
	tRepo *mock.MockITodoRepo
}

func TestTodoService_GetTodoByID(t *testing.T) {

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
