package migrate

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/models"
	"github.com/abibby/bob/schema"
	"github.com/abibby/bob/selects"
	"github.com/abibby/bob/set"
	"github.com/abibby/bob/slices"
	"github.com/davecgh/go-spew/spew"
)

func update(m models.Model) (*schema.UpdateTableBuilder, *schema.UpdateTableBuilder) {
	err := selects.InitializeRelationships(m)
	if err != nil {
		panic(err)
	}

	tableName := builder.GetTable(m)
	fields := getFields(m)
	oldTable := Blueprint(tableName)

	changedColumns := []*schema.ColumnBuilder{}
	newColumns := []string{}
	droppedColumns := []*schema.ColumnBuilder{}

	return schema.Table(tableName, func(table *schema.Blueprint) {
			addedForeignKeys := []*selects.ForeignKey{}
			addedColumns := set.New[string]()
			for _, f := range fields {
				if f.relation != nil {
					foreignKeys := f.relation.ForeignKeys()
					for _, foreignKey := range foreignKeys {
						if slices.Has(addedForeignKeys, foreignKey) {
							continue
						}

						addedForeignKeys = append(addedForeignKeys, foreignKey)
						table.ForeignKey(foreignKey.LocalKey, foreignKey.RelatedTable, foreignKey.RelatedKey)
					}
				} else {
					c, ok := oldTable.Column(f.tag.Name)
					if ok {
						addedColumns.Add(f.tag.Name)

						if c.Matches(f.dataType, f.nullable, f.tag) {
							continue
						}
					}
					b := table.OfType(f.dataType, f.tag.Name)
					if ok {
						b.Change()
						changedColumns = append(changedColumns, c)
					} else {
						spew.Dump("drop", f.tag.Name)
						newColumns = append(newColumns, f.tag.Name)
					}
					if f.nullable {
						b.Nullable()
					}
					if f.tag.Index {
						b.Index()
					}
					if f.tag.AutoIncrement {
						b.AutoIncrement()
					}
					if f.tag.Primary {
						b.Primary()
					}
				}
			}
			for _, c := range oldTable.Columns() {
				if addedColumns.Has(c.Name()) {
					continue
				}

				droppedColumns = append(droppedColumns, c)
				table.DropColumn(c.Name())

			}
		}),
		schema.Table(tableName, func(table *schema.Blueprint) {
			for _, c := range changedColumns {
				table.AddColumn(c)
			}
			for _, c := range droppedColumns {
				table.AddColumn(c)
			}
			for _, c := range newColumns {
				table.DropColumn(c)
			}
		})
}

func Blueprint(tableName string) *schema.Blueprint {
	result := &schema.Blueprint{}
	for _, m := range migrations {
		blueprinter, ok := m.Up().(schema.Blueprinter)
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
