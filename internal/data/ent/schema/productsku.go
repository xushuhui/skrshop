package schema

import "entgo.io/ent"

// ProductSku holds the schema definition for the ProductSku entity.
type ProductSku struct {
	ent.Schema
}

// Fields of the ProductSku.
func (ProductSku) Fields() []ent.Field {
	return nil
}

// Edges of the ProductSku.
func (ProductSku) Edges() []ent.Edge {
	return nil
}
