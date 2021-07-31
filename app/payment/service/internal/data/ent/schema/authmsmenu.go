package schema

import (
	"entgo.io/ent"
)

// AuthMsMenu holds the schema definition for the AuthMsMenu entity.
type AuthMsMenu struct {
	ent.Schema
}

// Fields of the AuthMsMenu.
func (AuthMsMenu) Fields() []ent.Field {
	return []ent.Field{
		
	}
}

// Edges of the AuthMsMenu.
func (AuthMsMenu) Edges() []ent.Edge {
	return nil
}
func (AuthMsMenu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
