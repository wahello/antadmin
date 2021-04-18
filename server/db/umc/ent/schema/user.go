package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/antbiz/antadmin/db/mixins"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Immutable().
			NotEmpty().
			Unique().
			Comment("用户名"),
		field.String("password").
			Sensitive().
			Comment("密码"),
		field.String("phone").
			MaxLen(11).
			Nillable().
			Optional().
			Unique().
			Comment("手机号"),
		field.String("email").
			Nillable().
			Optional().
			Unique().
			Comment("邮箱"),
		field.String("avatar").
			Optional().
			Comment("头像"),
		field.Int("gender").
			Min(0).
			Max(2).
			Default(0).
			Comment("性别"),
		field.Bool("disabled").
			Default(false).
			Comment("是否禁用"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).
			Ref("users"),
	}
}
