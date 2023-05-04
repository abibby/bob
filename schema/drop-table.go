package schema

import "github.com/abibby/bob/builder"

func Drop(table string) builder.ToSQLer {
	return builder.Concat(builder.Raw("DROP TABLE "), builder.Identifier(table))
}
func DropIfExists(table string) builder.ToSQLer {
	return builder.Concat(builder.Raw("DROP TABLE IF EXISTS "), builder.Identifier(table))
}
