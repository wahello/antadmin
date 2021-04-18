package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/antbiz/antadmin/db/mixins"
)

// Menu holds the schema definition for the Menu entity.
type Menu struct {
	ent.Schema
}

// Mixin of the Menu
func (Menu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("菜单名称"),
		field.Bool("hide").
			Default(false).
			Comment("是否隐藏"),
		field.String("path").
			NotEmpty().
			Comment("菜单路由"),
		field.Int("sort").
			Default(0).
			Comment("展示顺序"),
		field.String("icon").
			Comment("图标"),
		field.String("parent").
			Comment("上级菜单"),
	}
}

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return nil
}
