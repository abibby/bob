package selects

import (
	"context"
	"fmt"
	"reflect"

	"github.com/jmoiron/sqlx"
)

var relationType = reflect.TypeOf((*Relationship)(nil)).Elem()

func InitializeRelationships(v any) error {
	return each(v, initializeRelationships)
}

func initializeRelationships(v reflect.Value, pointer bool) error {
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		ft := t.Field(i)

		if ft.Anonymous {
			err := initializeRelationships(v.Field(i), ft.Type.Kind() == reflect.Ptr)
			if err != nil {
				return err
			}
			continue
		}

		if ft.Type.Implements(relationType) {
			fv := v.Field(i)
			if ft.Type.Kind() == reflect.Ptr {
				fv.Set(reflect.New(ft.Type.Elem()))
			} else {
				fv.Set(reflect.New(ft.Type).Elem())
			}
			iValue := v
			if pointer {
				iValue = iValue.Addr()
			}
			err := fv.Interface().(Relationship).Initialize(iValue.Interface(), ft)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func Load(tx *sqlx.Tx, v any, relation string) error {
	return LoadContext(context.Background(), tx, v, relation)
}
func LoadContext(ctx context.Context, tx *sqlx.Tx, v any, relation string) error {
	relations := []Relationship{}
	err := each(v, func(v reflect.Value, pointer bool) error {
		r, ok := getRelation(v, relation)
		if !ok {
			return fmt.Errorf("%s has no relation %s: %w", v.Type().Name(), relation, ErrMissingRelationship)
		}

		relations = append(relations, r)
		return nil
	})
	if err != nil {
		return err
	}

	if len(relations) == 0 {
		return nil
	}
	return relations[0].Load(ctx, tx, relations)
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

func each(v any, cb func(v reflect.Value, pointer bool) error) error {
	pointer := false
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		pointer = true
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

	return cb(rv, pointer)
}
