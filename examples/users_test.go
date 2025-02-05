package examples

import (
	"context"
	"errors"
	"github.com/dal-go/dalgo/dal"
	"github.com/dal-go/mocks4dalgo/mock_dal"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

func TestSelectUserByEmail(t *testing.T) {
	type args struct {
		ctx   context.Context
		db    dal.ReadSession
		email string
	}
	mockCtrl := gomock.NewController(t)
	dbMock := mock_dal.NewMockDB(mockCtrl)
	tests := []struct {
		name         string
		args         args
		selectResult mock_dal.SelectResult
		want         *userData
		wantErr      error
	}{
		{
			name: "should return nil",
			args: args{
				db:    dbMock,
				ctx:   context.Background(),
				email: "unknown@example.com",
			},
			selectResult: mock_dal.NewSelectResult(
				nil,
				dal.ErrRecordNotFound,
			),
			want:    nil,
			wantErr: dal.ErrRecordNotFound,
		},
		{
			name: "should succeed",
			args: args{
				db:    dbMock,
				ctx:   context.Background(),
				email: "test@example.com",
			},
			selectResult: mock_dal.NewSelectResult(
				mock_dal.NewRecordsReader(0,
					dal.NewRecordWithData(
						dal.NewKeyWithID("users", "user1"),
						&userData{Email: "test@example.com"},
					).SetError(nil),
				),
				nil,
			),
			want:    &userData{Email: "test@example.com"},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbMock.EXPECT().
				QueryReader(gomock.Any(), gomock.Any()).
				DoAndReturn(func(ctx context.Context, query dal.Query) (dal.Reader, error) {
					return tt.selectResult.Reader, tt.selectResult.Err
				})
			userRecord, err := SelectUserByEmail(tt.args.ctx, tt.args.db, tt.args.email)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("SelectUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if got := userRecord.Data().(*userData); !reflect.DeepEqual(*got, *tt.want) {
					t.Errorf("SelectUserByEmail() returned %T=%+v, want %T=%+v", got, got, tt.want, tt.want)
				}
			}
		})
	}
}
