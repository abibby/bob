package dialects

type Dialect interface {
	Identifier(string) string
}
