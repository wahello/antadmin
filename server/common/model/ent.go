package model

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/guid"
)

// CommonMixin .
type CommonMixin struct {
	mixin.Schema
}

// Fields .
func (CommonMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return guid.S()
			}).
			Comment("编号"),
		field.Time("createdAt").
			Immutable().
			Default(time.Now).
			Comment("创建时间"),
		field.Time("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间"),
		field.Time("deletedAt").
			Optional().
			Nillable().
			Comment("删除时间"),
		field.String("createdBy").
			Optional().
			Comment("创建者"),
		field.String("updatedBy").
			Optional().
			Comment("修改者"),
	}
}

// Indexes .
func (CommonMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("createdAt"),
		index.Fields("createdBy"),
	}
}

// Hooks .
func (CommonMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				v, err := next.Mutate(ctx, m)
				if err != nil {
					g.DB().GetLogger().Error(err)
					return nil, err
				}
				return v, nil
			})
		},
	}
}
