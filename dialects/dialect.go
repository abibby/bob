package dialects

import (
	"github.com/abibby/bob/set"
)

type DataType string

const (
	DataTypeBlob            = DataType("blob")
	DataTypeBoolean         = DataType("bool")
	DataTypeDate            = DataType("date")
	DataTypeDateTime        = DataType("date-time")
	DataTypeFloat           = DataType("float")
	DataTypeInteger         = DataType("int")
	DataTypeJSON            = DataType("json")
	DataTypeString          = DataType("string")
	DataTypeUnsignedInteger = DataType("uint")
)

var dataTypes = set.Set[DataType]{
	DataTypeBlob:            struct{}{},
	DataTypeBoolean:         struct{}{},
	DataTypeDate:            struct{}{},
	DataTypeDateTime:        struct{}{},
	DataTypeFloat:           struct{}{},
	DataTypeInteger:         struct{}{},
	DataTypeJSON:            struct{}{},
	DataTypeString:          struct{}{},
	DataTypeUnsignedInteger: struct{}{},
}

func (d DataType) IsValid() bool {
	return dataTypes.Has(d)
}

// DataTyper must not be implemented on an interface
type DataTyper interface {
	DataType() DataType
}

type Dialect interface {
	Identifier(string) string
	DataType(DataType) string
	CurrentTime() string
	TableQuery() string
}

type unsetDialect struct{}

func (*unsetDialect) Identifier(s string) string {
	return s
}

func (*unsetDialect) DataType(t DataType) string {
	return string(t)
}

func (*unsetDialect) CurrentTime() string {
	return "CURRENT_TIMESTAMP"
}

func (*unsetDialect) TableQuery() string {
	return ""
}

var DefaultDialect Dialect = &unsetDialect{}
