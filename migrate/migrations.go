package migrate

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/models"
	"github.com/abibby/bob/schema"
	"github.com/davecgh/go-spew/spew"
)

type Migrations struct {
	migrations []*Migration
}

func New() *Migrations {
	return &Migrations{
		migrations: []*Migration{},
	}
}

func (m *Migrations) Add(migration *Migration) {
	m.migrations = append(m.migrations, migration)
}

func (m *Migrations) isTableCreated(table string) bool {
	for _, m := range m.migrations {
		blueprinter, ok := m.Up().(schema.Blueprinter)
		if !ok {
			spew.Dump("not ok")
			continue
		}
		if blueprinter.Type() == schema.BlueprintTypeCreate {
			if blueprinter.GetBlueprint().TableName() == table {
				return true
			}
		}
	}

	return false
}

func (m *Migrations) GenerateMigration(migrationName, packageName string, model models.Model) (string, error) {
	if !m.isTableCreated(builder.GetTable(model)) {
		return SrcFile(migrationName, packageName, create(model), drop(model))
	}

	up, down, err := m.update(model)
	if err != nil {
		return "", err
	}
	return SrcFile(migrationName, packageName, up, down)
}

func (m *Migrations) Blueprint(tableName string) *schema.Blueprint {
	result := &schema.Blueprint{}

	for _, migration := range m.migrations {
		blueprinter, ok := migration.Up().(schema.Blueprinter)
		if !ok {
			continue
		}
		blueprint := blueprinter.GetBlueprint()
		if blueprint.TableName() != tableName {
			continue
		}

		if blueprinter.Type() == schema.BlueprintTypeCreate {
			result = blueprint
		} else {
			result.Merge(blueprint)
		}
	}
	return result
}
