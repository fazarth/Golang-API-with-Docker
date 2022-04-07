package inventory

type OBORHDR struct {
	ORDER_ID      uint64 `gorm:"primary_key:auto_increment;column:ORDER_ID;type:int(9);not null" json:"order_id"`
	DOC_ID        string `gorm:"column:DOC_ID;type:varchar(255)" json:"doc_id"`
	TASK_TYPE     string `gorm:"column:TASK_TYPE;type:varchar(255)" json:"task_type"`
	STAT_CODE     uint64 `gorm:"column:STAT_CODE;type:int(9)" json:"stat_code"`
	TASK_PRIOTITY uint64 `gorm:"column:TASK_PRIOTITY;type:int(9)" json:"task_priority"`
	START_TIME    string `gorm:"column:START_TIME;type:datetime" json:"start_time"`
	FINISH_TIME   string `gorm:"column:FINISH_TIME;type:datetime" json:"finish_time"`
	COMMENT       string `gorm:"column:COMMENT;type:varchar(255)" json:"comment"`
	ACTIVE        string `gorm:"column:ACTIVE;type:char(1)" json:"active"`
	CREATE_USER   uint64 `gorm:"column:CREATE_USER;type:int(9)" json:"create_user"`
	CREATE_DATE   string `gorm:"column:CREATE_DATE;type:datetime" json:"create_date"`
	UPDATE_USER   uint64 `gorm:"column:UPDATE_USER;type:int(9)" json:"update_user"`
	UPDATE_DATE   string `gorm:"column:UPDATE_DATE;type:datetime" json:"update_date"`
}
