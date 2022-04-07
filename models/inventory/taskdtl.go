package inventory

type TASKDTL struct {
	TASK_ID      uint64 `gorm:"primary_key:auto_increment;column:TASK_ID;type:int(9);not null" json:"task_id"`
	TASK_SEQ_NO  uint64 `gorm:"column:TASK_SEQ_NO;type:int(9)" json:"task_seq_no"`
	ITEM_ID      uint64 `gorm:"column:ITEM_ID;type:int(9)" json:"item_id"`
	PULL_LOCN_ID uint64 `gorm:"column:PULL_LOCN_ID;type:int(9)" json:"pull_locn_id"`
	DEST_LOCN_ID uint64 `gorm:"column:DEST_LOCN_ID;type:int(9)" json:"dest_locn_id"`
	PULL_CONT_ID uint64 `gorm:"column:PULL_CONT_ID;type:int(9)" json:"pull_cont_id"`
	DEST_CONT_ID uint64 `gorm:"column:DEST_CONT_ID;type:int(9)" json:"dest_cont_id"`
	QTY_REQUEST  uint64 `gorm:"column:QTY_REQUEST;type:int(9)" json:"qty_request"`
	QTY_ACTUAL   uint64 `gorm:"column:QTY_ACTUAL;type:int(9)" json:"qty_actual"`
	START_TIME   string `gorm:"column:START_TIME;type:datetime" json:"start_time"`
	FINISH_TIME  string `gorm:"column:FINISH_TIME;type:datetime" json:"finish_time"`
	MISC_FIELD_1 string `gorm:"column:MISC_FIELD_1;type:varchar(255)" json:"misc_field_1"`
	MISC_FIELD_2 string `gorm:"column:MISC_FIELD_2;type:varchar(255)" json:"misc_field_2"`
	MISC_FIELD_3 string `gorm:"column:MISC_FIELD_3;type:varchar(255)" json:"misc_field_3"`
	MISC_FIELD_4 string `gorm:"column:MISC_FIELD_4;type:varchar(255)" json:"misc_field_4"`
	MISC_FIELD_5 string `gorm:"column:MISC_FIELD_5;type:varchar(255)" json:"misc_field_5"`
	COMMENT      string `gorm:"column:COMMENT;type:varchar(255)" json:"comment"`
	ACTIVE       string `gorm:"column:ACTIVE;type:char(1)" json:"active"`
	CREATE_USER  uint64 `gorm:"column:CREATE_USER;type:int(9)" json:"create_user"`
	CREATE_DATE  string `gorm:"column:CREATE_DATE;type:datetime" json:"create_date"`
	UPDATE_USER  uint64 `gorm:"column:UPDATE_USER;type:int(9)" json:"update_user"`
	UPDATE_DATE  string `gorm:"column:UPDATE_DATE;type:datetime" json:"update_date"`
	PREV_USER_ID uint64 `gorm:"column:PREV_USER_ID;type:int(9)" json:"prev_user_id"`
	CURR_USER_ID uint64 `gorm:"column:CURR_USER_ID;type:int(9)" json:"curr_user_id"`
	STAT_CODE    uint64 `gorm:"column:STAT_CODE;type:int(9)" json:"stat_code"`
}
