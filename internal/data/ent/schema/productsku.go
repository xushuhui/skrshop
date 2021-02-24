package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ProductSku holds the schema definition for the ProductSku entity.
type ProductSku struct {
	ent.Schema
}

// Fields of the ProductSku.
func (ProductSku) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("spu_id").Default(0).Positive().Comment("账号 id"),
		field.Text("attrs").Comment("销售属性值 {attr_value_id}-{attr_value_id} 多个销售属性值逗号分隔"),
		field.Text("banner_url").Comment("banner 图片 多个图片逗号分隔"),
		field.Text("main_url").Comment("商品介绍主图 多个图片逗号分隔"),
		field.Int64("price_fee").Default(0).Positive().Comment("售价，整数方式保存"),
		field.Int8("price_scale").Default(0).Comment("售价，金额对应的小数位数"),
		field.Int64("market_price_fee").Default(0).Positive().Comment("市场价，整数方式保存"),
		field.Int8("market_price_scale").Default(0).Comment("市场价，金额对应的小数位数"),
		field.Time("create_at").Comment("创建时间"),
		field.Time("update_at").Comment("更新时间"),
		field.Int8("status").Default(0).Comment("状态 1:enable, 0:disable, -1:deleted"),
	}
}

// Edges of the ProductSku.
func (ProductSku) Edges() []ent.Edge {
	return nil
}
