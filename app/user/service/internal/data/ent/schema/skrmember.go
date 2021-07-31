package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// SkrMember holds the schema definition for the SkrMember entity.
type SkrMember struct {
	ent.Schema
}

// Fields of the SkrMember.
func (SkrMember) Fields() []ent.Field {
	return []ent.Field{
		//TODO: uid是否去掉
		field.Int64("uid").Default(0).Positive().Comment("账号 id"),

		field.String("nickname").MaxLen(30).Default("").Comment(""),
		field.String("avatar").MaxLen(255).Default("").Comment(""),
		field.Int8("gender").Default(1).Comment("性别 1 unknow 2 male 3 female"),
		field.Int8("role").Default(0).Comment("角色 0: 普通用户 1:vip"),
		field.Time("create_at").Comment("创建时间"),
		field.Time("update_at").Comment("更新时间"),
	}
}

// Edges of the SkrMember.
func (SkrMember) Edges() []ent.Edge {
	return nil
}
func (SkrMember) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uid"),
	}
}
func (SkrMember) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
