package systemsetting

import (
	"fmt"
	"log"

	"backend/models/global"

	"github.com/mashingan/smapping"
)

//SystemSettingService is a ....
type SystemSettingService interface {
	InsertSystemSetting(b global.SYSTEMSETTING) global.SYSTEMSETTING
	UpdateSystemSetting(b global.SYSTEMSETTING) global.SYSTEMSETTING
	DeleteSystemSetting(b global.SYSTEMSETTING)
	GetAllSystemSetting() []global.SYSTEMSETTING
	FindSystemSettingByID(systemsettingsID uint64) global.SYSTEMSETTING
	IsAllowedToEdit(userID string, systemsettingsID uint64) bool
}

type modulesService struct {
	modulesRepository SystemSettingRepository
}

//NewSystemSettingService .....
func NewSystemSettingService(modulesRepo SystemSettingRepository) SystemSettingService {
	return &modulesService{
		modulesRepository: modulesRepo,
	}
}

func (service *modulesService) InsertSystemSetting(b global.SYSTEMSETTING) global.SYSTEMSETTING {
	modules := global.SYSTEMSETTING{}
	err := smapping.FillStruct(&modules, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.modulesRepository.InsertSystemSetting(modules)
	return res
}

func (service *modulesService) UpdateSystemSetting(b global.SYSTEMSETTING) global.SYSTEMSETTING {
	res := service.modulesRepository.UpdateSystemSetting(b)
	return res
}

func (service *modulesService) GetAllSystemSetting() []global.SYSTEMSETTING {
	return service.modulesRepository.GetAllSystemSetting()
}

func (service *modulesService) FindSystemSettingByID(systemsettingsID uint64) global.SYSTEMSETTING {
	return service.modulesRepository.FindSystemSettingByID(systemsettingsID)
}

func (service *modulesService) DeleteSystemSetting(b global.SYSTEMSETTING) {
	service.modulesRepository.DeleteSystemSetting(b)
}

func (service *modulesService) IsAllowedToEdit(userID string, modulesID uint64) bool {
	b := service.modulesRepository.FindSystemSettingByID(modulesID)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return userID == id
}
