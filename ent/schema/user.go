package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_name").Unique(),
		field.String("nick_name").Default("unknown"),
		field.String("password"),
		field.Int("user_type").Default(1),
		field.Int("status").Default(1),
		field.Int64("delete_at").Optional(),
		field.Int64("create_at").Default(time.Now().Unix()).Immutable(),
		field.Int64("update_at").Default(time.Now().Unix()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
