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
		field.Time("create_at").Comment("创建时间"),
		field.Time("update_at").Comment("更新时间"),
		field.Int8("status").Default(0).Comment("状态 1:enable, 0:disable, -1:deleted"),
	}
}

// Edges of the ProductSpuSkuAttrMap.
func (ProductSpuSkuAttrMap) Edges() []ent.Edge {
	return nil
}
