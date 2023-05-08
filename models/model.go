package models

import (
	"context"

	"github.com/abibby/bob/builder"
)

type Contexter interface {
	Context() context.Context
}

type Model interface {
	InDatabase() bool
}

type BaseModel struct {
	inDatabase bool
	ctx        context.Context
}

var _ Model = &BaseModel{}

func (m *BaseModel) InDatabase() bool {
	return m.inDatabase
}
func (m *BaseModel) Context() context.Context {
	return m.ctx
}

func (m *BaseModel) AfterLoad(ctx context.Context, tx builder.QueryExecer) error {
	m.inDatabase = true
	m.ctx = ctx
	return nil
}
func (m *BaseModel) AfterSave(ctx context.Context, tx builder.QueryExecer) error {
	m.inDatabase = true
	m.ctx = ctx
	return nil
}
