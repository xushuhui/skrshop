package model

type ProductSpuSkuAttrMap struct {
	Id            uint   `gorm:"column:id;type:int unsigned" json:"id"`                                       // 自增ID
	SpuId         uint   `gorm:"column:spu_id;type:int unsigned;default:'0'" json:"spu_id"`                   // SPU ID
	SkuId         uint   `gorm:"column:sku_id;type:int unsigned;default:'0'" json:"sku_id"`                   // SKU ID
	AttrId        uint   `gorm:"column:attr_id;type:int unsigned;default:'0'" json:"attr_id"`                 // 销售属性ID
	AttrName      string `gorm:"column:attr_name;type:varchar(255);default:'0'" json:"attr_name"`             // 销售属性名称
	AttrValueId   uint   `gorm:"column:attr_value_id;type:int unsigned;default:'0'" json:"attr_value_id"`     // 销售属性值ID
	AttrValueName string `gorm:"column:attr_value_name;type:varchar(255);default:'0'" json:"attr_value_name"` // 销售属性值
	CreateAt      uint   `gorm:"column:create_at;type:int unsigned;default:'0'" json:"create_at"`             // 创建时间
	CreateBy      uint   `gorm:"column:create_by;type:int unsigned;default:'0'" json:"create_by"`             // 创建人staff_id
	Status        uint8  `gorm:"column:status;type:tinyint unsigned;default:'0'" json:"status"`               // 状态 1:enable, 0:disable, -1:deleted
}

//get real primary key name
func (productSpuSkuAttrMap *ProductSpuSkuAttrMap) GetKey() string {
	return "id"
}

//get primary key in model
func (productSpuSkuAttrMap *ProductSpuSkuAttrMap) GetKeyProperty() uint {
	return productSpuSkuAttrMap.Id
}

//set primary key
func (productSpuSkuAttrMap *ProductSpuSkuAttrMap) SetKeyProperty(id uint) {
	productSpuSkuAttrMap.Id = id
}

//get real table name
func (productSpuSkuAttrMap *ProductSpuSkuAttrMap) TableName() string {
	return "product_spu_sku_attr_map"
}

func GetProductSpuSkuAttrMapById(id string) (productSpuSkuAttrMap *ProductSpuSkuAttrMap) {
	err := DB.Model(productSpuSkuAttrMap).First(productSpuSkuAttrMap, productSpuSkuAttrMap.GetKey()+" = '"+id+"'").GetErrors()
	if len(err) > 0 {
		return nil
	}
	return
}

func GetProductSpuSkuAttrMapOne(where string, args ...interface{}) (productSpuSkuAttrMap *ProductSpuSkuAttrMap) {
	err := DB.Model(productSpuSkuAttrMap).First(productSpuSkuAttrMap, where, args).GetErrors()
	if len(err) > 0 {
		return nil
	}
	return
}

func GetProductSpuSkuAttrMapList(page, limit int64, where string, condition interface{}) (list []*ProductSpuSkuAttrMap) {
	err := DB.Limit(limit).Offset((page-1)*limit).Find(list, where, condition).GetErrors()
	if err != nil {
		return nil
	}
	return
}

func (productSpuSkuAttrMap *ProductSpuSkuAttrMap) Create() []error {
	return DB.Model(productSpuSkuAttrMap).Create(productSpuSkuAttrMap).GetErrors()
}

func (productSpuSkuAttrMap *ProductSpuSkuAttrMap) Update(update ProductSpuSkuAttrMap) []error {
	return DB.Model(productSpuSkuAttrMap).UpdateColumns(update).GetErrors()
}

func (productSpuSkuAttrMap *ProductSpuSkuAttrMap) Delete() {
	DB.Model(productSpuSkuAttrMap).Delete(productSpuSkuAttrMap)
}
