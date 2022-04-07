package usergroup

import (
	"backend/models/global"

	"gorm.io/gorm"
)

//UserGroupRepository is contract what userRepository can do to db
type UserGroupRepository interface {
	InsertUserGroup(b global.USERGROUP) global.USERGROUP
	GetAllUserGroup() []global.USERGROUP
	FindUserGroupByID(userGroupID uint64) global.USERGROUP
	UpdateUserGroup(b global.USERGROUP) global.USERGROUP
	DeleteUserGroup(b global.USERGROUP)
}

type userGroupConnection struct {
	connection *gorm.DB
}

//NewUserGroupRepository creates an instance UserGroupRepository
func NewUserGroupRepository(dbConn *gorm.DB) UserGroupRepository {
	return &userGroupConnection{
		connection: dbConn,
	}
}

func (db *userGroupConnection) InsertUserGroup(b global.USERGROUP) global.USERGROUP {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *userGroupConnection) GetAllUserGroup() []global.USERGROUP {
	var userGroup []global.USERGROUP
	db.connection.Find(&userGroup)
	return userGroup
}

func (db *userGroupConnection) FindUserGroupByID(userGroupID uint64) global.USERGROUP {
	var userGroup global.USERGROUP
	db.connection.Find(&userGroup, userGroupID)
	return userGroup
}

func (db *userGroupConnection) UpdateUserGroup(b global.USERGROUP) global.USERGROUP {
	var userGroup global.USERGROUP
	db.connection.Find(&userGroup).Where("USER_GROUP_ID = ?", b.USER_GROUP_ID)
	userGroup.COMMENT = b.COMMENT
	db.connection.Updates(&b)
	return b
}

func (db *userGroupConnection) DeleteUserGroup(b global.USERGROUP) {
	db.connection.Delete(&b)
}
