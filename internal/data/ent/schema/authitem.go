package schema

import (
	"entgo.io/ent"
)

// AuthItem holds the schema definition for the AuthItem entity.
type AuthItem struct {
	ent.Schema
}

// Fields of the AuthItem.
func (AuthItem) Fields() []ent.Field {
	return []ent.Field{}
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
