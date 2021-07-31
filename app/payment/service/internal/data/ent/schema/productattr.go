package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ProductAttr holds the schema definition for the ProductAttr entity.
type ProductAttr struct {
	ent.Schema
}

// Fields of the ProductAttr.
func (ProductAttr) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(230).Default("").Comment("名称"),
		field.String("desc").MaxLen(230).Default("").Comment("描述"),
	}
}

// Edges of the ProductAttr.
func (ProductAttr) Edges() []ent.Edge {
	return nil
}
func (ProductAttr) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
