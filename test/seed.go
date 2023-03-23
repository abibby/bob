package test

import (
	"database/sql"
	"fmt"
)

var fooID = 1

func CreateFoo(tx *sql.Tx) *Foo {
	f := &Foo{
		ID:   fooID,
		Name: fmt.Sprintf("name %d", fooID),
	}
	fooID++

	_, err := tx.Exec("INSERT INTO foo (id, name) values (?,?)", f.ID, f.Name)
	if err != nil {
		panic(err)
	}

	return f
}
