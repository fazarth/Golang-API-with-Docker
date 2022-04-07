package global

type GLOBAL struct {
	GLOBAL_ID   uint64 `gorm:"column:GLOBAL_ID;primary_key:auto_increment;type:int(6);not null" json:"global_id"`
	GLOBAL_TYPE string `gorm:"column:GLOBAL_TYPE;type:varchar(255)" json:"global_type"`
	CODE        string `gorm:"column:CODE;type:varchar(255)" json:"code"`
	ALIAS       string `gorm:"column:ALIAS;type:varchar(255)" json:"alias"`
	VALUE       string `gorm:"column:VALUE;type:varchar(255)" json:"value"`
	COMMENT     string `gorm:"column:COMMENT;type:varchar(255)" json:"comment"`
	ACTIVE      string `gorm:"column:ACTIVE;type:char(1)" json:"active"`
	CREATE_USER uint64 `gorm:"column:CREATE_USER;type:int(9)" json:"create_user"`
	CREATE_DATE string `gorm:"column:CREATE_DATE;type:datetime" json:"create_date"`
	UPDATE_USER uint64 `gorm:"column:UPDATE_USER;type:int(9)" json:"update_user"`
	UPDATE_DATE string `gorm:"column:UPDATE_DATE;type:datetime" json:"update_date"`
}
