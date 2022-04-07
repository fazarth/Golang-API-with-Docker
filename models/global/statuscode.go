package global

type STATUSCODE struct {
	STATUS_ID   uint64 `gorm:"primary_key:auto_increment;column:STATUS_ID;type:int(6);not null" json:"status_id"`
	STATUS_CODE uint64 `gorm:"column:STATUS_CODE;type:int(6);;not null" json:"status_code"`
	STATUS_NAME string `gorm:"column:STATUS_NAME;type:varchar(255)" json:"status_name"`
	COMMENT     string `gorm:"column:COMMENT;type:varchar(255)" json:"comment"`
	ACTIVE      string `gorm:"column:ACTIVE;type:char(1)" json:"active"`
	CREATE_USER uint64 `gorm:"column:CREATE_USER;type:int(9)" json:"create_user"`
	CREATE_DATE string `gorm:"column:CREATE_DATE;type:datetime" json:"create_date"`
	UPDATE_USER uint64 `gorm:"column:UPDATE_USER;type:int(9)" json:"update_user"`
	UPDATE_DATE string `gorm:"column:UPDATE_DATE;type:datetime" json:"update_date"`
}
