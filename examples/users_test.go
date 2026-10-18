package examples

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/dal-go/dalgo/dal"
	"github.com/dal-go/dalgo/mocks/mock_dal"
	"github.com/dal-go/record"
	"go.uber.org/mock/gomock"
)

func TestSelectUserByEmail(t *testing.T) {
	type args struct {
		ctx   context.Context
		db    dal.ReadSession
		email string
	}
	mockCtrl := gomock.NewController(t)
	dbMock := mock_dal.NewMockReadSession(mockCtrl)
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
				record.ErrRecordNotFound,
			),
			want:    nil,
			wantErr: record.ErrRecordNotFound,
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
					record.NewRecordWithData(
						record.NewKeyWithID("users", "user1"),
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
				ExecuteQueryToRecordsReader(gomock.Any(), gomock.Any()).
				DoAndReturn(func(ctx context.Context, query dal.Query) (dal.RecordsReader, error) {
					if tt.selectResult.Reader == nil {
						return nil, tt.selectResult.Err
					}
					return tt.selectResult.Reader.(dal.RecordsReader), tt.selectResult.Err
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

func TestUser_RecordWithIncompleteKey(t *testing.T) {
	recFn := User.RecordWithIncompleteKey()
	rec := recFn()
	if rec == nil {
		t.Fatal("expected non-nil record")
	}
	if rec.Key() == nil {
		t.Fatal("expected non-nil key")
	}
}

func TestSelectUserByEmail_NilDB(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic on nil db")
		}
	}()
	_, _ = SelectUserByEmail(context.Background(), nil, "test@example.com")
}

func TestSelectUserByEmail_NilReader(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	dbMock := mock_dal.NewMockReadSession(mockCtrl)
	dbMock.EXPECT().
		ExecuteQueryToRecordsReader(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic on nil reader with no error")
		}
	}()
	_, _ = SelectUserByEmail(context.Background(), dbMock, "test@example.com")
}
