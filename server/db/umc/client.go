package umc

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/antbiz/antadmin/db/umc/ent"
	_ "github.com/antbiz/antadmin/db/umc/ent/runtime"
	"github.com/gogf/gf/frame/g"
)

// Client is alias
type Client *ent.Client

// NewClient .
func NewClient(drv *sql.Driver) Client {
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
		g.Log().Fatalf("[Module UPMS] - failed creating schema resources: %v", err)
	}

	return client
}
