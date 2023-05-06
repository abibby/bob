package migrate

import (
	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/models"
	"github.com/abibby/bob/schema"
	"github.com/abibby/bob/selects"
	"github.com/abibby/bob/slices"
)

func create(m models.Model) *schema.CreateTableBuilder {
	err := selects.InitializeRelationships(m)
	if err != nil {
		panic(err)
	}

	tableName := builder.GetTable(m)
	fields := getFields(m)
	return schema.Create(tableName, func(table *schema.Blueprint) {
		addedForeignKeys := []*selects.ForeignKey{}
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
				b := table.OfType(f.dataType, f.tag.Name)
				if f.nullable {
					b.Nullable()
				}
				if f.tag.Index {
					b.Index()
					// table.Index(fmt.Sprintf("%s-%s", tableName, field.tag.Name)).AddColumn(field.tag.Name)
				}
				if f.tag.AutoIncrement {
					b.AutoIncrement()
				}
				if f.tag.Primary {
					b.Primary()
				}
			}
		}
	})
}