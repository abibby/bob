package migrate

import (
	"fmt"
	"go/format"

	"github.com/abibby/bob/schema"
	"golang.org/x/tools/imports"
)

type Migration struct {
	Name string
	Up   schema.Runner
	Down schema.Runner
}

func srcFile(migrationName, packageName string, up, down ToGoer) (string, error) {
	outFile := "test.go"
	initSrc := `package %s
	
import (
	"github.com/abibby/bob/migrate"
	"github.com/abibby/bob/builder"
)

func init() {
	migrations.Add(&migrate.Migration{
		Name: %#v,
		Up: %s,
		Down: %s,
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
