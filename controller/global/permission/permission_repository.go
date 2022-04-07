package permission

import (
	"backend/models/global"

	"gorm.io/gorm"
)

//ModuleRepository is contract what userRepository can do to db
type PermissionRepository interface {
	InsertPermission(b global.PERMISSION) global.PERMISSION
	GetAllPermission() []global.PERMISSION
	FindPermissionByID(permissionID uint64) global.PERMISSION
	UpdatePermission(b global.PERMISSION) global.PERMISSION
	DeletePermission(b global.PERMISSION)
}

type permissionConnection struct {
	connection *gorm.DB
}

//NewPermissionRepository creates an instance PermissionRepository
func NewPermissionRepository(dbConn *gorm.DB) PermissionRepository {
	return &permissionConnection{
		connection: dbConn,
	}
}

func (db *permissionConnection) InsertPermission(b global.PERMISSION) global.PERMISSION {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *permissionConnection) GetAllPermission() []global.PERMISSION {
	var permission []global.PERMISSION
	db.connection.Find(&permission)
	return permission
}

func (db *permissionConnection) FindPermissionByID(permissionID uint64) global.PERMISSION {
	var permission global.PERMISSION
	db.connection.Find(&permission, permissionID)
	return permission
}

func (db *permissionConnection) UpdatePermission(b global.PERMISSION) global.PERMISSION {
	var permission global.PERMISSION
	db.connection.Find(&permission).Where("USER_PERMISSION_ID = ?", b.USER_PERMISSION_ID)
	permission.COMMENT = b.COMMENT
	db.connection.Updates(&b)
	return b
}

func (db *permissionConnection) DeletePermission(b global.PERMISSION) {
	db.connection.Delete(&b)
}
