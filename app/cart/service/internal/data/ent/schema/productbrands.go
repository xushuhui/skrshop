package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ProductBrands holds the schema definition for the ProductBrands entity.
type ProductBrands struct {
	ent.Schema
}

// Fields of the ProductBrands.
func (ProductBrands) Fields() []ent.Field {
	return []ent.Field{

		field.String("name").MaxLen(255).Default("").Comment("名称"),
		field.String("desc").MaxLen(255).Default("").Comment("描述"),

		field.String("logo_url").MaxLen(255).Default("").Comment(""),
	}
}

// Edges of the ProductBrands.
func (ProductBrands) Edges() []ent.Edge {
	return nil
}
func (ProductBrands) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
