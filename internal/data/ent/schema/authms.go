package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AuthMs holds the schema definition for the AuthMs entity.
type AuthMs struct {
	ent.Schema
}

// Fields of the AuthMs.
func (AuthMs) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("status").Default(0).Comment("状态 1:enable, 0:disable, -1:deleted"),
		field.Time("create_at").Comment("创建时间"),
		field.Time("update_at").Comment("更新时间"),
	}
}

// Edges of the AuthMs.
func (AuthMs) Edges() []ent.Edge {
	return nil
}
