package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Orders struct {
	ent.Schema
}

func (Orders) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_time").Comment(""),
		field.Time("delete_time").Comment(""),
		field.Time("update_time").Comment(""),
	}
}
