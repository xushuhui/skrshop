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
		field.Int64("ms_id").Default(0).Positive().Comment("系统 id"),
	}
}

// Edges of the AuthItem.
func (AuthItem) Edges() []ent.Edge {
	return nil
}
func (AuthItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
