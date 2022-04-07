package departement

import (
	"backend/models/global"

	"gorm.io/gorm"
)

//ModuleRepository is contract what userRepository can do to db
type DepartementRepository interface {
	InsertDepartement(b global.DEPARTEMENT) global.DEPARTEMENT
	GetAllDepartement() []global.DEPARTEMENT
	FindDepartementByID(departementID uint64) global.DEPARTEMENT
	UpdateDepartement(b global.DEPARTEMENT) global.DEPARTEMENT
	DeleteDepartement(b global.DEPARTEMENT)
}

type departementConnection struct {
	connection *gorm.DB
}

//NewDepartementRepository creates an instance DepartementRepository
func NewDepartementRepository(dbConn *gorm.DB) DepartementRepository {
	return &departementConnection{
		connection: dbConn,
	}
}

func (db *departementConnection) InsertDepartement(b global.DEPARTEMENT) global.DEPARTEMENT {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *departementConnection) GetAllDepartement() []global.DEPARTEMENT {
	var departement []global.DEPARTEMENT
	db.connection.Find(&departement)
	return departement
}

func (db *departementConnection) FindDepartementByID(departementID uint64) global.DEPARTEMENT {
	var departement global.DEPARTEMENT
	db.connection.Find(&departement, departementID)
	return departement
}

func (db *departementConnection) UpdateDepartement(b global.DEPARTEMENT) global.DEPARTEMENT {
	var departement global.DEPARTEMENT
	db.connection.Find(&departement).Where("DEPT_ID = ?", b.DEPT_ID)
	departement.COMMENT = b.COMMENT
	db.connection.Updates(&b)
	return b
}

func (db *departementConnection) DeleteDepartement(b global.DEPARTEMENT) {
	db.connection.Delete(&b)
}
