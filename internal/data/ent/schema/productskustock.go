package schema

import "entgo.io/ent"

// ProductSkuStock holds the schema definition for the ProductSkuStock entity.
type ProductSkuStock struct {
	ent.Schema
}

// Fields of the ProductSkuStock.
func (ProductSkuStock) Fields() []ent.Field {
	return nil
}

// Edges of the ProductSkuStock.
func (ProductSkuStock) Edges() []ent.Edge {
	return nil
}
