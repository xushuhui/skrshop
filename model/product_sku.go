package model

type ProductSku struct {
	Id               uint   `gorm:"column:id;type:int unsigned" json:"id"`                                                 // SKU ID
	SpuId            uint   `gorm:"column:spu_id;type:int unsigned;default:'0'" json:"spu_id"`                             // SPU ID
	Attrs            string `gorm:"column:attrs;type:text;default:null" json:"attrs"`                                      // 销售属性值{attr_value_id}-{attr_value_id} 多个销售属性值逗号分隔
	BannerUrl        string `gorm:"column:banner_url;type:text;default:null" json:"banner_url"`                            // banner图片 多个图片逗号分隔
	MainUrl          string `gorm:"column:main_url;type:text;default:null" json:"main_url"`                                // 商品介绍主图 多个图片逗号分隔
	PriceFee         uint   `gorm:"column:price_fee;type:int unsigned;default:'0'" json:"price_fee"`                       // 售价，整数方式保存
	PriceScale       uint8  `gorm:"column:price_scale;type:tinyint unsigned;default:'0'" json:"price_scale"`               // 售价，金额对应的小数位数
	MarketPriceFee   uint   `gorm:"column:market_price_fee;type:int unsigned;default:'0'" json:"market_price_fee"`         // 市场价，整数方式保存
	MarketPriceScale uint8  `gorm:"column:market_price_scale;type:tinyint unsigned;default:'0'" json:"market_price_scale"` // 市场价，金额对应的小数位数
	CreateAt         uint   `gorm:"column:create_at;type:int unsigned;default:'0'" json:"create_at"`                       // 创建时间
	CreateBy         uint   `gorm:"column:create_by;type:int unsigned;default:'0'" json:"create_by"`                       // 创建人staff_id
	UpdateAt         uint   `gorm:"column:update_at;type:int unsigned;default:'0'" json:"update_at"`                       // 更新时间
	UpdateBy         uint   `gorm:"column:update_by;type:int unsigned;default:'0'" json:"update_by"`                       // 修改人staff_id
	Status           uint8  `gorm:"column:status;type:tinyint unsigned;default:'0'" json:"status"`                         // 状态 1:enable, 0:disable, -1:deleted
}

//get real primary key name
func (productSku *ProductSku) GetKey() string {
	return "id"
}

//get primary key in model
func (productSku *ProductSku) GetKeyProperty() uint {
	return productSku.Id
}

//set primary key
func (productSku *ProductSku) SetKeyProperty(id uint) {
	productSku.Id = id
}

//get real table name
func (productSku *ProductSku) TableName() string {
	return "product_sku"
}

func GetProductSkuById(id string) (productSku *ProductSku) {
	err := DB.Model(productSku).First(productSku, productSku.GetKey()+" = '"+id+"'").GetErrors()
	if len(err) > 0 {
		return nil
	}
	return
}

func GetProductSkuOne(where string, args ...interface{}) (productSku *ProductSku) {
	err := DB.Model(productSku).First(productSku, where, args).GetErrors()
	if len(err) > 0 {
		return nil
	}
	return
}

func GetProductSkuList(page, limit int64, where string, condition interface{}) (list []*ProductSku) {
	err := DB.Limit(limit).Offset((page-1)*limit).Find(list, where, condition).GetErrors()
	if err != nil {
		return nil
	}
	return
}

func (productSku *ProductSku) Create() []error {
	return DB.Model(productSku).Create(productSku).GetErrors()
}

func (productSku *ProductSku) Update(update ProductSku) []error {
	return DB.Model(productSku).UpdateColumns(update).GetErrors()
}

func (productSku *ProductSku) Delete() {
	DB.Model(productSku).Delete(productSku)
}
