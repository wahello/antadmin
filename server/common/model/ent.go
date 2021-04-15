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
			}),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			Optional().
			Nillable(),
		field.String("created_by").
			Immutable(),
		field.String("updated_by"),
	}
}
