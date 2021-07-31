package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ProductSpuSkuAttrMap holds the schema definition for the ProductSpuSkuAttrMap entity.
type ProductSpuSkuAttrMap struct {
	ent.Schema
}

// Fields of the ProductSpuSkuAttrMap.
func (ProductSpuSkuAttrMap) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("sku_id").Default(0).Positive().Comment("sku_id"),
	}
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
