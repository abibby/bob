package dialects

type DataType string

const (
	DataTypeString  = DataType("string")
	DataTypeInteger = DataType("int")
	DataTypeFloat   = DataType("float")
)

type Dialect interface {
	Identifier(string) string
	DataType(DataType) string
}
