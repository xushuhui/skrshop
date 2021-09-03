package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type ActivityCoupon struct {
	ent.Schema
}

func (ActivityCoupon) Fields() []ent.Field {
	return []ent.Field{
		field.Int("coupon_id").Comment(""),
		field.Int("activity_id").Comment(""),
	}
}
