package bobtesting

import (
	"context"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type Runner[T any] interface {
	Run(name string, cb func(t T)) bool
}

var migrate func(db *sqlx.DB) error

func SetMigrate(cb func(db *sqlx.DB) error) {
	migrate = cb
}

func RunWithDatabase[T testing.TB](t T, name string, cb func(t T, tx *sqlx.Tx)) bool {
	var err error
	if db == nil {
		if migrate == nil {
			panic(fmt.Errorf("no migrate function defined call bobtesting.SetMigrate() first"))
		}
		db, err = sqlx.Open("sqlite3", ":memory:")
		if err != nil {
			panic(fmt.Errorf("failed to open database: %w", err))
		}
		err = migrate(db)
		if err != nil {
			panic(fmt.Errorf("failed to create tables: %w", err))
		}
	}
	tx, err := db.BeginTxx(context.Background(), nil)
	if err != nil {
		panic(fmt.Errorf("failed to begin transaction: %w", err))
	}

	var tAny any = t
	result := tAny.(Runner[T]).Run(name, func(t T) {
		cb(t, tx)
	})

	err = tx.Rollback()
	if err != nil {
		panic(fmt.Errorf("failed to rollback transaction: %w", err))
	}
	return result
}
