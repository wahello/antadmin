package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/antbiz/antadmin/db/mixins"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Mixin of the Role
func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Unique().
			Comment("角色名称"),
		field.Bool("disabled").
			Default(false).
			Comment("是否禁用"),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
