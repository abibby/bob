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
	InsertOrUpdate(table, primaryKey string, columns []string, values []any) (string, []any)
}

type unsetDialect struct{}

func (*unsetDialect) Identifier(s string) string {
	return s
}

func (*unsetDialect) DataType(t DataType) string {
	return string(t)
}
func (d *unsetDialect) InsertOrUpdate(table, primaryKey string, columns []string, values []any) (string, []any) {
	return "", nil
}

var DefaultDialect Dialect = &unsetDialect{}
