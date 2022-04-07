package oborhdr

import (
	"backend/models/inventory"
	"fmt"
	"log"

	"github.com/mashingan/smapping"
)

type OborHDRService interface {
	InsertOborHDR(b inventory.OBORHDR) inventory.OBORHDR
	ReadAllOborHDR() []inventory.OBORHDR
	FindOborHDRByID(order_id uint64) inventory.OBORHDR
	UpdateOborHDR(b inventory.OBORHDR) inventory.OBORHDR
	DeleteOborHDR(b inventory.OBORHDR)
	IsAllowedToEdit(createUser string, order_id uint64) bool
}

type OBORHDR_Service struct {
	locationRepository OborHDRRepository
}

//NewOborHDRService .....
func NewOborHDRService(OborHDRRepo OborHDRRepository) OborHDRService {
	return &OBORHDR_Service{
		locationRepository: OborHDRRepo,
	}
}

func (service *OBORHDR_Service) InsertOborHDR(b inventory.OBORHDR) inventory.OBORHDR {
	OborHDR := inventory.OBORHDR{}
	err := smapping.FillStruct(&OborHDR, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.locationRepository.InsertOborHDR(OborHDR)
	return res
}

func (service *OBORHDR_Service) ReadAllOborHDR() []inventory.OBORHDR {
	return service.locationRepository.ReadAllOborHDR()
}

func (service *OBORHDR_Service) FindOborHDRByID(order_id uint64) inventory.OBORHDR {
	return service.locationRepository.FindOborHDRByID(order_id)
}

func (service *OBORHDR_Service) UpdateOborHDR(b inventory.OBORHDR) inventory.OBORHDR {
	res := service.locationRepository.UpdateOborHDR(b)
	return res
}

func (service *OBORHDR_Service) DeleteOborHDR(b inventory.OBORHDR) {
	service.locationRepository.DeleteOborHDR(b)
}

func (service *OBORHDR_Service) IsAllowedToEdit(createUser string, order_id uint64) bool {
	b := service.locationRepository.FindOborHDRByID(order_id)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return createUser == id
}
