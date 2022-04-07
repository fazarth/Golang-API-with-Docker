package inventory

type CONTAINER struct {
	CONT_ID          uint64 `gorm:"primary_key:auto_increment;column:CONT_ID;type:int(9);not null" json:"cont_id"`
	CONT_DISPLAY     uint64 `gorm:"column:CONT_DISPLAY;type:varchar(255);not null" json:"cont_display"`
	CONT_BRCD        uint64 `gorm:"column:CONT_BRCD;type:varchar(255);not null" json:"cont_brcd"`
	PREV_LOCN_ID     uint64 `gorm:"column:PREV_LOCN_ID;type:int(9);not null" json:"prev_locn_id"`
	CURR_LOCN_ID     uint64 `gorm:"column:CURR_LOCN_ID;type:int(9);not null" json:"curr-locn_id"`
	DEST_LOCN_ID     uint64 `gorm:"column:DEST_LOCN_ID;type:int(9);not null" json:"dest_locn_id"`
	PREV_TASK_HDR_ID uint64 `gorm:"column:PREV_TASK_HDR_ID;type:int(9);not null" json:"prev_task_hdr_id"`
	CURR_TASK_HDR_ID uint64 `gorm:"column:CURR_TASK_HDR_ID;type:int(9);not null" json:"curr_task_hdr_id"`
	CONT_LEN         uint64 `gorm:"column:CONT_LEN;type:float(9);not null" json:"cont_len"`
	CONT_WID         uint64 `gorm:"column:CONT_WID;type:float(9);not null" json:"cont_wid"`
	CONT_HGT         uint64 `gorm:"column:CONT_HGT;type:float(9);not null" json:"cont_hgt"`
	CURR_VOL         uint64 `gorm:"column:CURR_VOL;type:float(9);not null" json:"curr_vol"`
	CURR_ITEM        uint64 `gorm:"column:CURR_ITEM;type:int(9);not null" json:"curr_item"`
	CURR_WEIGHT_BRT  uint64 `gorm:"column:CURR_WEIGHT_BRT;type:float(9);not null" json:"curr_weigth_brt"`
	CURR_WEIGHT_NET  uint64 `gorm:"column:CURR_WEIGHT_NET;type:float(9);not null" json:"curr_weigth_net"`
	LAST_LOCK_DATE   uint64 `gorm:"column:LAST_LOCK_DATE;type:datetime;not null" json:"last_lock_date"`
	LOCK             uint64 `gorm:"column:LOCK;type:int(9);not null" json:"lock"`
	LAST_OPNAME_DATE uint64 `gorm:"column:LAST_OPNAME_DATE;type:datetime;not null" json:"last_opname_date"`
	OPNAME_PENDING   uint64 `gorm:"column:OPNAME_PENDING;type:char(1);not null" json:"opname_pending"`
	COMMENT          string `gorm:"column:COMMENT;type:varchar(255)" json:"comment"`
	ACTIVE           string `gorm:"column:ACTIVE;type:char(1)" json:"active"`
	CREATE_USER      uint64 `gorm:"column:CREATE_USER;type:int(9)" json:"create_user"`
	CREATE_DATE      string `gorm:"column:CREATE_DATE;type:datetime" json:"create_date"`
	UPDATE_USER      uint64 `gorm:"column:UPDATE_USER;type:int(9)" json:"update_user"`
	UPDATE_DATE      string `gorm:"column:UPDATE_DATE;type:datetime" json:"update_date"`
	DOC_NUM_ISO      uint64 `gorm:"column:DOC_NUM_ISO;type:varchar(255);not null" json:"doc_num_iso"`
}
