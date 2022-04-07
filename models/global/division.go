package global

type DIVISION struct {
	DIV_ID           uint64 `gorm:"column:DIV_ID;primary_key:auto_increment;type:int(6);not null" json:"div_id"`
	DEPT_ID          uint64 `gorm:"column:DEPT_ID;type:int(6);not null" json:"dept_id"`
	DIV_NAME         string `gorm:"column:DIV_NAME;type:varchar(255)" json:"div_name"`
	SUB_DIV_ID       uint64 `gorm:"column:SUB_DIV_ID;type:int(6);not null" json:"sub_div_id"`
	USER_EMPLOYEE_ID uint64 `gorm:"column:USER_EMPLOYEE_ID;type:int(6);not null" json:"user_employee_id"`
	COMMENT          string `gorm:"column:COMMENT;type:varchar(255)" json:"comment"`
	ACTIVE           string `gorm:"column:ACTIVE;type:char(1)" json:"active"`
	CREATE_USER      uint64 `gorm:"column:CREATE_USER;type:int(9)" json:"create_user"`
	CREATE_DATE      string `gorm:"column:CREATE_DATE;type:datetime" json:"create_date"`
	UPDATE_USER      uint64 `gorm:"column:UPDATE_USER;type:int(9)" json:"update_user"`
	UPDATE_DATE      string `gorm:"column:UPDATE_DATE;type:datetime" json:"update_date"`
}
