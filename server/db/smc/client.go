package smc

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/antbiz/antadmin/db/helper"
	"github.com/antbiz/antadmin/db/smc/ent"
	_ "github.com/antbiz/antadmin/db/smc/ent/runtime"
	"github.com/gogf/gf/frame/g"
)

// Client is alias
type Client *ent.Client

func newClient(drv *sql.Driver) Client {
	client := ent.NewClient(ent.Driver(drv)).Debug()

	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			v, err := next.Mutate(ctx, m)
			if err != nil {
				g.DB().GetLogger().Error(err)
				return nil, err
			}
			return v, nil
		})
	})

	if err := client.Schema.Create(context.Background()); err != nil {
		g.Log().Fatalf("[Module SMC] - failed creating schema resources: %v", err)
	}

	return client
}

// InitClient 初始化主从 ent client
func InitClient() (master Client, slave Client, err error) {
	masterDrv, er := helper.GetMasterDriver()
	if er != nil {
		g.Log().Errorf("init smc-module master db failed: %v", er)
		return nil, nil, er
	}
	master = newClient(masterDrv)

	slaveDrv, er := helper.GetSlaveDriver()
	if er != nil {
		g.Log().Errorf("init smc-module slave db failed: %v", er)
		slave = master
	} else {
		slave = newClient(slaveDrv)
	}

	return
}
