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
			MaxLen(11),
		field.String("email"),
		field.String("avatar"),
		field.Int("gender").
			Default(0),
		field.String("remark"),
		field.Int("status").
			Default(0),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
