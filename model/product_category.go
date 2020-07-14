package model

type ProductCategory struct {
	Id       uint   `gorm:"column:id;type:int unsigned" json:"id"`                           // 分类ID
	Pid      uint   `gorm:"column:pid;type:int unsigned;default:'0'" json:"pid"`             // 父ID
	Name     string `gorm:"column:name;type:varchar(255);default:''" json:"name"`            // 分类名称
	Desc     string `gorm:"column:desc;type:varchar(255);default:''" json:"desc"`            // 分类描述
	PicUrl   string `gorm:"column:pic_url;type:varchar(255);default:''" json:"pic_url"`      // 分类图片
	Path     string `gorm:"column:path;type:varchar(255);default:''" json:"path"`            // 分类地址{pid}-{child_id}-...
	CreateAt uint   `gorm:"column:create_at;type:int unsigned;default:'0'" json:"create_at"` // 创建时间
	CreateBy uint   `gorm:"column:create_by;type:int unsigned;default:'0'" json:"create_by"` // 创建人staff_id
	UpdateAt uint   `gorm:"column:update_at;type:int unsigned;default:'0'" json:"update_at"` // 更新时间
	UpdateBy uint   `gorm:"column:update_by;type:int unsigned;default:'0'" json:"update_by"` // 修改人staff_id
	Status   uint8  `gorm:"column:status;type:tinyint unsigned;default:'0'" json:"status"`   // 状态 1:enable, 0:disable, -1:deleted
}

//get real primary key name
func (productCategory *ProductCategory) GetKey() string {
	return "id"
}

//get primary key in model
func (productCategory *ProductCategory) GetKeyProperty() uint {
	return productCategory.Id
}

//set primary key
func (productCategory *ProductCategory) SetKeyProperty(id uint) {
	productCategory.Id = id
}

//get real table name
func (productCategory *ProductCategory) TableName() string {
	return "product_category"
}

func GetProductCategoryById(id string) (productCategory *ProductCategory, err error) {
	err = DB.Model(productCategory).First(productCategory, productCategory.GetKey()+" = '"+id+"'").Error

	return
}

func GetProductCategoryOne(where string, args ...interface{}) (productCategory *ProductCategory, err error) {
	err = DB.Model(productCategory).First(productCategory, where, args).Error

	return
}

func GetProductCategoryList(page, limit int64, where string, condition interface{}) (list []*ProductCategory, err error) {
	err = DB.Limit(limit).Offset((page-1)*limit).Find(list, where, condition).Error

	return
}

func (productCategory *ProductCategory) Create() (err error) {
	return DB.Model(productCategory).Create(productCategory).Error
}

func (productCategory *ProductCategory) Delete() {
	DB.Model(productCategory).Delete(productCategory)
}
