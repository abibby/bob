package migrate

import (
	"fmt"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/models"
	"github.com/abibby/bob/schema"
	"github.com/abibby/bob/selects"
)

var (
	ErrNoChanges = fmt.Errorf("no changes")
)

func (m *Migrations) update(model models.Model) (*schema.UpdateTableBuilder, *schema.UpdateTableBuilder, error) {
	err := selects.InitializeRelationships(model)
	if err != nil {
		return nil, nil, err
	}

	tableName := builder.GetTable(model)
	fields, err := getFields(model)
	if err != nil {
		return nil, nil, err
	}
	oldTable := m.Blueprint(tableName)
	newTable := blueprintFromFields(tableName, fields)

	hasChanges := false

	up := schema.Table(tableName, func(table *schema.Blueprint) {
		hasChanges = table.Update(oldTable, newTable)
	})
	down := schema.Table(tableName, func(table *schema.Blueprint) {
		table.Update(newTable, oldTable)
	})
	if !hasChanges {
		return nil, nil, ErrNoChanges
	}

	return up, down, nil
}
