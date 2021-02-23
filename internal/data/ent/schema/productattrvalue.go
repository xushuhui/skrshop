package schema

import "entgo.io/ent"

// ProductAttrValue holds the schema definition for the ProductAttrValue entity.
type ProductAttrValue struct {
	ent.Schema
}

// Fields of the ProductAttrValue.
func (ProductAttrValue) Fields() []ent.Field {
	return nil
}

// Edges of the ProductAttrValue.
func (ProductAttrValue) Edges() []ent.Edge {
	return nil
}
