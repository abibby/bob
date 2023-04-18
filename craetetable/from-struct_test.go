package craetetable_test

import (
	"testing"

	"github.com/abibby/bob/craetetable"
	"github.com/abibby/bob/models"
	"github.com/abibby/bob/test"
	"github.com/abibby/nulls"
)

func TestFromStruct(t *testing.T) {
	type A struct {
		models.BaseModel
		ID       int           `db:"id,primary"`
		Nullable *nulls.String `db:"nullable"`
		Indexed  bool          `db:"indexed,index"`
	}
	test.QueryTest(t, []test.Case{
		{
			Name:    "create table",
			Builder: craetetable.FromStruct(&test.Foo{}),
			ExpectedSQL: "CREATE TABLE \"foos\" (\"id\" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, \"name\" TEXT NOT NULL); " +
				"CONSTRAINT \"bars\" FOREIGN KEY (\"id\") REFERENCES \"bars\"(\"foo_id\");",
			ExpectedBindings: []any{},
		},
		{
			Name:    "create table",
			Builder: craetetable.FromStruct(&A{}),
			ExpectedSQL: "CREATE TABLE \"as\" (\"id\" INTEGER PRIMARY KEY NOT NULL, \"nullable\" TEXT, \"indexed\" INTEGER NOT NULL); " +
				"CREATE INDEX IF NOT EXIST \"as-indexed\" ON \"as\" (\"indexed\");",
			ExpectedBindings: []any{},
		},
	})
}
