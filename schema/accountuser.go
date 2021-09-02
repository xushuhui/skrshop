package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// AccountUser holds the schema definition for the AccountUser entity.
type AccountUser struct {
	ent.Schema
}

// Fields of the AccountUser.
func (AccountUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").MaxLen(30).Default("").Comment("邮箱"),
		field.String("phone").MaxLen(15).Default("").Comment("手机号"),
		field.String("password").MaxLen(32).Default("").Comment("密码"),
	}

}

// Edges of the AccountUser.
func (AccountUser) Edges() []ent.Edge {
	return nil
}
func (AccountUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email", "phone"),
	}
}
func (AccountUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
