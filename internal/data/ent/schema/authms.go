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
	return []ent.Field{}
}

// Edges of the AuthMs.
func (AuthMs) Edges() []ent.Edge {
	return nil
}
func (AuthMs) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
