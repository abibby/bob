package selects_test

import (
	"github.com/abibby/bob/insert"
	"github.com/abibby/bob/models"
	"github.com/abibby/bob/selects"
	"github.com/abibby/bob/test"
	"github.com/jmoiron/sqlx"
)

func NewTestBuilder() *selects.Builder[*test.Foo] {
	return selects.New[*test.Foo]().Select("*").From("foo")
}

func MustSave(tx *sqlx.Tx, v models.Model) {
	err := insert.Save(tx, v)
	if err != nil {
		panic(err)
	}
}
