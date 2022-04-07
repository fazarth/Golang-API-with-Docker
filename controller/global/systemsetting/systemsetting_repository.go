package systemsetting

import (
	"backend/models/global"

	"gorm.io/gorm"
)

//ModuleRepository is contract what userRepository can do to db
type SystemSettingRepository interface {
	InsertSystemSetting(b global.SYSTEMSETTING) global.SYSTEMSETTING
	GetAllSystemSetting() []global.SYSTEMSETTING
	FindSystemSettingByID(systemsettingsID uint64) global.SYSTEMSETTING
	UpdateSystemSetting(b global.SYSTEMSETTING) global.SYSTEMSETTING
	DeleteSystemSetting(b global.SYSTEMSETTING)
}

type systemsettingsConnection struct {
	connection *gorm.DB
}

//NewSystemSettingRepository creates an instance SystemSettingRepository
func NewSystemSettingRepository(dbConn *gorm.DB) SystemSettingRepository {
	return &systemsettingsConnection{
		connection: dbConn,
	}
}

func (db *systemsettingsConnection) InsertSystemSetting(b global.SYSTEMSETTING) global.SYSTEMSETTING {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *systemsettingsConnection) GetAllSystemSetting() []global.SYSTEMSETTING {
	var systemsettings []global.SYSTEMSETTING
	db.connection.Find(&systemsettings)
	return systemsettings
}

func (db *systemsettingsConnection) FindSystemSettingByID(systemsettingsID uint64) global.SYSTEMSETTING {
	var systemsettings global.SYSTEMSETTING
	db.connection.Find(&systemsettings, systemsettingsID)
	return systemsettings
}

func (db *systemsettingsConnection) UpdateSystemSetting(b global.SYSTEMSETTING) global.SYSTEMSETTING {
	var systemsettings global.SYSTEMSETTING
	db.connection.Find(&systemsettings).Where("SYSTEM_SETTING_ID = ?", b.SYSTEM_SETTING_ID)
	systemsettings.COMMENT = b.COMMENT
	db.connection.Updates(&b)
	return b
}

func (db *systemsettingsConnection) DeleteSystemSetting(b global.SYSTEMSETTING) {
	db.connection.Delete(&b)
}
