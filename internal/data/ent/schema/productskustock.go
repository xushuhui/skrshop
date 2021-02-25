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
	return []ent.Field{
		field.Time("create_at").Comment("创建时间"),
		field.Time("update_at").Comment("更新时间"),
		field.Int8("status").Default(0).Comment("状态 1:enable, 0:disable, -1:deleted"),
	}
}

// Edges of the ProductSkuStock.
func (ProductSkuStock) Edges() []ent.Edge {
	return nil
}
