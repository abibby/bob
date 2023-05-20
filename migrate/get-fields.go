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
				case reflect.Int8:
					f.dataType = dialects.DataTypeInt8
				case reflect.Int16:
					f.dataType = dialects.DataTypeInt16
				case reflect.Int32:
					f.dataType = dialects.DataTypeInt32
				case reflect.Int, reflect.Int64:
					f.dataType = dialects.DataTypeInt32
				case reflect.Uint8:
					f.dataType = dialects.DataTypeUInt8
				case reflect.Uint16:
					f.dataType = dialects.DataTypeUInt16
				case reflect.Uint32:
					f.dataType = dialects.DataTypeUInt32
				case reflect.Uint, reflect.Uint64:
					f.dataType = dialects.DataTypeUInt32
				case reflect.Float32:
					f.dataType = dialects.DataTypeFloat32
				case reflect.Float64:
					f.dataType = dialects.DataTypeFloat64
				case reflect.String:
					f.dataType = dialects.DataTypeString
				case reflect.Map, reflect.Slice, reflect.Struct:
					f.dataType = dialects.DataTypeJSON
				case reflect.Array:
					f.dataType = dialects.DataTypeBlob
				default:
					return fmt.Errorf("no datatype for %v", t.Kind())
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
