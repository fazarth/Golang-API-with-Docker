package global

type USER struct {
	USER_ID          uint64 `gorm:"column:USER_ID;primary_key:auto_increment;type:int(9)" json:"user_id"`
	EMPLOYEE_ID      uint64 `gorm:"column:EMPLOYEE_ID;type:int(9)" json:"employee_id"`
	PARTNER_ID       uint64 `gorm:"column:PARTNER_ID;type:int(9);not null" json:"partner_id"`
	USER_GROUP_ID    uint64 `gorm:"column:USER_GROUP_ID;type:int(9);not null" json:"user_group_id"`
	USERNAME         string `gorm:"column:USERNAME;type:varchar(255)" json:"username"`
	PASSWORD         string `gorm:"column:PASSWORD;->;<-;not null" json:"-"`
	TOKEN            string `gorm:"-" json:"token,omitempty"`
	LAST_LOGIN_DATE  string `gorm:"column:LAST_LOGIN_DATE;type:datetime" json:"last_login_date"`
	LAST_LOGOUT_DATE string `gorm:"column:LAST_LOGOUT_DATE;type:datetime" json:"last_logout_date"`
	LAST_LOGIN_IP    string `gorm:"column:LAST_LOGIN_IP;type:varchar(255)" json:"last_login_ip"`
	IS_LOGGED_IN     string `gorm:"column:IS_LOGGED_IN;type:varchar(255)" json:"is_logged_in"`
	ONLINE_STATUS    string `gorm:"column:ONLINE_STATUS;type:varchar(255)" json:"online_status"`
	COMMENT          string `gorm:"column:COMMENT;type:varchar(255)" json:"comment"`
	ACTIVE           string `gorm:"column:ACTIVE;type:char(1)" json:"active"`
	CREATE_USER      uint64 `gorm:"column:CREATE_USER;type:int(9)" json:"create_user"`
	CREATE_DATE      string `gorm:"column:CREATE_DATE;type:datetime" json:"create_date"`
	UPDATE_USER      uint64 `gorm:"column:UPDATE_USER;type:int(9)" json:"update_user"`
	UPDATE_DATE      string `gorm:"column:UPDATE_DATE;type:datetime" json:"update_date"`
}

type CredentialsLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
