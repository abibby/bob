package bob

import (
	"github.com/abibby/bob/models"
	"github.com/abibby/bob/selects"
)

func New[T models.Model]() *selects.Builder[T] {
	return selects.New[T]()
}

func From[T models.Model]() *selects.Builder[T] {
	return selects.From[T]()
}

func NewEmpty[T models.Model]() *selects.Builder[T] {
	return selects.NewEmpty[T]()
}
