package departement

import (
	"fmt"
	"log"

	"backend/models/global"

	"github.com/mashingan/smapping"
)

//DepartementService is a ....
type DepartementService interface {
	InsertDepartement(b global.DEPARTEMENT) global.DEPARTEMENT
	UpdateDepartement(b global.DEPARTEMENT) global.DEPARTEMENT
	DeleteDepartement(b global.DEPARTEMENT)
	GetAllDepartement() []global.DEPARTEMENT
	FindDepartementByID(modulID uint64) global.DEPARTEMENT
	IsAllowedToEdit(userID string, modulID uint64) bool
}

type departementService struct {
	departementRepository DepartementRepository
}

//NewDepartementService .....
func NewDepartementService(departementRepo DepartementRepository) DepartementService {
	return &departementService{
		departementRepository: departementRepo,
	}
}

func (service *departementService) InsertDepartement(b global.DEPARTEMENT) global.DEPARTEMENT {
	departement := global.DEPARTEMENT{}
	err := smapping.FillStruct(&departement, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.departementRepository.InsertDepartement(departement)
	return res
}

func (service *departementService) UpdateDepartement(b global.DEPARTEMENT) global.DEPARTEMENT {
	res := service.departementRepository.UpdateDepartement(b)
	return res
}

func (service *departementService) GetAllDepartement() []global.DEPARTEMENT {
	return service.departementRepository.GetAllDepartement()
}

func (service *departementService) FindDepartementByID(modulID uint64) global.DEPARTEMENT {
	return service.departementRepository.FindDepartementByID(modulID)
}

func (service *departementService) DeleteDepartement(b global.DEPARTEMENT) {
	service.departementRepository.DeleteDepartement(b)
}

func (service *departementService) IsAllowedToEdit(userID string, departementID uint64) bool {
	b := service.departementRepository.FindDepartementByID(departementID)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return userID == id
}
