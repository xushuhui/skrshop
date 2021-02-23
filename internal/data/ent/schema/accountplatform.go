package schema

import "entgo.io/ent"

// AccountPlatform holds the schema definition for the AccountPlatform entity.
type AccountPlatform struct {
	ent.Schema
}

// Fields of the AccountPlatform.
func (AccountPlatform) Fields() []ent.Field {
	return nil
}

// Edges of the AccountPlatform.
func (AccountPlatform) Edges() []ent.Edge {
	return nil
}
