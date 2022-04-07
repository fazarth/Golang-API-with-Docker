package partner

import (
	"fmt"
	"log"

	"backend/models/global"

	"github.com/mashingan/smapping"
)

//PartnerService is a ....
type PartnerService interface {
	InsertPartner(b global.PARTNER) global.PARTNER
	UpdatePartner(b global.PARTNER) global.PARTNER
	DeletePartner(b global.PARTNER)
	GetAllPartner() []global.PARTNER
	FindPartnerByID(parnerID uint64) global.PARTNER
	IsAllowedToEdit(userID string, parnerID uint64) bool
}

type modulesService struct {
	modulesRepository PartnerRepository
}

//NewPartnerService .....
func NewPartnerService(modulesRepo PartnerRepository) PartnerService {
	return &modulesService{
		modulesRepository: modulesRepo,
	}
}

func (service *modulesService) InsertPartner(b global.PARTNER) global.PARTNER {
	modules := global.PARTNER{}
	err := smapping.FillStruct(&modules, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.modulesRepository.InsertPartner(modules)
	return res
}

func (service *modulesService) UpdatePartner(b global.PARTNER) global.PARTNER {
	res := service.modulesRepository.UpdatePartner(b)
	return res
}

func (service *modulesService) GetAllPartner() []global.PARTNER {
	return service.modulesRepository.GetAllPartner()
}

func (service *modulesService) FindPartnerByID(parnerID uint64) global.PARTNER {
	return service.modulesRepository.FindPartnerByID(parnerID)
}

func (service *modulesService) DeletePartner(b global.PARTNER) {
	service.modulesRepository.DeletePartner(b)
}

func (service *modulesService) IsAllowedToEdit(userID string, modulesID uint64) bool {
	b := service.modulesRepository.FindPartnerByID(modulesID)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return userID == id
}
