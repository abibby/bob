package migrate_test

import (
	"testing"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/migrate"
	"github.com/abibby/bob/models"
	"github.com/abibby/bob/schema"
	"github.com/abibby/nulls"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/assert"
)

func TestGenerateMigration(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		type TestModel struct {
			models.BaseModel
			ID       int           `db:"id,primary"`
			Nullable *nulls.String `db:"nullable"`
			Indexed  bool          `db:"indexed,index"`
		}
		src, err := migrate.New().GenerateMigration("2023-01-01T00:00:00Z create test model", "packageName", &TestModel{})
		assert.NoError(t, err)
		cupaloy.SnapshotT(t, src)
	})

	t.Run("add column", func(t *testing.T) {
		m := migrate.New()
		m.Add(&migrate.Migration{
			Name: "2023-01-01T00:00:00Z create test model",
			Up: func() builder.ToSQLer {
				return schema.Create("test_models", func(table *schema.Blueprint) {
					table.Int("id").Primary()
				})
			},
			Down: func() builder.ToSQLer {
				return schema.DropIfExists("test_models")
			},
		})

		type TestModel struct {
			models.BaseModel
			ID    int    `db:"id,primary"`
			ToAdd string `db:"to_add"`
		}
		src, err := m.GenerateMigration("2023-01-01T00:00:00Z create test model", "packageName", &TestModel{})
		assert.NoError(t, err)
		cupaloy.SnapshotT(t, src)
	})

	t.Run("drop column", func(t *testing.T) {
		m := migrate.New()
		m.Add(&migrate.Migration{
			Name: "2023-01-01T00:00:00Z create test model",
			Up: func() builder.ToSQLer {
				return schema.Create("test_models", func(table *schema.Blueprint) {
					table.Int("id").Primary()
					table.String("to_drop")
				})
			},
			Down: func() builder.ToSQLer {
				return schema.DropIfExists("test_models")
			},
		})

		type TestModel struct {
			models.BaseModel
			ID int `db:"id,primary"`
		}
		src, err := m.GenerateMigration("2023-01-01T00:00:00Z create test model", "packageName", &TestModel{})
		assert.NoError(t, err)
		cupaloy.SnapshotT(t, src)
	})

	t.Run("change", func(t *testing.T) {
		m := migrate.New()
		m.Add(&migrate.Migration{
			Name: "2023-01-01T00:00:00Z create test model",
			Up: func() builder.ToSQLer {
				return schema.Create("test_models", func(table *schema.Blueprint) {
					table.Int("id").Primary()
				})
			},
			Down: func() builder.ToSQLer {
				return schema.DropIfExists("test_models")
			},
		})

		type TestModel struct {
			models.BaseModel
			ID string `db:"id,primary"`
		}
		src, err := m.GenerateMigration("2023-01-01T00:00:00Z create test model", "packageName", &TestModel{})
		assert.NoError(t, err)
		cupaloy.SnapshotT(t, src)
	})

	t.Run("no changes", func(t *testing.T) {
		m := migrate.New()
		m.Add(&migrate.Migration{
			Name: "2023-01-01T00:00:00Z create test model",
			Up: func() builder.ToSQLer {
				return schema.Create("test_models", func(table *schema.Blueprint) {
					table.Int("id").Primary()
				})
			},
			Down: func() builder.ToSQLer {
				return schema.DropIfExists("test_models")
			},
		})

		type TestModel struct {
			models.BaseModel
			ID int `db:"id,primary"`
		}
		_, err := m.GenerateMigration("2023-01-01T00:00:00Z create test model", "packageName", &TestModel{})
		assert.ErrorIs(t, err, migrate.ErrNoChanges)
	})

	t.Run("multiple migrations", func(t *testing.T) {
		m := migrate.New()
		m.Add(&migrate.Migration{
			Name: "2023-01-01T00:00:00Z create test model",
			Up: func() builder.ToSQLer {
				return schema.Create("test_models", func(table *schema.Blueprint) {
					table.Int("id").Primary()
				})
			},
			Down: nil,
		})
		m.Add(&migrate.Migration{
			Name: "2023-01-01T00:00:01Z change",
			Up: func() builder.ToSQLer {
				return schema.Table("test_models", func(table *schema.Blueprint) {
					table.String("id").Primary().Change()
				})
			},
			Down: nil,
		})

		type TestModel struct {
			models.BaseModel
			ID int `db:"id,primary"`
		}
		src, err := m.GenerateMigration("2023-01-01T00:00:00Z create test model", "packageName", &TestModel{})
		assert.NoError(t, err)
		cupaloy.SnapshotT(t, src)
	})
}
