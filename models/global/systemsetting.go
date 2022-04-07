package global

type SYSTEMSETTING struct {
	SYSTEM_SETTING_ID        uint64 `gorm:"column:SYSTEM_SETTING_ID;primary_key:auto_increment;type:int(11);not null" json:"system_setting_id"`
	APPLICATION_NAME         string `gorm:"column:APPLICATION_NAME;type:varchar(255)" json:"application_name"`
	DEFAULT_CURENCY          string `gorm:"column:DEFAULT_CURENCY;type:varchar(255)" json:"default_curency"`
	DEFAULT_CURENCY_SYMBOL   string `gorm:"column:DEFAULT_CURENCY_SYMBOL;type:varchar(255)" json:"default_curency_symbol"`
	LOGIN_WITH               string `gorm:"column:LOGIN_WITH;type:varchar(255)" json:"login_with"`
	APPLICATION_LOGO         string `gorm:"column:APPLICATION_LOGO;type:varchar(255)" json:"application_logo"`
	FOOTER_TEXT              string `gorm:"column:FOOTER_TEXT;type:varchar(255)" json:"footer_text"`
	DEFAULT_LANGUAGE         string `gorm:"column:DEFAULT_LANGUAGE;type:varchar(255)" json:"default_language"`
	SYSTEM_TIMEZONE          string `gorm:"column:SYSTEM_TIMEZONE;type:datetime" json:"system_timezone"`
	SYSTEM_IP                string `gorm:"column:SYSTEM_IP;type:varchar(255)" json:"system_ip"`
	GOOGLE_MAPS_APIKEY       string `gorm:"column:GOOGLE_MAPS_APIKEY;type:varchar(255)" json:"google_maps_apikey"`
	APPLICATION_RELEASE_DATE string `gorm:"column:APPLICATION_RELEASE_DATE;type:datetime" json:"application_release_date"`
	COMMENT                  string `gorm:"column:COMMENT;type:varchar(255)" json:"comment"`
	ACTIVE                   string `gorm:"column:ACTIVE;type:char(1)" json:"active"`
	CREATE_USER              uint64 `gorm:"column:CREATE_USER;type:int(9)" json:"create_user"`
	CREATE_DATE              string `gorm:"column:CREATE_DATE;type:datetime" json:"create_date"`
	UPDATE_USER              uint64 `gorm:"column:UPDATE_USER;type:int(9)" json:"update_user"`
	UPDATE_DATE              string `gorm:"column:UPDATE_DATE;type:datetime" json:"update_date"`
}
