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
			Unique().
			Comment("用户名").
			StructTag(`json:"username" v:"required|length:2,16#用户名不能为空|用户名长度应当在:min到:max之前"`),
		field.String("password").
			Sensitive().
			Comment("密码"),
		field.String("phone").
			MaxLen(11).
			Nillable().
			Optional().
			Unique().
			Comment("手机号").
			StructTag(`json:"phone,omitempty" v:"phone#手机号格式不正确"`),
		field.String("email").
			Nillable().
			Optional().
			Unique().
			Comment("邮箱").
			StructTag(`json:"email,omitempty" v:"email#邮箱格式不正确"`),
		field.String("avatar").
			Comment("头像"),
		field.Int("gender").
			Min(0).
			Max(2).
			Default(0).
			Comment("性别").
			StructTag(`json:"gender" v:"between:0,2#性别参数大小为0到2"`),
		field.Int("status").
			Min(0).
			Max(2).
			Default(0).
			Comment("用户状态").
			StructTag(`json:"status" v:"between:0,2#用户状态参数大小为0到2"`),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
