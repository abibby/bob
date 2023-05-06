package migrate

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

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

var dataTyperInterface = reflect.TypeOf((*dialects.DataTyper)(nil)).Elem()

func getFields(m models.Model) ([]*field, error) {
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
		tag := builder.DBTag(sf)
		if tag.Name == "-" {
			return nil
		}

		f := &field{
			tag:      tag,
			nullable: false,
		}
		t := sf.Type
		if t.Kind() == reflect.Pointer {
			t = t.Elem()
			fv = reflect.New(t).Elem()
			f.nullable = true
		}

		if tag.Type != "" {
			if !tag.Type.IsValid() {
				return fmt.Errorf("data type %s is not valid", tag.Type)
			}
			f.dataType = tag.Type
		} else {
			switch field := fv.Interface().(type) {
			case dialects.DataTyper:
				f.dataType = field.DataType()
			case time.Time:
				f.dataType = dialects.DataTypeDateTime
			case []byte:
				f.dataType = dialects.DataTypeBlob
			case json.RawMessage:
				f.dataType = dialects.DataTypeJSON
			default:
				switch t.Kind() {
				case reflect.Bool:
					f.dataType = dialects.DataTypeBoolean
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					f.dataType = dialects.DataTypeInteger
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					f.dataType = dialects.DataTypeUnsignedInteger
				case reflect.Float32, reflect.Float64:
					f.dataType = dialects.DataTypeFloat
				case reflect.String:
					f.dataType = dialects.DataTypeString
				case reflect.Map, reflect.Slice, reflect.Struct:
					f.dataType = dialects.DataTypeJSON
				case reflect.Array:
					f.dataType = dialects.DataTypeBlob
					// case reflect.Complex64, reflect.Complex128:
				}
			}
		}

		fields = append(fields, f)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return fields, nil
}
