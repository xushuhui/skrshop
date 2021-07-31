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
		field.String("value").MaxLen(230).Default("").Comment("属性值"),
		field.String("desc").MaxLen(230).Default("").Comment("描述"),
	}
}

// Edges of the ProductAttrValue.
func (ProductAttrValue) Edges() []ent.Edge {
	return nil
}
func (ProductAttrValue) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
