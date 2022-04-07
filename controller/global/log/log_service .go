package log

import (
	"fmt"
	"log"

	"backend/models/global"

	"github.com/mashingan/smapping"
)

//LogService is a ....
type LogService interface {
	InsertLog(b global.LOG) global.LOG
	UpdateLog(b global.LOG) global.LOG
	DeleteLog(b global.LOG)
	GetAllLog() []global.LOG
	FindLogByID(logID uint64) global.LOG
	IsAllowedToEdit(userID string, logID uint64) bool
}

type logService struct {
	logRepository LogRepository
}

//NewLogService .....
func NewLogService(logRepo LogRepository) LogService {
	return &logService{
		logRepository: logRepo,
	}
}

func (service *logService) InsertLog(b global.LOG) global.LOG {
	logs := global.LOG{}
	err := smapping.FillStruct(&logs, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.logRepository.InsertLog(logs)
	return res
}

func (service *logService) UpdateLog(b global.LOG) global.LOG {
	res := service.logRepository.UpdateLog(b)
	return res
}

func (service *logService) GetAllLog() []global.LOG {
	return service.logRepository.GetAllLog()
}

func (service *logService) FindLogByID(logID uint64) global.LOG {
	return service.logRepository.FindLogByID(logID)
}

func (service *logService) DeleteLog(b global.LOG) {
	service.logRepository.DeleteLog(b)
}

func (service *logService) IsAllowedToEdit(userID string, logID uint64) bool {
	b := service.logRepository.FindLogByID(logID)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return userID == id
}
