// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_name", Type: field.TypeString, Unique: true},
		{Name: "nick_name", Type: field.TypeString, Default: "unknown"},
		{Name: "password", Type: field.TypeString},
		{Name: "user_type", Type: field.TypeInt, Default: 1},
		{Name: "status", Type: field.TypeInt, Default: 1},
		{Name: "delete_at", Type: field.TypeInt64, Nullable: true},
		{Name: "create_at", Type: field.TypeInt64, Default: 1709103601},
		{Name: "update_at", Type: field.TypeInt64, Default: 1709103601},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
	}
)

func init() {
}
