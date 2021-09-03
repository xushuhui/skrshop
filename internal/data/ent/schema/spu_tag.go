package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type SpuTag struct {
	ent.Schema
}

func (SpuTag) Fields() []ent.Field {
	return []ent.Field{
		field.Int("spu_id").Comment(""),
		field.Int("tag_id").Comment(""),
	}
}
