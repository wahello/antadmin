package model

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
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
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Comment("创建时间"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间"),
		field.Time("deleted_at").
			Optional().
			Nillable().
			Comment("删除时间"),
		field.String("created_by").
			Immutable().
			Comment("创建者"),
		field.String("updated_by").
			Comment("修改者"),
	}
}
