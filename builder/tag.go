package builder

import (
	"reflect"
	"strings"

	"github.com/abibby/bob/dialects"
)

type Tag struct {
	Name          string
	Primary       bool
	AutoIncrement bool
	Readonly      bool
	Index         bool
	Type          dialects.DataType
}

// var relationshipInterface = reflect.TypeOf((*selects.Relationship)(nil)).Elem()

func DBTag(f reflect.StructField) *Tag {
	if f.Type.Kind() == reflect.Struct {
		return &Tag{
			Name: "-",
		}
	}
	dbTag, ok := f.Tag.Lookup("db")
	if !ok {
		return &Tag{
			Name: f.Name,
		}
	}
	parts := strings.Split(dbTag, ",")
	tag := &Tag{
		Name: parts[0],
	}

	tagValue := reflect.ValueOf(tag).Elem()
	tagType := tagValue.Type()
	for _, p := range parts[1:] {
		typePrefix := "type:"
		if strings.HasPrefix(p, typePrefix) {
			tag.Type = dialects.DataType(p[len(typePrefix):])
			continue
		}
		for i := 0; i < tagType.NumField(); i++ {
			f := tagType.Field(i)
			if strings.ToLower(f.Name) == p && f.Type.Kind() == reflect.Bool {
				tagValue.Field(i).SetBool(true)
			}
		}
	}

	return tag
}
