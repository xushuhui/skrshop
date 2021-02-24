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
		field.Int8("status").Default(0).Comment("状态 1:enable, 0:disable, -1:deleted"),
		field.Time("create_at").Comment("创建时间"),
		field.Time("update_at").Comment("更新时间"),
	}
}

// Edges of the ProductBrands.
func (ProductBrands) Edges() []ent.Edge {
	return nil
}
