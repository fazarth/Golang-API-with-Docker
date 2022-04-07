package inventory

type CASE struct {
	CASE_ID          uint64  `gorm:"primary_key:auto_increment;column:CASE_ID" json:"case_id"`
	CASE_DISPLAY     string  `gorm:"column:CASE_DISPLAY;type:varchar(255);not null" json:"case_display"`
	CASE_BRCD        string  `gorm:"column:CASE_BRCD;type:varchar(255);not null" json:"case_brcd"`
	GREIGE_ID        uint64  `gorm:"column:GREIGE_ID;type:int(9);not null" json:"greige_id"`
	GREIGE_BARCODE   string  `gorm:"column:GREIGE_BARCODE;type:varchar(255);not null" json:"greige_barcode"`
	MARKETING_NAME   string  `gorm:"column:MARKETING_NAME;type:varchar(255);not null" json:"marketing_name"`
	ITEM_PAID        string  `gorm:"column:ITEM_PAID;type:varchar(255);not null" json:"item_paid"`
	CUST_NAME        string  `gorm:"column:CUST_NAME;type:varchar(255);not null" json:"cust_name"`
	CUST_ADR         string  `gorm:"column:CUST_ADR;type:varchar(255);not null" json:"cust_adr"`
	PIL_SOREFF       string  `gorm:"column:PIL_SOREFF;type:varchar(255);not null" json:"pil_soreff"`
	CUST_PO          string  `gorm:"column:CUST_PO;type:varchar(255);not null" json:"cust_po"`
	DESIGN_ID        string  `gorm:"column:DESIGN_ID;type:varchar(255);not null" json:"design_id"`
	SAP_ITEM_CODE    uint64  `gorm:"column:SAP_ITEM_CODE;type:int(9);not null" json:"sap_item_code"`
	SAP_ITEM_NAME    string  `gorm:"column:SAP_ITEM_NAME;type:varchar(255);not null" json:"sap_item_name"`
	SAP_ITEM_COLOR   string  `gorm:"column:SAP_ITEM_COLOR;type:varchar(255);not null" json:"sap_item_color"`
	SAP_COLOR_CODE   uint64  `gorm:"column:SAP_COLOR_CODE;type:int(9);not null" json:"sap_color_code"`
	SAP_ITEM_PROCESS string  `gorm:"column:SAP_ITEM_PROCESS;type:varchar(255);not null" json:"sap_item_process"`
	GRG_MAKLOON      string  `gorm:"column:GRG_MAKLOON;type:varchar(255);not null" json:"grg_makloon"`
	SAP_SO           uint64  `gorm:"column:SAP_SO;type:int(9);not null" json:"sap_so"`
	SAP_SO_LINE      uint64  `gorm:"column:SAP_SO_LINE;type:int(9);not null" json:"sap_so_line"`
	SAP_DUE_DATE     string  `gorm:"column:SAP_DUE_DATE;type:varchar(255);not null" json:"sap_due_date"`
	SAP_ORDER_TYPE   string  `gorm:"column:SAP_ORDER_TYPE;type:varchar(255);not null" json:"sap_order_type"`
	SAP_PDO_ROLL     uint64  `gorm:"column:SAP_PDO_ROLL;type:int(9);not null" json:"sap_pdo_roll"`
	SAP_PDO_WEIGHT   float64 `gorm:"column:SAP_PDO_WEIGHT;type:float;not null" json:"sap_pdo_weight"`
	KNIT_OPR_NAME    string  `gorm:"column:KNIT_OPR_NAME;type:varchar(255);not null" json:"knit_opr_name"`
	KNIT_MACHINE     string  `gorm:"column:KNIT_MACHINE;type:varchar(255);not null" json:"knit_machine"`
	KNIT_PROD_DATE   string  `gorm:"column:KNIT_PROD_DATE;type:varchar(255);not null" json:"knit_prod_date"`
	KNIT_PROD_SHIFT  uint64  `gorm:"column:KNIT_PROD_SHIFT;type:int(9);not null" json:"knit_prod_shift"`
	KNIT_QC_NAME     string  `gorm:"column:KNIT_QC_NAME;type:varchar(255);not null" json:"knit_qc_name"`
	KNIT_QC_TIME     string  `gorm:"column:KNIT_QC_TIME;type:varchar(255);not null" json:"knit_qc_time"`
	KNIT_SCALE_NAME  string  `gorm:"column:KNIT_SCALE_NAME;type:varchar(255);not null" json:"knit_scale_name"`
	KNIT_SCALE_TIME  string  `gorm:"column:KNIT_SCALE_TIME;type:varchar(255);not null" json:"knit_sacle_time"`
	KNIT_SCALE_BRUTO float64 `gorm:"column:KNIT_SCALE_BRUTO;type:float;not null" json:"knit_scale_bruto"`
	KNIT_SCALE_NETTO float64 `gorm:"column:KNIT_SCALE_NETTO;type:float;not null" json:"knit_scale_netto"`
	KNIT_LENGTH      float64 `gorm:"column:KNIT_LENGTH;type:float;not null" json:"knit_length"`
	KNIT_GSM         float64 `gorm:"column:KNIT_GSM;type:float;not null" json:"knit_gsm"`
	KNIT_REMARK      string  `gorm:"column:KNIT_REMARK;type:varchar(255);not null" json:"knit_remark"`
	KNIT_DO_NUM      uint64  `gorm:"column:KNIT_DO_NUM;type:int(9);not null" json:"knit_do_num"`
	KNIT_DO_DATE     string  `gorm:"column:KNIT_DO_DATE;type:varchar(255);not null" json:"knit_do_date"`
	PG_DYE_DATE      string  `gorm:"column:PG_DYE_DATE;type:varchar(255);not null" json:"pg_dye_date"`
	PG_DYE_OPR       string  `gorm:"column:PG_DYE_OPR;type:varchar(255);not null" json:"pg_dye_opr"`
	PG_QC_DATE       string  `gorm:"column:PG_QC_DATE;type:varchar(255);not null" json:"pg_qc_date"`
	PG_QC_OPR        string  `gorm:"column:PG_QC_OPR;type:varchar(255);not null" json:"pg_qc_opr"`
	PG_SCALE_DATE    string  `gorm:"column:PG_SCALE_DATE;type:varchar(255);not null" json:"pg_scale_date"`
	PG_SCALE_OPR     string  `gorm:"column:PG_SCALE_OPR;type:varchar(255);not null" json:"pg_scale_opr"`
	BATCH_ID         uint64  `gorm:"column:BATCH_ID;type:int(9);not null" json:"batch_id"`
	FINAL_WEIGHT     float64 `gorm:"column:FINAL_WEIGHT;type:float;not null" json:"final_weight"`
	COMMENT          string  `gorm:"column:COMMENT;type:varchar(255)" json:"comment"`
	ACTIVE           string  `gorm:"column:ACTIVE;type:char(1)" json:"active"`
	CREATE_USER      uint64  `gorm:"column:CREATE_USER;type:int(9)" json:"create_user"`
	CREATE_DATE      string  `gorm:"column:CREATE_DATE;type:datetime" json:"create_date"`
	UPDATE_USER      uint64  `gorm:"column:UPDATE_USER;type:int(9)" json:"update_user"`
	UPDATE_DATE      string  `gorm:"column:UPDATE_DATE;type:datetime" json:"update_date"`
	RECV_SHIP_ID     uint64  `gorm:"column:RECV_SHIP_ID;type:int(9);not null" json:"recv_ship_id"`
	DELV_SHIP_ID     uint64  `gorm:"column:DELV_SHIP_ID;type:int(9);not null" json:"delv-ship_id"`
	STAT_CODE        uint64  `gorm:"column:STAT_CODE;type:int(9);not null" json:"stat_code"`
	PREV_LOCN_ID     uint64  `gorm:"column:PREV_LOCN_ID;type:int(9);not null" json:"prev_locn_id"`
	LOCN_ID          uint64  `gorm:"column:LOCN_ID;type:int(9);not null" json:"locn_id"`
	DEST_LOCN_ID     uint64  `gorm:"column:DEST_LOCN_ID;type:int(9);not null" json:"dest_locn_id"`
	LAST_LOCK_DATE   string  `gorm:"column:LAST_LOCK_DATE;type:varchar(255);not null" json:"last_lock_id"`
	LOCK             string  `gorm:"column:LOCK;type:varchar(255);not null" json:"lock"`
	PACK_QTY         float64 `gorm:"column:PACK_QTY;type:float;not null" json:"pack_qty"`
	PACK_TYPE        string  `gorm:"column:PACK_TYPE;type:varchar(255);not null" json:"pack_type"`
}
