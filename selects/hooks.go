package selects

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type BeforeSaver interface {
	BeforeSave(ctx context.Context, tx *sqlx.Tx) error
}
type AfterSaver interface {
	AfterSave(ctx context.Context, tx *sqlx.Tx) error
}

type AfterLoader interface {
	AfterLoad(ctx context.Context, tx *sqlx.Tx) error
}
