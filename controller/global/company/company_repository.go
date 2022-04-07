package company

import (
	"backend/models/global"

	"gorm.io/gorm"
)

//ModuleRepository is contract what userRepository can do to db
type CompanyRepository interface {
	InsertCompany(b global.COMPANY) global.COMPANY
	GetAllCompany() []global.COMPANY
	FindCompanyByID(companyID uint64) global.COMPANY
	UpdateCompany(b global.COMPANY) global.COMPANY
	DeleteCompany(b global.COMPANY)
}

type companyConnection struct {
	connection *gorm.DB
}

//NewCompanyRepository creates an instance CompanyRepository
func NewCompanyRepository(dbConn *gorm.DB) CompanyRepository {
	return &companyConnection{
		connection: dbConn,
	}
}

func (db *companyConnection) InsertCompany(b global.COMPANY) global.COMPANY {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *companyConnection) GetAllCompany() []global.COMPANY {
	var getcompany []global.COMPANY
	db.connection.Find(&getcompany)
	return getcompany
}

func (db *companyConnection) FindCompanyByID(companyID uint64) global.COMPANY {
	var findcompanyID global.COMPANY
	db.connection.Find(&findcompanyID, companyID)
	return findcompanyID
}

func (db *companyConnection) UpdateCompany(b global.COMPANY) global.COMPANY {
	var updatecompany global.COMPANY
	db.connection.Find(&updatecompany).Where("COMPANY_ID = ?", b.COMPANY_ID)
	updatecompany.COMMENT = b.COMMENT
	db.connection.Updates(&b)
	return b
}

func (db *companyConnection) DeleteCompany(b global.COMPANY) {
	db.connection.Delete(&b)
}
