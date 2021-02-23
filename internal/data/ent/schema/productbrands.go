package schema

import "entgo.io/ent"

// ProductBrands holds the schema definition for the ProductBrands entity.
type ProductBrands struct {
	ent.Schema
}

// Fields of the ProductBrands.
func (ProductBrands) Fields() []ent.Field {
	return nil
}

// Edges of the ProductBrands.
func (ProductBrands) Edges() []ent.Edge {
	return nil
}
