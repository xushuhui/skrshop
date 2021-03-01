package schema

import (
	"entgo.io/ent"
)

// ProductSpuSkuAttrMap holds the schema definition for the ProductSpuSkuAttrMap entity.
type ProductSpuSkuAttrMap struct {
	ent.Schema
}

// Fields of the ProductSpuSkuAttrMap.
func (ProductSpuSkuAttrMap) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the ProductSpuSkuAttrMap.
func (ProductSpuSkuAttrMap) Edges() []ent.Edge {
	return nil
}
func (ProductSpuSkuAttrMap) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
