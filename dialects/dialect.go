package dialects

type DataType string

const (
	DataTypeString          = DataType("string")
	DataTypeInteger         = DataType("int")
	DataTypeUnsignedInteger = DataType("uint")
	DataTypeFloat           = DataType("float")
	DataTypeBoolean         = DataType("bool")
	DataTypeJSON            = DataType("json")
	DataTypeDate            = DataType("date")
	DataTypeDateTime        = DataType("date-time")
)

type Dialect interface {
	Identifier(string) string
	DataType(DataType) string
}

type unsetDialect struct{}

func (*unsetDialect) Identifier(s string) string {
	return s
}

func (*unsetDialect) DataType(t DataType) string {
	return string(t)
}

var DefaultDialect Dialect = &unsetDialect{}
