package bob

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/models"
	"github.com/abibby/bob/selects"
)

func New[T models.Model]() *selects.Builder[T] {
	return selects.New[T]().Select("*")
}

func From[T models.Model](m T) *selects.Builder[T] {
	return selects.New[T]().Select("*").From(builder.GetTable(m))
}

func NewEmpty[T models.Model]() *selects.Builder[T] {
	return selects.New[T]()
}
