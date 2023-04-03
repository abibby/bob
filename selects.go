package bob

import (
	"github.com/abibby/bob/models"
	"github.com/abibby/bob/selects"
)

type Scope = selects.Scope
type ScopeFunc = selects.ScopeFunc
type Scoper = selects.Scoper

var SoftDeletes = &selects.Scope{
	Name: "soft-deletes",
	Apply: func(b *selects.SubBuilder) *selects.SubBuilder {
		return b.Where("deleted_at", "=", nil)
	},
}

func New[T models.Model]() *selects.Builder[T] {
	return selects.New[T]()
}

func From[T models.Model]() *selects.Builder[T] {
	return selects.From[T]()
}

func NewEmpty[T models.Model]() *selects.Builder[T] {
	return selects.NewEmpty[T]()
}
