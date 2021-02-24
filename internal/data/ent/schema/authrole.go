package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AuthRole holds the schema definition for the AuthRole entity.
type AuthRole struct {
	ent.Schema
}

// Fields of the AuthRole.
func (AuthRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("status").Default(0).Comment("状态 1:enable, 0:disable, -1:deleted"),
		field.Time("create_at").Comment("创建时间"),
		field.Time("update_at").Comment("更新时间"),
	}
}

// Edges of the AuthRole.
func (AuthRole) Edges() []ent.Edge {
	return nil
}
