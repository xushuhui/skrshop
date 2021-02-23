package schema

import "entgo.io/ent"

// ProductCategory holds the schema definition for the ProductCategory entity.
type ProductCategory struct {
	ent.Schema
}

// Fields of the ProductCategory.
func (ProductCategory) Fields() []ent.Field {
	return nil
}

// Edges of the ProductCategory.
func (ProductCategory) Edges() []ent.Edge {
	return nil
}
