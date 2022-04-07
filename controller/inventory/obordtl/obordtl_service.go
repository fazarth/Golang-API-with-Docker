package obordtl

import (
	"backend/models/inventory"
	"fmt"
	"log"

	"github.com/mashingan/smapping"
)

type OborDTLService interface {
	InsertOborDTL(b inventory.OBORDTL) inventory.OBORDTL
	ReadAllOborDTL() []inventory.OBORDTL
	FindOborDTLByID(task_id uint64) inventory.OBORDTL
	UpdateOborDTL(b inventory.OBORDTL) inventory.OBORDTL
	DeleteOborDTL(b inventory.OBORDTL)
	IsAllowedToEdit(createUser string, task_id uint64) bool
}

type OBORDTL_Service struct {
	oborDTLRepository OborDTLRepository
}

//NewOborDTLService .....
func NewOborDTLService(OborDTLRepo OborDTLRepository) OborDTLService {
	return &OBORDTL_Service{
		oborDTLRepository: OborDTLRepo,
	}
}

func (service *OBORDTL_Service) InsertOborDTL(b inventory.OBORDTL) inventory.OBORDTL {
	OborDTL := inventory.OBORDTL{}
	err := smapping.FillStruct(&OborDTL, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.oborDTLRepository.InsertOborDTL(OborDTL)
	return res
}

func (service *OBORDTL_Service) ReadAllOborDTL() []inventory.OBORDTL {
	return service.oborDTLRepository.ReadAllOborDTL()
}

func (service *OBORDTL_Service) FindOborDTLByID(task_id uint64) inventory.OBORDTL {
	return service.oborDTLRepository.FindOborDTLByID(task_id)
}

func (service *OBORDTL_Service) UpdateOborDTL(b inventory.OBORDTL) inventory.OBORDTL {
	res := service.oborDTLRepository.UpdateOborDTL(b)
	return res
}

func (service *OBORDTL_Service) DeleteOborDTL(b inventory.OBORDTL) {
	service.oborDTLRepository.DeleteOborDTL(b)
}

func (service *OBORDTL_Service) IsAllowedToEdit(createUser string, task_id uint64) bool {
	b := service.oborDTLRepository.FindOborDTLByID(task_id)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return createUser == id
}
