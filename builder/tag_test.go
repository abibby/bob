package builder_test

import (
	"reflect"
	"testing"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	type Foo struct {
		ID  int
		Foo string `db:"foo,autoincrement,primary,type:date"`
	}

	rt := reflect.TypeOf(Foo{})
	assert.Equal(
		t,
		&builder.Tag{
			Name:          "ID",
			Primary:       false,
			AutoIncrement: false,
			Readonly:      false,
			Index:         false,
		},
		builder.DBTag(rt.Field(0)),
	)
	assert.Equal(
		t,
		&builder.Tag{
			Name:          "foo",
			Primary:       true,
			AutoIncrement: true,
			Readonly:      false,
			Index:         false,
			Type:          dialects.DataType("date"),
		},
		builder.DBTag(rt.Field(1)),
	)
}
