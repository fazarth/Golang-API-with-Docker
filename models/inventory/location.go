package inventory

type LOCATION struct {
	LOCN_ID          uint64 `gorm:"primary_key:auto_increment;column:LOCN_ID;type:int(9);not null" json:"locn_id"`
	WHSE             uint64 `gorm:"column:WHSE;type:int(4);not null" json:"whse"`
	CHECK_DIGIT      uint64 `gorm:"column:CHECK_DIGIT;type:int(4);not null" json:"check_digit"`
	LOCN_DISPLAY     string `gorm:"column:LOCN_DISPLAY;type:varchar(255);not null" json:"locn_display"`
	LOCN_BRCD        string `gorm:"column:LOCN_BRCD;type:varchar(255);not null" json:"locn_brcd"`
	ZONE             string `gorm:"column:ZONE;type:char(1);not null" json:"zone"`
	AISLE            string `gorm:"column:AISLE;type:varchar(9);not null" json:"aisle"`
	BAY              uint64 `gorm:"column:BAY;type:int(3);not null" json:"bay"`
	LVL              uint64 `gorm:"column:LVL;type:int(3);not null" json:"lvl"`
	X_COORD          uint64 `gorm:"column:X_COORD;type:int(3);not null" json:"x_coord"`
	Y_COORD          uint64 `gorm:"column:Y_COORD;type:int(3);not null" json:"y_coord"`
	Z_COORD          uint64 `gorm:"column:Z_COORD;type:int(3);not null" json:"z_coord"`
	IB_SEQ           uint64 `gorm:"column:IB_SEQ;type:int(9);not null" json:"ib_seq"`
	OB_SEQ           uint64 `gorm:"column:OB_SEQ;type:int(9);not null" json:"ob_seq"`
	LEN              uint64 `gorm:"column:LEN;type:float(53);not null" json:"len"`
	WID              uint64 `gorm:"column:WID;type:float(53);not null" json:"wid"`
	HGT              uint64 `gorm:"column:HGT;type:float(53);not null" json:"hgt"`
	CURR_VOL         uint64 `gorm:"column:CURR_VOL;type:int(9);not null" json:"curr_vol"`
	CURR_CONT        uint64 `gorm:"column:CURR_CONT;type:int(9);not null" json:"curr_cont"`
	CURR_ITEM        uint64 `gorm:"column:CURR_ITEM;type:int(9);not null" json:"curr_item"`
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
