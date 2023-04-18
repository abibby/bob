package craetetable

import (
	"fmt"
	"reflect"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/dialects"
	"github.com/abibby/bob/models"
	"github.com/abibby/bob/selects"
)

type field struct {
	dataType dialects.DataType
	tag      *builder.Tag
	nullable bool
	relation selects.Relationship
}

func getFields(m models.Model) []*field {
	fields := []*field{}
	relationshipInterface := reflect.TypeOf((*selects.Relationship)(nil)).Elem()
	err := builder.EachField(reflect.ValueOf(m), func(sf reflect.StructField, fv reflect.Value) error {
		if !sf.IsExported() {
			return nil
		}
		if sf.Type.Implements(relationshipInterface) {
			fields = append(fields, &field{
				relation: fv.Interface().(selects.Relationship),
			})

			return nil
		}
		f := &field{
			tag:      builder.DBTag(sf),
			nullable: false,
		}
		t := sf.Type
		if t.Kind() == reflect.Pointer {
			t = t.Elem()
			f.nullable = true
		}
		switch t.Kind() {
		case reflect.Bool:
			f.dataType = dialects.DataTypeBoolean
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			f.dataType = dialects.DataTypeUnsignedInteger
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			f.dataType = dialects.DataTypeInteger
		case reflect.Float32, reflect.Float64:
			f.dataType = dialects.DataTypeFloat
		case reflect.String:
			f.dataType = dialects.DataTypeString
		case reflect.Array, reflect.Map, reflect.Slice, reflect.Struct:
			f.dataType = dialects.DataTypeJSON
			// case reflect.Complex64, reflect.Complex128:
		}

		fields = append(fields, f)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return fields
}

func FromStruct(m models.Model) *Builder {
	err := selects.InitializeRelationships(m)
	if err != nil {
		panic(err)
	}
	// relationshipInterface := reflect.TypeOf((*selects.Relationship)(nil)).Elem()
	tableName := builder.GetTable(m)
	fields := getFields(m)
	return CreateTable(tableName, func(table *Table) {
		for _, field := range fields {
			if field.relation != nil {
				for _, foreignKey := range field.relation.ForeignKeys() {
					table.ForeignKey(foreignKey.LocalKey, foreignKey.RelatedTable, foreignKey.RelatedKey)
				}
			} else {
				b := table.OfType(field.dataType, field.tag.Name).withTag(field.tag)
				if field.nullable {
					b.Nullable()
				}
				if field.tag.Index {
					table.Index(fmt.Sprintf("%s-%s", tableName, field.tag.Name)).AddColumn(field.tag.Name)
				}
			}
		}
	})
}

// func initialize[T any](v T) T {
// 	rv := reflect.ValueOf(v)
// 	if rv.Kind() != reflect.Pointer || !rv.IsNil() {
// 		return v
// 	}
// 	return reflect.New(rv.Type().Elem()).Interface().(T)
// }
