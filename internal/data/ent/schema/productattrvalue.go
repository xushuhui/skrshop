package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ProductAttrValue holds the schema definition for the ProductAttrValue entity.
type ProductAttrValue struct {
	ent.Schema
}

// Fields of the ProductAttrValue.
func (ProductAttrValue) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("status").Default(0).Comment("状态 1:enable, 0:disable, -1:deleted"),
		field.Time("create_at").Comment("创建时间"),
		field.Time("update_at").Comment("更新时间"),
	}
}

// Edges of the ProductAttrValue.
func (ProductAttrValue) Edges() []ent.Edge {
	return nil
}
