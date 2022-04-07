package hrd

type EMPLOYEE struct {
	EMPLOYEE_ID        uint64 `gorm:"primary_key:auto_increment;column:EMPLOYEE_ID" json:"employee_id"`
	FIRST_NAME         string `gorm:"column:FIRST_NAME;type:varchar(255);not null" json:"first_name"`
	SURE_NAME          string `gorm:"column:SURE_NAME;type:varchar(255);not null" json:"sure_name"`
	LAST_NAME          uint64 `gorm:"column:LAST_NAME;type:int(9);not null" json:"last_name"`
	EMAIL              string `gorm:"column:EMAIL;type:varchar(255);not null" json:"email"`
	DATE_OF_BIRTH      string `gorm:"column:DATE_OF_BIRTH;type:datetime;not null" json:"date_of_birth"`
	GENDER             string `gorm:"column:GENDER;type:varchar(255);not null" json:"gender"`
	USER_GROUP_ID      uint64 `gorm:"column:USER_GROUP_ID;type:int(9);not null" json:"user_group_id"`
	COMPANY_ID         uint64 `gorm:"column:COMPANY_ID;type:int(9);not null" json:"company_id"`
	DEPT_ID            uint64 `gorm:"column:DEPT_ID;type:int(9);not null" json:"dept_id"`
	DIV_ID             uint64 `gorm:"column:DIV_ID;type:int(9);not null" json:"div_id"`
	POSITION           string `gorm:"column:POSITION;type:varchar(255);not null" json:"position"`
	GRADE              uint64 `gorm:"column:GRADE;type:int(9);not null" json:"grade"`
	STATUS_EMPLOYEE    string `gorm:"column:STATUS_EMPLOYEE;type:varchar(255);not null" json:"status_employee"`
	SALARY_TEMPLATE    string `gorm:"column:SALARY_TEMPLATE;type:varchar(255)" json:"salary_template"`
	DATE_JOINING       string `gorm:"column:DATE_JOINING;type:datetime" json:"date_joining"`
	DATE_LEAVING       string `gorm:"column:DATE_LEAVING;type:datetime" json:"date_leaving"`
	MARITIAL_STATUS    string `gorm:"column:MARITIAL_STATUS;type:varchar(255)" json:"maritial_status"`
	BASIC_SALARY       string `gorm:"column:BASIC_SALARY;type:varchar(255)" json:"basic_salary"`
	ADDRESS            string `gorm:"column:ADDRESS;type:varchar(255)" json:"address"`
	STATE              string `gorm:"column:STATE;type:varchar(255)" json:"state"`
	CITY               string `gorm:"column:CITY;type:varchar(255)" json:"city"`
	ZIPCODE            string `gorm:"column:ZIPCODE;type:varchar(255)" json:"zipcode"`
	PROFILE_PICTURE    string `gorm:"column:PROFILE_PICTURE;type:varchar(255)" json:"profile_picture"`
	PROFILE_BACKGROUND string `gorm:"column:PROFILE_BACKGROUND;type:varchar(255)" json:"profile_background"`
	RESUME             string `gorm:"column:RESUME;type:varchar(255)" json:"resume"`
	PHONE_NUM          string `gorm:"column:PHONE_NUM;type:varchar(255)" json:"phone_num"`
	FACEBOOK_LINK      string `gorm:"column:FACEBOOK_LINK;type:varchar(255)" json:"facebook_link"`
	TWITTER_LINK       string `gorm:"column:TWITTER_LINK;type:varchar(255)" json:"twittter_link"`
	INSTAGRAM_LINK     string `gorm:"column:INSTAGRAM_LINK;type:varchar(255)" json:"instagram_link"`
	LINKEDIN_LINK      string `gorm:"column:LINKEDIN_LINK;type:varchar(255)" json:"linkedin_link"`
	OTHER_LINK         string `gorm:"column:OTHER_LINK;type:varchar(255)" json:"other_link"`
	LAST_LOGIN_DATE    string `gorm:"column:LAST_LOGIN_DATE;type:varchar(255)" json:"last_login_date"`
	LAST_LOGOUT_DATE   string `gorm:"column:LAST_LOGOUT_DATE;type:varchar(255)" json:"last_logout_date"`
	LAST_LOGIN_IP      string `gorm:"column:LAST_LOGIN_IP;type:varchar(255)" json:"las_login_ip"`
	IS_LOGGED_IN       string `gorm:"column:IS_LOGGED_IN;type:varchar(255)" json:"is_logged_in"`
	ONLINE_STATUS      string `gorm:"column:ONLINE_STATUS;type:varchar(255)" json:"online_status"`
	COMMENT            string `gorm:"column:COMMENT;type:varchar(255)" json:"comment"`
	ACTIVE             string `gorm:"column:ACTIVE;type:char(1)" json:"active"`
	CREATE_USER        uint64 `gorm:"column:CREATE_USER;type:int(9)" json:"create_user"`
	CREATE_DATE        string `gorm:"column:CREATE_DATE;type:datetime" json:"create_date"`
	UPDATE_USER        uint64 `gorm:"column:UPDATE_USER;type:int(9)" json:"update_user"`
	UPDATE_DATE        string `gorm:"column:UPDATE_DATE;type:datetime" json:"update_date"`
}
