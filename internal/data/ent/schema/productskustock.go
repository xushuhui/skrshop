package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ProductSkuStock holds the schema definition for the ProductSkuStock entity.
type ProductSkuStock struct {
	ent.Schema
}

// Fields of the ProductSkuStock.
func (ProductSkuStock) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the ProductSkuStock.
func (ProductSkuStock) Edges() []ent.Edge {
	return nil
}
func (ProductSkuStock) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
