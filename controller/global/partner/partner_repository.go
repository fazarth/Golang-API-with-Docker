package partner

import (
	"backend/models/global"

	"gorm.io/gorm"
)

//ModuleRepository is contract what userRepository can do to db
type PartnerRepository interface {
	InsertPartner(b global.PARTNER) global.PARTNER
	GetAllPartner() []global.PARTNER
	FindPartnerByID(partnerID uint64) global.PARTNER
	UpdatePartner(b global.PARTNER) global.PARTNER
	DeletePartner(b global.PARTNER)
}

type partnerConnection struct {
	connection *gorm.DB
}

//NewPartnerRepository creates an instance PartnerRepository
func NewPartnerRepository(dbConn *gorm.DB) PartnerRepository {
	return &partnerConnection{
		connection: dbConn,
	}
}

func (db *partnerConnection) InsertPartner(b global.PARTNER) global.PARTNER {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *partnerConnection) GetAllPartner() []global.PARTNER {
	var partner []global.PARTNER
	db.connection.Find(&partner)
	return partner
}

func (db *partnerConnection) FindPartnerByID(partnerID uint64) global.PARTNER {
	var partner global.PARTNER
	db.connection.Find(&partner, partnerID)
	return partner
}

func (db *partnerConnection) UpdatePartner(b global.PARTNER) global.PARTNER {
	var partner global.PARTNER
	db.connection.Find(&partner).Where("PARTNER_ID = ?", b.PARTNER_ID)
	partner.COMMENT = b.COMMENT
	db.connection.Updates(&b)
	return b
}

func (db *partnerConnection) DeletePartner(b global.PARTNER) {
	db.connection.Delete(&b)
}
