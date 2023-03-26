package models

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Model interface {
	InDatabase() bool
}

type BaseModel struct {
	inDatabase bool
}

var _ Model = &BaseModel{}

func (m *BaseModel) InDatabase() bool {
	return m.inDatabase
}

func (m *BaseModel) AfterLoad(ctx context.Context, tx *sqlx.Tx) error {
	m.inDatabase = true
	return nil
}
func (m *BaseModel) AfterSave(ctx context.Context, tx *sqlx.Tx) error {
	m.inDatabase = true
	return nil
}
