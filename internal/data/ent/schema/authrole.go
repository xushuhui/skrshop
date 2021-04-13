package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AuthRole holds the schema definition for the AuthRole entity.
type AuthRole struct {
	ent.Schema
}

// Fields of the AuthRole.
func (AuthRole) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(230).Default("").Comment("名称"),
		field.String("desc").MaxLen(230).Default("").Comment("描述"),
	}
}

// Edges of the AuthRole.
func (AuthRole) Edges() []ent.Edge {
	return nil
}
func (AuthRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
