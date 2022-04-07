package module

import (
	"backend/models/global"

	"gorm.io/gorm"
)

//ModuleRepository is contract what userRepository can do to db
type ModulesRepository interface {
	InsertModules(b global.MODULE) global.MODULE
	GetAllModules() []global.MODULE
	FindModulesByID(modulesID uint64) global.MODULE
	UpdateModules(b global.MODULE) global.MODULE
	DeleteModules(b global.MODULE)
}

type modulesConnection struct {
	connection *gorm.DB
}

//NewModulesRepository creates an instance ModulesRepository
func NewModulesRepository(dbConn *gorm.DB) ModulesRepository {
	return &modulesConnection{
		connection: dbConn,
	}
}

func (db *modulesConnection) InsertModules(b global.MODULE) global.MODULE {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *modulesConnection) GetAllModules() []global.MODULE {
	var modules []global.MODULE
	db.connection.Find(&modules)
	return modules
}

func (db *modulesConnection) FindModulesByID(modulesID uint64) global.MODULE {
	var modules global.MODULE
	db.connection.Find(&modules, modulesID)
	return modules
}

func (db *modulesConnection) UpdateModules(b global.MODULE) global.MODULE {
	var modules global.MODULE
	db.connection.Find(&modules).Where("MODULE_ID = ?", b.MODULE_ID)
	modules.COMMENT = b.COMMENT
	db.connection.Updates(&b)
	return b
}

func (db *modulesConnection) DeleteModules(b global.MODULE) {
	db.connection.Delete(&b)
}
