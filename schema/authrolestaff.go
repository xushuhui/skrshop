package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AuthRoleStaff holds the schema definition for the AuthRoleStaff entity.
type AuthRoleStaff struct {
	ent.Schema
}

// Fields of the AuthRoleStaff.
func (AuthRoleStaff) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("staff_id").Default(0).Positive().Comment("员工 id"),
	}
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
