package schema

import "entgo.io/ent"

// SkrMember holds the schema definition for the SkrMember entity.
type SkrMember struct {
	ent.Schema
}

// Fields of the SkrMember.
func (SkrMember) Fields() []ent.Field {
	return nil
}

// Edges of the SkrMember.
func (SkrMember) Edges() []ent.Edge {
	return nil
}
