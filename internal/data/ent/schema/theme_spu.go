package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type ThemeSpu struct {
	ent.Schema
}

func (ThemeSpu) Fields() []ent.Field {
	return []ent.Field{
		field.Int("theme_id").Comment(""),
		field.Int("spu_id").Comment(""),
	}
}
