package selects_test

import (
	"github.com/abibby/bob/insert"
	"github.com/abibby/bob/selects"
	"github.com/jmoiron/sqlx"
)

func NewTestBuilder() *selects.Builder {
	return selects.New().Select("*").From("foo")
}

func MustSave(tx *sqlx.Tx, v any) {
	err := insert.Save(tx, v)
	if err != nil {
		panic(err)
	}
}
