package examples

import (
	"context"
	"github.com/dal-go/dalgo/dal"
	"github.com/dal-go/dalgo/orm"
	"reflect"
)

type user struct {
	Email     orm.FieldDefinition[string]
	FirstName orm.FieldDefinition[string]
	LastName  orm.FieldDefinition[string]
}

// User defines user collection
var User = user{
	FirstName: orm.NewField[string]("fist_name"),
	LastName:  orm.NewField[string]("last_name"),
	Email:     orm.NewField[string]("email"),
}

func (v user) Collection() *dal.CollectionRef {
	return &dal.CollectionRef{
		Name: "users",
	}
}

type userData struct {
	Email string `json:"email"`
}

// SelectUserByEmail is a examples facade method
func SelectUserByEmail(ctx context.Context, db dal.ReadSession, email string) (record dal.Record, err error) {
	if db == nil {
		panic("db is a required parameter")
	}
	q := dal.
		From(User.Collection().Name).
		WhereField("Email", dal.Equal, User.Email.EqualTo(email)).
		Limit(1).
		SelectInto(func() dal.Record {
			return dal.NewRecordWithIncompleteKey(User.Collection().Name, reflect.String, &userData{})
		})
	reader, err := db.QueryReader(ctx, q)
	if err != nil {
		return nil, err
	}
	if reader == nil {
		panic("db.Select() returned no error and nil reader")
	}
	return reader.Next()
}
