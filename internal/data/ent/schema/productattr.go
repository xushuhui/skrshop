package schema

import (
	"entgo.io/ent"
)

// ProductAttr holds the schema definition for the ProductAttr entity.
type ProductAttr struct {
	ent.Schema
}

// Fields of the ProductAttr.
func (ProductAttr) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the ProductAttr.
func (ProductAttr) Edges() []ent.Edge {
	return nil
}
func (ProductAttr) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
