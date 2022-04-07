package global

type COMPANY struct {
	COMPANY_ID      uint64 `gorm:"primary_key:auto_increment;column:COMPANY_ID;type:int(6);not null" json:"company_id"`
	COMPANY_NAME    string `gorm:"column:COMPANY_NAME;type:varchar(255);not null" json:"company_name"`
	EMAIL           string `gorm:"column:EMAIL;type:varchar(255)" json:"email"`
	LOGO            string `gorm:"column:LOGO;type:varchar(255)" json:"logo"`
	PHONE           string `gorm:"column:PHONE;type:varchar(255)" json:"phone"`
	WEBSITE_URL     string `gorm:"column:WEBSITE_URL;type:varchar(255)" json:"website_url"`
	ADDRESS         string `gorm:"column:ADDRESS;type:varchar(255)" json:"address"`
	CITY            string `gorm:"column:CITY;type:varchar(255)" json:"city"`
	STATE           string `gorm:"column:STATE;type:varchar(255)" json:"state"`
	ZIPCODE         string `gorm:"column:ZIPCODE;type:varchar(255)" json:"zip_code"`
	COUNTRY         string `gorm:"column:COUNTRY;type:varchar(255)" json:"country"`
	NPWP            string `gorm:"column:NPWP;type:varchar(255)" json:"npwp"`
	REGISTRATION_NO uint64 `gorm:"column:REGISTRATION_NO;type:int(11)" json:"registration_no"`
	TAX             string `gorm:"column:TAX;type:varchar(255)" json:"tax"`
	COMMENT         string `gorm:"column:COMMENT;type:varchar(255)" json:"comment"`
	ACTIVE          string `gorm:"column:ACTIVE;type:char(1)" json:"active"`
	CREATE_USER     uint64 `gorm:"column:CREATE_USER;type:int(9)" json:"create_user"`
	CREATE_DATE     string `gorm:"column:CREATE_DATE;type:datetime" json:"create_date"`
	UPDATE_USER     uint64 `gorm:"column:UPDATE_USER;type:int(9)" json:"update_user"`
	UPDATE_DATE     string `gorm:"column:UPDATE_DATE;type:datetime" json:"update_date"`
}
