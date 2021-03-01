package schema

import (
	"entgo.io/ent"
)

// AuthRoleStaff holds the schema definition for the AuthRoleStaff entity.
type AuthRoleStaff struct {
	ent.Schema
}

// Fields of the AuthRoleStaff.
func (AuthRoleStaff) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the AuthRoleStaff.
func (AuthRoleStaff) Edges() []ent.Edge {
	return nil
}
func (AuthRoleStaff) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
