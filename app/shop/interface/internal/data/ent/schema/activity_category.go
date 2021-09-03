package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type ActivityCategory struct {
	ent.Schema
}

func (ActivityCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Int("category_id").Comment(""),
		field.Int("activity_id").Comment(""),
	}
}
