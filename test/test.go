package test

import (
	"context"
	"testing"

	"github.com/abibby/bob/bobtest"
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
	"github.com/abibby/bob/dialects/sqlite"
	"github.com/abibby/bob/migrate"
	"github.com/abibby/bob/models"
	"github.com/abibby/bob/selects"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	Name             string
	Builder          builder.ToSQLer
	ExpectedSQL      string
	ExpectedBindings []any
}

func QueryTest(t *testing.T, testCases []Case) {

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			q, bindings, err := tc.Builder.ToSQL(dialects.New())
			assert.NoError(t, err)

			assert.Equal(t, tc.ExpectedSQL, q)
			assert.Equal(t, tc.ExpectedBindings, bindings)
		})
	}
}

var runner = bobtest.NewRunner(func() (*sqlx.DB, error) {
	sqlite.UseSQLite()
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	err = migrate.RunModelCreate(ctx, db, &Foo{}, &Bar{})
	if err != nil {
		return nil, err
	}
	return db, nil
})

var Run = runner.Run
var RunBenchmark = runner.RunBenchmark

type Foo struct {
	models.BaseModel
	ID   int                    `json:"id"   db:"id,primary,autoincrement"`
	Name string                 `json:"name" db:"name"`
	Bar  *selects.HasOne[*Bar]  `json:"bar"`
	Bars *selects.HasMany[*Bar] `json:"bars"`
}

func (h *Foo) Table() string {
	return "foos"
}

type Bar struct {
	models.BaseModel
	ID    int                      `json:"id"     db:"id,primary,autoincrement"`
	FooID int                      `json:"foo_id" db:"foo_id"`
	Foo   *selects.BelongsTo[*Foo] `json:"foo"`
}

func (h *Bar) Table() string {
	return "bars"
}
