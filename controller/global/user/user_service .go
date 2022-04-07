package user

import (
	"log"

	"backend/models/global"

	"github.com/mashingan/smapping"
)

//UserService is a contract.....
type UserService interface {
	Update(user global.USER) global.USER
	Profile(userID string) global.USER
}

type userService struct {
	userRepository UsersRepository
}

//NewUserService creates a new instance of UserService
func NewUserService(userRepo UsersRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user global.USER) global.USER {
	userToUpdate := global.USER{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) global.USER {
	return service.userRepository.ProfileUser(userID)
}
