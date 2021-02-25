package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AuthItem holds the schema definition for the AuthItem entity.
type AuthItem struct {
	ent.Schema
}

// Fields of the AuthItem.
func (AuthItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("status").Default(0).Comment("状态 1:enable, 0:disable, -1:deleted"),
		field.Time("create_at").Comment("创建时间"),
		field.Time("update_at").Comment("更新时间"),
	}
}

// Edges of the AuthItem.
func (AuthItem) Edges() []ent.Edge {
	return nil
}
