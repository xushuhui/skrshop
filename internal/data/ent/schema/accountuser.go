package schema

import "entgo.io/ent"

// AccountUser holds the schema definition for the AccountUser entity.
type AccountUser struct {
	ent.Schema
}

// Fields of the AccountUser.
func (AccountUser) Fields() []ent.Field {
	return nil
}

// Edges of the AccountUser.
func (AccountUser) Edges() []ent.Edge {
	return nil
}
