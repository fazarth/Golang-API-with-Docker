package module

import (
	"fmt"
	"log"

	"backend/models/global"

	"github.com/mashingan/smapping"
)

//ModulesService is a ....
type ModulesService interface {
	InsertModules(b global.MODULE) global.MODULE
	UpdateModules(b global.MODULE) global.MODULE
	DeleteModules(b global.MODULE)
	GetAllModules() []global.MODULE
	FindModulesByID(modulID uint64) global.MODULE
	IsAllowedToEdit(userID string, modulID uint64) bool
}

type modulesService struct {
	modulesRepository ModulesRepository
}

//NewModulesService .....
func NewModulesService(modulesRepo ModulesRepository) ModulesService {
	return &modulesService{
		modulesRepository: modulesRepo,
	}
}

func (service *modulesService) InsertModules(b global.MODULE) global.MODULE {
	modules := global.MODULE{}
	err := smapping.FillStruct(&modules, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.modulesRepository.InsertModules(modules)
	return res
}

func (service *modulesService) UpdateModules(b global.MODULE) global.MODULE {
	res := service.modulesRepository.UpdateModules(b)
	return res
}

func (service *modulesService) GetAllModules() []global.MODULE {
	return service.modulesRepository.GetAllModules()
}

func (service *modulesService) FindModulesByID(modulID uint64) global.MODULE {
	return service.modulesRepository.FindModulesByID(modulID)
}

func (service *modulesService) DeleteModules(b global.MODULE) {
	service.modulesRepository.DeleteModules(b)
}

func (service *modulesService) IsAllowedToEdit(userID string, modulesID uint64) bool {
	b := service.modulesRepository.FindModulesByID(modulesID)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return userID == id
}
