package global

type PARTNER struct {
	PARTNER_ID       uint64 `gorm:"column:PARTNER_ID;primary_key:auto_increment;type:int(6);not null" json:"partner_id"`
	TYPE_PARTNER     string `gorm:"column:TYPE_PARTNER;type:varchar(255)" json:"type_partner"`
	PARTNER_NAME     string `gorm:"column:PARTNER_NAME;type:varchar(255)" json:"partner_name"`
	EMAIL            string `gorm:"column:EMAIL;type:varchar(255)" json:"email"`
	USERNAME         string `gorm:"column:USERNAME;type:varchar(255)" json:"username"`
	PASSWORD         string `gorm:"column:PASSWORD;type:varchar(255)" json:"password"`
	PARTNER_PROFILE  string `gorm:"column:PARTNER_PROFILE;type:varchar(255)" json:"partner_profile"`
	PHONE_NUM        string `gorm:"column:PHONE_NUM;type:varchar(255)" json:"phone_num"`
	WEBSITE_URL      string `gorm:"column:WEBSITE_URL;type:varchar(255)" json:"website_url"`
	ADDRESS          string `gorm:"column:ADDRESS;type:varchar(255)" json:"address"`
	CITY             string `gorm:"column:CITY;type:varchar(255)" json:"city"`
	STATE            string `gorm:"column:STATE;type:varchar(255)" json:"state"`
	ZIP_CODE         string `gorm:"column:ZIP_CODE;type:varchar(255)" json:"zip_code"`
	COUNTRY          string `gorm:"column:COUNTRY;type:varchar(255)" json:"country"`
	NPWP             string `gorm:"column:NPWP;type:varchar(255)" json:"npwp"`
	TAX              string `gorm:"column:TAX;type:varchar(255)" json:"tax"`
	LAST_LOGIN_DATE  string `gorm:"column:LAST_LOGIN_DATE;type:datetime" json:"last_login_date"`
	LAST_LOGOUT_DATE string `gorm:"column:LAST_LOGOUT_DATE;type:datetime" json:"last_logout_date"`
	IS_LOGGED_IN     string `gorm:"column:IS_LOGGED_IN;type:varchar(255)" json:"is_logged_in"`
	COMMENT          string `gorm:"column:COMMENT;type:varchar(255)" json:"comment"`
	ACTIVE           string `gorm:"column:ACTIVE;type:char(1)" json:"active"`
	CREATE_USER      uint64 `gorm:"column:CREATE_USER;type:int(9)" json:"create_user"`
	CREATE_DATE      string `gorm:"column:CREATE_DATE;type:datetime" json:"create_date"`
	UPDATE_USER      uint64 `gorm:"column:UPDATE_USER;type:int(9)" json:"update_user"`
	UPDATE_DATE      string `gorm:"column:UPDATE_DATE;type:datetime" json:"update_date"`
}
