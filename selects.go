package bob

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/selects"
)

func New() *selects.Builder {
	return selects.New().Select("*")
}

func From(m any) *selects.Builder {
	return selects.New().Select("*").From(builder.GetTable(m))
}

func NewEmpty() *selects.Builder {
	return selects.New()
}
