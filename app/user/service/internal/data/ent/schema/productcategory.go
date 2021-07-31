package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ProductCategory holds the schema definition for the ProductCategory entity.
type ProductCategory struct {
	ent.Schema
}

// Fields of the ProductCategory.
func (ProductCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("pid").Default(0).Positive().Comment("父 id"),
		field.String("name").MaxLen(255).Default("").Comment("名称"),
		field.String("desc").MaxLen(255).Default("").Comment("描述"),

		field.String("pic_url").MaxLen(255).Default("").Comment(""),
		field.String("path").MaxLen(255).Default("").Comment("分类地址 {pid}-{child_id}-..."),
	}
}

// Edges of the ProductCategory.
func (ProductCategory) Edges() []ent.Edge {
	return nil
}
func (ProductCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
