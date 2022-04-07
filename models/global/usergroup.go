package global

type USERGROUP struct {
	USER_GROUP_ID   uint64 `gorm:"column:USER_GROUP_ID;primary_key:auto_increment;type:int(6);not null" json:"user_group_id"`
	USER_ID         uint64 `gorm:"column:USER_ID;type:int(6);not null" json:"user_id"`
	PERMISSION_ID   uint64 `gorm:"column:PERMISSION_ID;type:int(6);not null" json:"permission_id"`
	COMPANY_ID      uint64 `gorm:"column:COMPANY_ID;type:int(6);not null" json:"company_id"`
	USER_GROUP_NAME string `gorm:"column:USER_GROUP_NAME;type:varchar(255)" json:"user_group_name"`
	COMMENT         string `gorm:"column:COMMENT;type:varchar(255)" json:"comment"`
	ACTIVE          string `gorm:"column:ACTIVE;type:char(1)" json:"active"`
	CREATE_USER     uint64 `gorm:"column:CREATE_USER;type:int(9)" json:"create_user"`
	CREATE_DATE     string `gorm:"column:CREATE_DATE;type:datetime" json:"create_date"`
	UPDATE_USER     uint64 `gorm:"column:UPDATE_USER;type:int(9)" json:"update_user"`
	UPDATE_DATE     string `gorm:"column:UPDATE_DATE;type:datetime" json:"update_date"`
}
