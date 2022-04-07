package permission

import (
	"fmt"
	"log"

	"backend/models/global"

	"github.com/mashingan/smapping"
)

//PermissionService is a ....
type PermissionService interface {
	InsertPermission(b global.PERMISSION) global.PERMISSION
	UpdatePermission(b global.PERMISSION) global.PERMISSION
	DeletePermission(b global.PERMISSION)
	GetAllPermission() []global.PERMISSION
	FindPermissionByID(permissionID uint64) global.PERMISSION
	IsAllowedToEdit(userID string, permissionID uint64) bool
}

type permissionService struct {
	permissionRepository PermissionRepository
}

//NewPermissionService .....
func NewPermissionService(permissionRepo PermissionRepository) PermissionService {
	return &permissionService{
		permissionRepository: permissionRepo,
	}
}

func (service *permissionService) InsertPermission(b global.PERMISSION) global.PERMISSION {
	permission := global.PERMISSION{}
	err := smapping.FillStruct(&permission, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.permissionRepository.InsertPermission(permission)
	return res
}

func (service *permissionService) UpdatePermission(b global.PERMISSION) global.PERMISSION {
	res := service.permissionRepository.UpdatePermission(b)
	return res
}

func (service *permissionService) GetAllPermission() []global.PERMISSION {
	return service.permissionRepository.GetAllPermission()
}

func (service *permissionService) FindPermissionByID(permissionID uint64) global.PERMISSION {
	return service.permissionRepository.FindPermissionByID(permissionID)
}

func (service *permissionService) DeletePermission(b global.PERMISSION) {
	service.permissionRepository.DeletePermission(b)
}

func (service *permissionService) IsAllowedToEdit(userID string, permissionID uint64) bool {
	b := service.permissionRepository.FindPermissionByID(permissionID)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return userID == id
}
