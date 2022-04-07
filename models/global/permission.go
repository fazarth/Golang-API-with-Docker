package global

type PERMISSION struct {
	USER_PERMISSION_ID uint64 `gorm:"primary_key:auto_increment;column:USER_PERMISSION_ID;type:int(6);not null" json:"user_permission_id"`
	MODULE_GROUP       string `gorm:"column:MODULE_GROUP;type:varchar(255)" json:"module_group"`
	MODULE_ATTRIBUTE   uint64 `gorm:"column:MODULE_ATTRIBUTE;type:int(6);not null" json:"module_attribute"`
	COMMENT            string `gorm:"column:COMMENT;type:varchar(255)" json:"comment"`
	ACTIVE             string `gorm:"column:ACTIVE;type:char(1)" json:"active"`
	CREATE_USER        uint64 `gorm:"column:CREATE_USER;type:int(9)" json:"create_user"`
	CREATE_DATE        string `gorm:"column:CREATE_DATE;type:datetime" json:"create_date"`
	UPDATE_USER        uint64 `gorm:"column:UPDATE_USER;type:int(9)" json:"update_user"`
	UPDATE_DATE        string `gorm:"column:UPDATE_DATE;type:datetime" json:"update_date"`
}
