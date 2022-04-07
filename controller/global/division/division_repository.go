package division

import (
	"backend/models/global"

	"gorm.io/gorm"
)

//ModuleRepository is contract what userRepository can do to db
type DivisionRepository interface {
	InsertDivision(b global.DIVISION) global.DIVISION
	GetAllDivision() []global.DIVISION
	FindDivisionByID(divID uint64) global.DIVISION
	UpdateDivision(b global.DIVISION) global.DIVISION
	DeleteDivision(b global.DIVISION)
}

type modulesConnection struct {
	connection *gorm.DB
}

//NewDivisionRepository creates an instance DivisionRepository
func NewDivisionRepository(dbConn *gorm.DB) DivisionRepository {
	return &modulesConnection{
		connection: dbConn,
	}
}

func (db *modulesConnection) InsertDivision(b global.DIVISION) global.DIVISION {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *modulesConnection) GetAllDivision() []global.DIVISION {
	var modules []global.DIVISION
	db.connection.Find(&modules)
	return modules
}

func (db *modulesConnection) FindDivisionByID(divID uint64) global.DIVISION {
	var modules global.DIVISION
	db.connection.Find(&modules, divID)
	return modules
}

func (db *modulesConnection) UpdateDivision(b global.DIVISION) global.DIVISION {
	var modules global.DIVISION
	db.connection.Find(&modules).Where("DIV_ID = ?", b.DIV_ID)
	modules.COMMENT = b.COMMENT
	db.connection.Updates(&b)
	return b
}

func (db *modulesConnection) DeleteDivision(b global.DIVISION) {
	db.connection.Delete(&b)
}
