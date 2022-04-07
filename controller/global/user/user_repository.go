package user

import (
	"backend/models/global"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type UsersRepository interface {
	InsertUser(user global.USER) global.USER
	UpdateUser(user global.USER) global.USER
	VerifyCredential(username string, password string) interface{}
	IsDuplicateUserName(username string) (tx *gorm.DB)
	FindByUserName(username string) global.USER
	ProfileUser(userID string) global.USER
}

type usersConnection struct {
	connection *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository
func NewUsersRepository(dbConn *gorm.DB) UsersRepository {
	return &usersConnection{
		connection: dbConn,
	}
}

func (db *usersConnection) InsertUser(user global.USER) global.USER {
	user.PASSWORD = hashAndSalt([]byte(user.PASSWORD))
	db.connection.Save(&user)
	return user
}

func (db *usersConnection) UpdateUser(user global.USER) global.USER {
	if user.PASSWORD != "" {
		user.PASSWORD = hashAndSalt([]byte(user.PASSWORD))
	} else {
		var tempUser global.USER
		db.connection.Find(&tempUser, user.USER_ID)
		user.PASSWORD = tempUser.PASSWORD
	}
	db.connection.Save(&user)
	return user
}

func (db *usersConnection) VerifyCredential(USERNAME string, PASSWORD string) interface{} {
	var user global.USER
	res := db.connection.Where("USERNAME = ?", USERNAME).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *usersConnection) IsDuplicateUserName(username string) (tx *gorm.DB) {
	var user global.USER
	return db.connection.Where("USERNAME = ?", username).Take(&user)
}

func (db *usersConnection) FindByUserName(username string) global.USER {
	var user global.USER
	db.connection.Where("USERNAME = ?", username).Take(&user)
	return user
}

func (db *usersConnection) ProfileUser(userID string) global.USER {
	var user global.USER
	db.connection.Preload("USERS").Preload("USERS.USERNAME").Find(&user, userID)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
