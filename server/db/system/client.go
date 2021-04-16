package system

import (
	"entgo.io/ent/dialect/sql"
	"github.com/antbiz/antadmin/db/system/ent"
)

// Client is alias
type Client *ent.Client

// NewClient .
func NewClient(drv *sql.Driver) Client {
	return ent.NewClient(ent.Driver(drv))
}
