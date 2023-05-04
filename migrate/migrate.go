package migrate

import (
	"fmt"
	"go/format"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/models"
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
	migrate.New(migrate.Migration{
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
	src, err := format.Source(src)
	if err != nil {
		panic(err)
	}
	src, err = imports.Process(outFile, src, nil)
	if err != nil {
		panic(err)
	}

	return string(src), nil
}

func GenerateMigration(migrationName, packageName string, m models.Model) (string, error) {
	return SrcFile(migrationName, packageName, create(m), drop(m))
}

func New(m Migration) {

}
