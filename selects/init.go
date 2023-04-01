package selects

import (
	"reflect"

	"github.com/jmoiron/sqlx"
)

func InitializeRelationships(v any) error {
	relationType := reflect.TypeOf((*Relationship)(nil)).Elem()
	return each(v, func(rv reflect.Value) error {
		rt := rv.Type()
		for i := 0; i < rv.NumField(); i++ {
			ft := rt.Field(i)

			if ft.Type.Implements(relationType) {
				fv := rv.Field(i)
				if ft.Type.Kind() == reflect.Ptr {
					fv.Set(reflect.New(ft.Type.Elem()))
				} else {
					fv.Set(reflect.New(ft.Type).Elem())
				}
				err := fv.Interface().(Relationship).Initialize(rv.Interface(), ft)
				if err != nil {
					return err
				}
			}

		}
		return nil
	})
}

func Load(tx *sqlx.Tx, v any, relation string) error {
	relations := []Relationship{}
	err := each(v, func(v reflect.Value) error {
		relation, ok := getRelation(v.Interface(), relation)
		if !ok {
			return nil
		}

		relations = append(relations, relation)
		return nil
	})
	if err != nil {
		return err
	}

	if len(relations) == 0 {
		return nil
	}
	return relations[0].Load(tx, relations)
}

func ofType[T Relationship](relations []Relationship) []T {
	relationsOfType := make([]T, 0, len(relations))
	for _, r := range relations {
		rOfType, ok := r.(T)
		if ok {
			relationsOfType = append(relationsOfType, rOfType)
		}
	}
	return relationsOfType
}

func each(v any, cb func(reflect.Value) error) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Slice {
		for i := 0; i < rv.Len(); i++ {
			err := each(rv.Index(i).Interface(), cb)
			if err != nil {
				return err
			}
		}
		return nil
	}
	if rv.Kind() != reflect.Struct {
		return nil
	}

	return cb(rv)
}
