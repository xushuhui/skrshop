package model

type ProductSkuStock struct {
	SkuId    uint  `gorm:"column:sku_id;type:int unsigned" json:"sku_id"`                   // SKU ID
	Quantity uint  `gorm:"column:quantity;type:int unsigned;default:'0'" json:"quantity"`   // 库存
	CreateAt uint  `gorm:"column:create_at;type:int unsigned;default:'0'" json:"create_at"` // 创建时间
	CreateBy uint  `gorm:"column:create_by;type:int unsigned;default:'0'" json:"create_by"` // 创建人staff_id
	UpdateAt uint  `gorm:"column:update_at;type:int unsigned;default:'0'" json:"update_at"` // 更新时间
	UpdateBy uint  `gorm:"column:update_by;type:int unsigned;default:'0'" json:"update_by"` // 修改人staff_id
	Status   uint8 `gorm:"column:status;type:tinyint unsigned;default:'0'" json:"status"`   // 状态 1:enable, 0:disable, -1:deleted
}

//get real primary key name
func (productSkuStock *ProductSkuStock) GetKey() string {
	return "sku_id"
}

//get primary key in model
func (productSkuStock *ProductSkuStock) GetKeyProperty() uint {
	return productSkuStock.SkuId
}

//set primary key
func (productSkuStock *ProductSkuStock) SetKeyProperty(id uint) {
	productSkuStock.SkuId = id
}

//get real table name
func (productSkuStock *ProductSkuStock) TableName() string {
	return "product_sku_stock"
}

func GetProductSkuStockById(id string) (productSkuStock *ProductSkuStock) {
	err := DB.Model(productSkuStock).First(productSkuStock, productSkuStock.GetKey()+" = '"+id+"'").GetErrors()
	if len(err) > 0 {
		return nil
	}
	return
}

func GetProductSkuStockOne(where string, args ...interface{}) (productSkuStock *ProductSkuStock) {
	err := DB.Model(productSkuStock).First(productSkuStock, where, args).GetErrors()
	if len(err) > 0 {
		return nil
	}
	return
}

func GetProductSkuStockList(page, limit int64, where string, condition interface{}) (list []*ProductSkuStock) {
	err := DB.Limit(limit).Offset((page-1)*limit).Find(list, where, condition).GetErrors()
	if err != nil {
		return nil
	}
	return
}

func (productSkuStock *ProductSkuStock) Create() []error {
	return DB.Model(productSkuStock).Create(productSkuStock).GetErrors()
}

func (productSkuStock *ProductSkuStock) Update(update ProductSkuStock) []error {
	return DB.Model(productSkuStock).UpdateColumns(update).GetErrors()
}

func (productSkuStock *ProductSkuStock) Delete() {
	DB.Model(productSkuStock).Delete(productSkuStock)
}
