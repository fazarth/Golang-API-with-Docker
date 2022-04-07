package global

type LOG struct {
	USER_LOGS_ID uint64 `gorm:"column:USER_LOGS_ID;primary_key:auto_increment;type:int(6);not null" json:"user_logs_id"`
	CODE_TRX     string `gorm:"column:CODE_TRX;type:varchar(255)" json:"code_trx"`
	ACTION_DESC  string `gorm:"column:ACTION_DESC;type:varchar(255)" json:"action_desc"`
	ROUTE_PAGE   string `gorm:"column:ROUTE_PAGE;type:varchar(255)" json:"route_page"`
	USER_ID      uint64 `gorm:"column:USER_ID;type:int(6)" json:"user_id"`
	USER_IP      string `gorm:"column:USER_IP;type:varchar(255)" json:"user_ip"`
	USER_DEVICE  string `gorm:"column:USER_DEVICE;type:varchar(255)" json:"user_device"`
	COMMENT      string `gorm:"column:COMMENT;type:varchar(255)" json:"comment"`
	ACTIVE       string `gorm:"column:ACTIVE;type:char(1)" json:"active"`
	CREATE_USER  uint64 `gorm:"column:CREATE_USER;type:int(9)" json:"create_user"`
	CREATE_DATE  string `gorm:"column:CREATE_DATE;type:datetime" json:"create_date"`
	UPDATE_USER  uint64 `gorm:"column:UPDATE_USER;type:int(9)" json:"update_user"`
	UPDATE_DATE  string `gorm:"column:UPDATE_DATE;type:datetime" json:"update_date"`
}
