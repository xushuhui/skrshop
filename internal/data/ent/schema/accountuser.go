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
		field.Time("create_at").Comment("创建时间"),
		field.Int8("status").Default(0).Comment("状态 1:enable, 0:disable, -1:deleted"),
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
