package global

import (
	"backend/models/global"

	"gorm.io/gorm"
)

//GlobalRepository is contract what GlobalRepository can do to db
type GlobalRepository interface {
	InsertGlobal(b global.GLOBAL) global.GLOBAL
	GetAllGlobal() []global.GLOBAL
	FindGlobalByID(globalID uint64) global.GLOBAL
	UpdateGlobal(b global.GLOBAL) global.GLOBAL
	DeleteGlobal(b global.GLOBAL)
}

type globalConnection struct {
	connection *gorm.DB
}

//NewGlobalRepository creates an instance GlobalRepository
func NewGlobalRepository(dbConn *gorm.DB) GlobalRepository {
	return &globalConnection{
		connection: dbConn,
	}
}

func (db *globalConnection) InsertGlobal(b global.GLOBAL) global.GLOBAL {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *globalConnection) GetAllGlobal() []global.GLOBAL {
	var global []global.GLOBAL
	db.connection.Find(&global)
	return global
}

func (db *globalConnection) FindGlobalByID(globalID uint64) global.GLOBAL {
	var global global.GLOBAL
	db.connection.Find(&global, globalID)
	return global
}

func (db *globalConnection) UpdateGlobal(b global.GLOBAL) global.GLOBAL {
	var global global.GLOBAL
	db.connection.Find(&global).Where("GLOBAL_ID = ?", b.GLOBAL_ID)
	global.COMMENT = b.COMMENT
	db.connection.Updates(&b)
	return b
}

func (db *globalConnection) DeleteGlobal(b global.GLOBAL) {
	db.connection.Delete(&b)
}
