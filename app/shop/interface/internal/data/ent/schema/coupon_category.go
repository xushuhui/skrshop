package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type CouponCategory struct {
	ent.Schema
}

func (CouponCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Int("category_id").Comment(""),
		field.Int("coupon_id").Comment(""),
	}
}
