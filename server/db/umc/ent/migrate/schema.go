// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "disabled", Type: field.TypeBool, Default: false},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:        "roles",
		Columns:     RolesColumns,
		PrimaryKey:  []*schema.Column{RolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "role_created_at",
				Unique:  false,
				Columns: []*schema.Column{RolesColumns[1]},
			},
			{
				Name:    "role_created_by",
				Unique:  false,
				Columns: []*schema.Column{RolesColumns[4]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString, Unique: true, Nullable: true, Size: 11},
		{Name: "email", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "gender", Type: field.TypeInt, Default: 0},
		{Name: "disabled", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "user_created_at",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[1]},
			},
			{
				Name:    "user_created_by",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[4]},
			},
		},
	}
	// RoleUsersColumns holds the columns for the "role_users" table.
	RoleUsersColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
	}
	// RoleUsersTable holds the schema information for the "role_users" table.
	RoleUsersTable = &schema.Table{
		Name:       "role_users",
		Columns:    RoleUsersColumns,
		PrimaryKey: []*schema.Column{RoleUsersColumns[0], RoleUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_users_role_id",
				Columns:    []*schema.Column{RoleUsersColumns[0]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_users_user_id",
				Columns:    []*schema.Column{RoleUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		RolesTable,
		UsersTable,
		RoleUsersTable,
	}
)

func init() {
	RoleUsersTable.ForeignKeys[0].RefTable = RolesTable
	RoleUsersTable.ForeignKeys[1].RefTable = UsersTable
}
