package global

import (
	"fmt"
	"log"

	"backend/models/global"

	"github.com/mashingan/smapping"
)

//GlobalService is a ....
type GlobalService interface {
	InsertGlobal(b global.GLOBAL) global.GLOBAL
	UpdateGlobal(b global.GLOBAL) global.GLOBAL
	DeleteGlobal(b global.GLOBAL)
	GetAllGlobal() []global.GLOBAL
	FindGlobalByID(globalID uint64) global.GLOBAL
	IsAllowedToEdit(userID string, globalID uint64) bool
}

type globalService struct {
	globalRepository GlobalRepository
}

//NewGlobalService .....
func NewGlobalService(globalRepo GlobalRepository) GlobalService {
	return &globalService{
		globalRepository: globalRepo,
	}
}

func (service *globalService) InsertGlobal(b global.GLOBAL) global.GLOBAL {
	global := global.GLOBAL{}
	err := smapping.FillStruct(&global, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.globalRepository.InsertGlobal(global)
	return res
}

func (service *globalService) UpdateGlobal(b global.GLOBAL) global.GLOBAL {
	res := service.globalRepository.UpdateGlobal(b)
	return res
}

func (service *globalService) GetAllGlobal() []global.GLOBAL {
	return service.globalRepository.GetAllGlobal()
}

func (service *globalService) FindGlobalByID(globalID uint64) global.GLOBAL {
	return service.globalRepository.FindGlobalByID(globalID)
}

func (service *globalService) DeleteGlobal(b global.GLOBAL) {
	service.globalRepository.DeleteGlobal(b)
}

func (service *globalService) IsAllowedToEdit(userID string, globalID uint64) bool {
	b := service.globalRepository.FindGlobalByID(globalID)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return userID == id
}
