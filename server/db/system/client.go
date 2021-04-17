package system

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/antbiz/antadmin/db/system/ent"
	_ "github.com/antbiz/antadmin/db/system/ent/runtime"
	"github.com/gogf/gf/frame/g"
)

// Client is alias
type Client *ent.Client

// NewClient .
func NewClient(drv *sql.Driver) Client {
	client := ent.NewClient(ent.Driver(drv)).Debug()
	if err := client.Schema.Create(context.Background()); err != nil {
		g.Log().Fatalf("[Module System] - failed creating schema resources: %v", err)
	}

	return client
}
