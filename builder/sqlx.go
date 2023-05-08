package builder

import "github.com/jmoiron/sqlx"

type QueryExecer interface {
	sqlx.QueryerContext
	sqlx.ExecerContext
}
