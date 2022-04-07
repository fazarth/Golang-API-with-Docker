package statuscode

import (
	"backend/models/global"

	"gorm.io/gorm"
)

//ModuleRepository is contract what userRepository can do to db
type StatusCodeRepository interface {
	InsertStatusCode(b global.STATUSCODE) global.STATUSCODE
	GetAllStatusCode() []global.STATUSCODE
	FindStatusCodeByID(companyID uint64) global.STATUSCODE
	UpdateStatusCode(b global.STATUSCODE) global.STATUSCODE
	DeleteStatusCode(b global.STATUSCODE)
}

type companyConnection struct {
	connection *gorm.DB
}

//NewStatusCodeRepository creates an instance StatusCodeRepository
func NewStatusCodeRepository(dbConn *gorm.DB) StatusCodeRepository {
	return &companyConnection{
		connection: dbConn,
	}
}

func (db *companyConnection) InsertStatusCode(b global.STATUSCODE) global.STATUSCODE {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *companyConnection) GetAllStatusCode() []global.STATUSCODE {
	var getcompany []global.STATUSCODE
	db.connection.Find(&getcompany)
	return getcompany
}

func (db *companyConnection) FindStatusCodeByID(companyID uint64) global.STATUSCODE {
	var findcompanyID global.STATUSCODE
	db.connection.Find(&findcompanyID, companyID)
	return findcompanyID
}

func (db *companyConnection) UpdateStatusCode(b global.STATUSCODE) global.STATUSCODE {
	var updatecompany global.STATUSCODE
	db.connection.Find(&updatecompany).Where("STATUSCODE_ID = ?", b.STATUSCODE_ID)
	updatecompany.COMMENT = b.COMMENT
	db.connection.Updates(&b)
	return b
}

func (db *companyConnection) DeleteStatusCode(b global.STATUSCODE) {
	db.connection.Delete(&b)
}
