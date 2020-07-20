package model

type ProductSpu struct {
	Id               uint   `gorm:"column:id;type:int unsigned" json:"id"`                                                 // SPU ID
	BrandId          uint   `gorm:"column:brand_id;type:int unsigned;default:'0'" json:"brand_id"`                         // 品牌ID
	CategoryId       uint   `gorm:"column:category_id;type:int unsigned;default:'0'" json:"category_id"`                   // 分类ID
	Name             string `gorm:"column:name;type:varchar(255);default:''" json:"name"`                                  // spu名称
	Desc             string `gorm:"column:desc;type:varchar(255);default:''" json:"desc"`                                  // spu描述
	SellingPoint     string `gorm:"column:selling_point;type:varchar(255);default:''" json:"selling_point"`                // 卖点
	Unit             string `gorm:"column:unit;type:varchar(255);default:''" json:"unit"`                                  // spu单位
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
func (productSpu *ProductSpu) GetKey() string {
	return "id"
}

//get primary key in model
func (productSpu *ProductSpu) GetKeyProperty() uint {
	return productSpu.Id
}

//set primary key
func (productSpu *ProductSpu) SetKeyProperty(id uint) {
	productSpu.Id = id
}

//get real table name
func (productSpu *ProductSpu) TableName() string {
	return "product_spu"
}

func GetProductSpuById(id string) (productSpu *ProductSpu, err error) {
	err = DB.Model(productSpu).First(productSpu, productSpu.GetKey()+" = '"+id+"'").Error

	return
}

func GetProductSpuOne(where string, args ...interface{}) (productSpu *ProductSpu, err error) {
	err = DB.Model(productSpu).First(productSpu, where, args).Error

	return
}

func GetProductSpuList(page, limit int64, where string, condition interface{}) (list []*ProductSpu) {
	err := DB.Limit(limit).Offset((page-1)*limit).Find(list, where, condition).GetErrors()
	if err != nil {
		return nil
	}
	return
}

func (productSpu *ProductSpu) Create() []error {
	return DB.Model(productSpu).Create(productSpu).GetErrors()
}

func (productSpu *ProductSpu) Update(update ProductSpu) []error {
	return DB.Model(productSpu).UpdateColumns(update).GetErrors()
}

func (productSpu *ProductSpu) Delete() {
	DB.Model(productSpu).Delete(productSpu)
}
