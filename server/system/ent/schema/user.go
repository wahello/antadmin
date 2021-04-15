package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/antbiz/antadmin/common/model"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		model.CommonMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Immutable().
			NotEmpty().
			Unique(),
		field.String("password").
			Sensitive(),
		field.String("phone").
			MaxLen(11).
			Nillable().
			Optional().
			Unique(),
		field.String("email").
			Nillable().
			Optional().
			Unique(),
		field.String("avatar"),
		field.Int("gender").
			Min(0).
			Max(2).
			Default(0),
		field.Int("status").
			Min(0).
			Max(2).
			Default(0),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
