package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
	_ "github.com/abibby/bob/dialects/sqlite"
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
			q, bindings, err := tc.Builder.ToSQL(dialects.DefaultDialect)
			assert.NoError(t, err)

			assert.Equal(t, tc.ExpectedSQL, q)
			assert.Equal(t, tc.ExpectedBindings, bindings)
		})
	}
}

type Foo struct {
	ID   int                    `db:"id"   json:"id"`
	Name string                 `db:"name" json:"name"`
	Bar  *selects.HasOne[*Bar]  `db:"-"    json:"bar"`
	Bars *selects.HasMany[*Bar] `db:"-"    json:"bars"`
}

type Bar struct {
	ID    int                      `db:"id"     json:"id"`
	FooID int                      `db:"foo_id" json:"foo_id"`
	Foo   *selects.BelongsTo[*Foo] `db:"-"      json:"foo"`
}

const createTables = `CREATE TABLE foos (
	id int not null,
	name varchar(255) not null default '',
	PRIMARY KEY (id)
);
CREATE TABLE bars (
	id int not null,
	foo_id int not null,
	PRIMARY KEY (id)
);`

func WithDatabase(cb func(tx *sqlx.Tx)) {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		panic(fmt.Errorf("failed to open database: %w", err))
	}
	_, err = db.Exec(createTables)
	if err != nil {
		panic(fmt.Errorf("failed to create tables: %w", err))
	}
	tx, err := db.BeginTxx(context.Background(), nil)
	if err != nil {
		panic(fmt.Errorf("failed to begin transaction: %w", err))
	}

	cb(tx)

	err = tx.Rollback()
	if err != nil {
		panic(fmt.Errorf("failed to rollback transaction: %w", err))
	}
}
