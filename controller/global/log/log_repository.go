package log

import (
	"backend/models/global"

	"gorm.io/gorm"
)

//ModuleRepository is contract what userRepository can do to db
type LogRepository interface {
	InsertLog(b global.LOG) global.LOG
	GetAllLog() []global.LOG
	FindLogByID(logID uint64) global.LOG
	UpdateLog(b global.LOG) global.LOG
	DeleteLog(b global.LOG)
}

type logConnection struct {
	connection *gorm.DB
}

//NewLogRepository creates an instance LogRepository
func NewLogRepository(dbConn *gorm.DB) LogRepository {
	return &logConnection{
		connection: dbConn,
	}
}

func (db *logConnection) InsertLog(b global.LOG) global.LOG {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *logConnection) GetAllLog() []global.LOG {
	var log []global.LOG
	db.connection.Find(&log)
	return log
}

func (db *logConnection) FindLogByID(logID uint64) global.LOG {
	var log global.LOG
	db.connection.Find(&log, logID)
	return log
}

func (db *logConnection) UpdateLog(b global.LOG) global.LOG {
	var log global.LOG
	db.connection.Find(&log).Where("USER_LOGS_ID = ?", b.USER_LOGS_ID)
	log.COMMENT = b.COMMENT
	db.connection.Updates(&b)
	return b
}

func (db *logConnection) DeleteLog(b global.LOG) {
	db.connection.Delete(&b)
}
