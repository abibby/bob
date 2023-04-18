package builder

import (
	"reflect"
	"strings"
)

type Tag struct {
	Name          string
	Primary       bool
	AutoIncrement bool
	Readonly      bool
	Index         bool
}

func DBTag(f reflect.StructField) *Tag {
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
		for i := 0; i < tagType.NumField(); i++ {
			f := tagType.Field(i)
			if strings.ToLower(f.Name) == p {
				tagValue.Field(i).SetBool(true)
			}
		}
	}
	return tag
}
