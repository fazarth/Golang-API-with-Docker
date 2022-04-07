package usergroup

import (
	"fmt"
	"log"

	"backend/models/global"

	"github.com/mashingan/smapping"
)

//UserGroupService is a ....
type UserGroupService interface {
	InsertUserGroup(b global.USERGROUP) global.USERGROUP
	UpdateUserGroup(b global.USERGROUP) global.USERGROUP
	DeleteUserGroup(b global.USERGROUP)
	GetAllUserGroup() []global.USERGROUP
	FindUserGroupByID(userGroupID uint64) global.USERGROUP
	IsAllowedToEdit(userID string, userGroupID uint64) bool
}

type userGroupIDService struct {
	userGroupIDRepository UserGroupRepository
}

//NewUserGroupService .....
func NewUserGroupService(userGroupIDRepo UserGroupRepository) UserGroupService {
	return &userGroupIDService{
		userGroupIDRepository: userGroupIDRepo,
	}
}

func (service *userGroupIDService) InsertUserGroup(b global.USERGROUP) global.USERGROUP {
	userGroupID := global.USERGROUP{}
	err := smapping.FillStruct(&userGroupID, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.userGroupIDRepository.InsertUserGroup(userGroupID)
	return res
}

func (service *userGroupIDService) UpdateUserGroup(b global.USERGROUP) global.USERGROUP {
	res := service.userGroupIDRepository.UpdateUserGroup(b)
	return res
}

func (service *userGroupIDService) GetAllUserGroup() []global.USERGROUP {
	return service.userGroupIDRepository.GetAllUserGroup()
}

func (service *userGroupIDService) FindUserGroupByID(userGroupID uint64) global.USERGROUP {
	return service.userGroupIDRepository.FindUserGroupByID(userGroupID)
}

func (service *userGroupIDService) DeleteUserGroup(b global.USERGROUP) {
	service.userGroupIDRepository.DeleteUserGroup(b)
}

func (service *userGroupIDService) IsAllowedToEdit(userID string, userGroupIDID uint64) bool {
	b := service.userGroupIDRepository.FindUserGroupByID(userGroupIDID)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return userID == id
}
