package data

import (
	"context"

	"github.com/antbiz/antadmin/module/system/data/ent"
	"github.com/gogf/gf/frame/g"
)

var client *ent.Client

func init() {
	cli, err := ent.Open(
		g.Cfg().GetString("database.default.name"),
		g.Cfg().GetString("database.default.link"),
	)
	if err != nil {
		g.Log().Fatalf("[Module-System] - failed opeing connection to database: %v", err)
	}

	client = cli
	if err := client.Schema.Create(context.Background()); err != nil {
		g.Log().Fatalf("[Module-System] - failed creating schema resources: %v", err)
	}
}
