package migrate

import (
	"fmt"
	"go/format"

	"github.com/abibby/bob/builder"
	"golang.org/x/tools/imports"
)

type Migration struct {
	Name string
	Up   func() builder.ToSQLer
	Down func() builder.ToSQLer
}

func SrcFile(migrationName, packageName string, up, down ToGoer) (string, error) {
	outFile := "test.go"
	initSrc := `package %s
	
import (
	"github.com/abibby/bob/migrate"
	"github.com/abibby/bob/builder"
)

func init() {
	migrations.Add(&migrate.Migration{
		Name: %#v,
		Up: func() builder.ToSQLer {
			return %s
		},
		Down: func() builder.ToSQLer {
			return %s
		},
	})
}`

	src := []byte(fmt.Sprintf(initSrc, packageName, migrationName, up.ToGo(), down.ToGo()))
	src, err := imports.Process(outFile, src, nil)
	if err != nil {
		panic(err)
	}

	src, err = format.Source(src)
	if err != nil {
		panic(err)
	}
	return string(src), nil
}
