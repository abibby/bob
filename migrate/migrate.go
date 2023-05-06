package migrate

import (
	"fmt"
	"go/format"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/models"
	"github.com/abibby/bob/schema"
	"golang.org/x/tools/imports"
)

type Migration struct {
	Name string
	Up   func() builder.ToSQLer
	Down func() builder.ToSQLer
}

var migrations = []*Migration{}

func SrcFile(migrationName, packageName string, up, down ToGoer) (string, error) {
	outFile := "test.go"
	initSrc := `package %s
	
import (
	"github.com/abibby/bob/migrate"
	"github.com/abibby/bob/builder"
)

func init() {
	migrate.Add(&migrate.Migration{
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

func GenerateMigration(migrationName, packageName string, m models.Model) (string, error) {
	if !isTableCreated(builder.GetTable(m)) {
		return SrcFile(migrationName, packageName, create(m), drop(m))
	}

	up, down, err := update(m)
	if err != nil {
		return "", err
	}
	return SrcFile(migrationName, packageName, up, down)
}

func isTableCreated(table string) bool {
	for _, m := range migrations {
		up := m.Up()
		if create, ok := up.(*schema.CreateTableBuilder); ok {
			if create.TableName() == table {
				return true
			}
		}
	}

	return false
}

func Add(m *Migration) {
	migrations = append(migrations, m)
}
