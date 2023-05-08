package selects

import (
	"context"

	"github.com/abibby/bob/builder"
)

type BeforeSaver interface {
	BeforeSave(ctx context.Context, tx builder.QueryExecer) error
}
type AfterSaver interface {
	AfterSave(ctx context.Context, tx builder.QueryExecer) error
}

type AfterLoader interface {
	AfterLoad(ctx context.Context, tx builder.QueryExecer) error
}
