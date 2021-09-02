package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// StaffInfo holds the schema definition for the StaffInfo entity.
type StaffInfo struct {
	ent.Schema
}

// Fields of the StaffInfo.
func (StaffInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("uid").Default(0).Positive().Comment("账号 id"),
		field.String("email").MaxLen(30).Default("").Comment("邮箱"),
		field.String("phone").MaxLen(15).Default("").Comment("手机号"),
		field.String("password").MaxLen(32).Default("").Comment("密码"),
		field.String("name").MaxLen(30).Default("").Comment("姓名"),
		field.String("avatar").MaxLen(255).Default("").Comment(""),
		field.Int8("gender").Default(1).Comment("性别 1 unknow 2 male 3 female"),
	}
}

// Edges of the StaffInfo.
func (StaffInfo) Edges() []ent.Edge {
	return nil
}
func (StaffInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email", "phone"),
	}
}
func (StaffInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Mixin() []ent.Field {
	return []ent.Field{
		field.Time("create_at").Comment("创建时间"),
		field.Time("update_at").Comment("更新时间"),
		field.Int8("status").Default(0).Comment("状态 1:enable, 0:disable, -1:deleted"),
	}
}
