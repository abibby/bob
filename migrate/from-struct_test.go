package migrate_test

import (
	"fmt"
	"testing"

	"github.com/abibby/bob/migrate"
	"github.com/abibby/bob/models"
	"github.com/abibby/nulls"
)

func TestFromStruct(t *testing.T) {
	type TestModel struct {
		models.BaseModel
		ID       int           `db:"id,primary"`
		Nullable *nulls.String `db:"nullable"`
		Indexed  bool          `db:"indexed,index"`
	}
	fmt.Println(migrate.GenerateMigration("2023-01-01T00:00:00Z create test model", "packageName", &TestModel{}))
	t.Fail()
	// test.QueryTest(t, []test.Case{
	// 	{
	// 		Name:    "create table",
	// 		Builder: migrate.FromStruct(&test.Foo{}),
	// 		ExpectedSQL: "CREATE TABLE \"foos\" (\"id\" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, \"name\" TEXT NOT NULL); " +
	// 			"CONSTRAINT \"bars\" FOREIGN KEY (\"id\") REFERENCES \"bars\"(\"foo_id\");",
	// 		ExpectedBindings: []any{},
	// 	},
	// 	{
	// 		Name:    "create table",
	// 		Builder: migrate.FromStruct(&A{}),
	// 		ExpectedSQL: "CREATE TABLE \"as\" (\"id\" INTEGER PRIMARY KEY NOT NULL, \"nullable\" TEXT, \"indexed\" INTEGER NOT NULL); " +
	// 			"CREATE INDEX IF NOT EXIST \"as-indexed\" ON \"as\" (\"indexed\");",
	// 		ExpectedBindings: []any{},
	// 	},
	// })
}
