package migrate

import (
	"fmt"

	"github.com/abibby/bob/builder"
	"github.com/abibby/bob/models"
)

type dropTable string

func drop(table models.Model) dropTable {
	return dropTable(builder.GetTable(table))
}

func (dt dropTable) ToGo() string {
	return fmt.Sprintf("schema.DropIfExists(%#v)", dt)
}
