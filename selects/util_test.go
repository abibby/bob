package selects_test

import (
	"encoding/json"
	"testing"

	"github.com/abibby/bob/insert"
	"github.com/abibby/bob/models"
	"github.com/abibby/bob/selects"
	"github.com/abibby/bob/test"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func NewTestBuilder() *selects.Builder[*test.Foo] {
	return selects.From[*test.Foo]()
}

func MustSave(tx *sqlx.Tx, v models.Model) {
	err := insert.Save(tx, v)
	if err != nil {
		panic(err)
	}
}

func assertJsonEqual(t *testing.T, rawJson string, v any) bool {
	b, err := json.Marshal(v)
	if !assert.NoError(t, err) {
		return false
	}
	var data any
	err = json.Unmarshal([]byte(rawJson), &data)
	if !assert.NoError(t, err) {
		return false
	}
	formattedJson, err := json.Marshal(data)
	if !assert.NoError(t, err) {
		return false
	}

	return assert.JSONEq(t, string(formattedJson), string(b))
}
