package global

type MODULE struct {
	MODULE_ID     uint64 `gorm:"column:MODULE_ID;primary_key:auto_increment;type:int(6);not null" json:"module_id"`
	MODULE_NAME   string `gorm:"column:MODULE_NAME;type:varchar(255)" json:"module_name"`
	IRL           string `gorm:"column:IRL;type:varchar(255)" json:"irl"`
	ICON          string `gorm:"column:ICON;type:varchar(255)" json:"icon"`
	MODULE_SEQ    string `gorm:"column:MODULE_SEQ;type:varchar(255)" json:"module_seq"`
	PARENT_MODULE string `gorm:"column:PARENT_MODULE;type:varchar(255)" json:"parent_module"`
	COMMENT       string `gorm:"column:COMMENT;type:varchar(255);default:''" json:"comment"`
	ACTIVE        string `gorm:"column:ACTIVE;type:char(1);default:''" json:"active"`
	CREATE_USER   uint64 `gorm:"column:CREATE_USER;type:int(9)" json:"create_user"`
	CREATE_DATE   string `gorm:"column:CREATE_DATE;type:datetime" json:"create_date"`
	UPDATE_USER   uint64 `gorm:"column:UPDATE_USER;type:int(9)" json:"update_user"`
	UPDATE_DATE   string `gorm:"column:UPDATE_DATE;type:datetime" json:"update_date"`
}
