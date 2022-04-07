package division

import (
	"fmt"
	"log"

	"backend/models/global"

	"github.com/mashingan/smapping"
)

//DivisionService is a ....
type DivisionService interface {
	InsertDivision(b global.DIVISION) global.DIVISION
	UpdateDivision(b global.DIVISION) global.DIVISION
	DeleteDivision(b global.DIVISION)
	GetAllDivision() []global.DIVISION
	FindDivisionByID(divID uint64) global.DIVISION
	IsAllowedToEdit(userID string, divID uint64) bool
}

type divisionService struct {
	divisionRepository DivisionRepository
}

//NewDivisionService .....
func NewDivisionService(divisionRepo DivisionRepository) DivisionService {
	return &divisionService{
		divisionRepository: divisionRepo,
	}
}

func (service *divisionService) InsertDivision(b global.DIVISION) global.DIVISION {
	division := global.DIVISION{}
	err := smapping.FillStruct(&division, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.divisionRepository.InsertDivision(division)
	return res
}

func (service *divisionService) UpdateDivision(b global.DIVISION) global.DIVISION {
	res := service.divisionRepository.UpdateDivision(b)
	return res
}

func (service *divisionService) GetAllDivision() []global.DIVISION {
	return service.divisionRepository.GetAllDivision()
}

func (service *divisionService) FindDivisionByID(divID uint64) global.DIVISION {
	return service.divisionRepository.FindDivisionByID(divID)
}

func (service *divisionService) DeleteDivision(b global.DIVISION) {
	service.divisionRepository.DeleteDivision(b)
}

func (service *divisionService) IsAllowedToEdit(userID string, divisionID uint64) bool {
	b := service.divisionRepository.FindDivisionByID(divisionID)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return userID == id
}
