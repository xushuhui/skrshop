package model

type SkrMember struct {
	Id       uint   `gorm:"column:id;type:int unsigned" json:"id"`                       // 用户id
	Uid      uint   `gorm:"column:uid;type:int unsigned;default:'0'" json:"uid"`         // 账号id
	Nickname string `gorm:"column:nickname;type:varchar(30);default:''" json:"nickname"` // 昵称
	Avatar   string `gorm:"column:avatar;type:varchar(255);default:''" json:"avatar"`    // 头像(相对路径)
	Gender   int8   `gorm:"column:gender;type:tinyint(1);default:'1'" json:"gender"`     // 员工性别 1 unknow 2 male 3 female
	Role     uint8  `gorm:"column:role;type:tinyint unsigned;default:'0'" json:"role"`   // 角色 0:普通用户 1:vip
	CreateAt int    `gorm:"column:create_at;type:int;default:'0'" json:"create_at"`      // 创建时间
	UpdateAt int    `gorm:"column:update_at;type:int;default:'0'" json:"update_at"`      // 更新时间
}

//get real primary key name
func (skrMember *SkrMember) GetKey() string {
	return "id"
}

//get primary key in model
func (skrMember *SkrMember) GetKeyProperty() uint {
	return skrMember.Id
}

//set primary key
func (skrMember *SkrMember) SetKeyProperty(id uint) {
	skrMember.Id = id
}

//get real table name
func (skrMember *SkrMember) TableName() string {
	return "skr_member"
}

func GetSkrMemberById(id string) (skrMember *SkrMember) {
	err := DB.Model(skrMember).First(skrMember, skrMember.GetKey()+" = '"+id+"'").GetErrors()
	if len(err) > 0 {
		return nil
	}
	return
}

func GetSkrMemberOne(where string, args ...interface{}) (skrMember *SkrMember) {
	err := DB.Model(skrMember).First(skrMember, where, args).GetErrors()
	if len(err) > 0 {
		return nil
	}
	return
}

func GetSkrMemberList(page, limit int64, where string, condition interface{}) (list []*SkrMember) {
	err := DB.Limit(limit).Offset((page-1)*limit).Find(list, where, condition).GetErrors()
	if err != nil {
		return nil
	}
	return
}

func (skrMember *SkrMember) Create() []error {
	return DB.Model(skrMember).Create(skrMember).GetErrors()
}

func (skrMember *SkrMember) Update(update SkrMember) []error {
	return DB.Model(skrMember).UpdateColumns(update).GetErrors()
}

func (skrMember *SkrMember) Delete() {
	DB.Model(skrMember).Delete(skrMember)
}
