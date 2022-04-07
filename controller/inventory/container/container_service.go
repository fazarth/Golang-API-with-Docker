package container

import (
	"backend/models/inventory"
	"fmt"
	"log"

	"github.com/mashingan/smapping"
)

type ContainerService interface {
	InsertContainer(b inventory.CONTAINER) inventory.CONTAINER
	ReadAllContainer() []inventory.CONTAINER
	FindContainerByID(cont_id uint64) inventory.CONTAINER
	UpdateContainer(b inventory.CONTAINER) inventory.CONTAINER
	DeleteContainer(b inventory.CONTAINER)
	IsAllowedToEdit(createUser string, cont_id uint64) bool
}

type CONTAINER_Service struct {
	containerRepository ContainerRepository
}

//NewContainerService .....
func NewContainerService(ContainerRepo ContainerRepository) ContainerService {
	return &CONTAINER_Service{
		containerRepository: ContainerRepo,
	}
}

func (service *CONTAINER_Service) InsertContainer(b inventory.CONTAINER) inventory.CONTAINER {
	Container := inventory.CONTAINER{}
	err := smapping.FillStruct(&Container, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.containerRepository.InsertContainer(Container)
	return res
}

func (service *CONTAINER_Service) ReadAllContainer() []inventory.CONTAINER {
	return service.containerRepository.ReadAllContainer()
}

func (service *CONTAINER_Service) FindContainerByID(cont_id uint64) inventory.CONTAINER {
	return service.containerRepository.FindContainerByID(cont_id)
}

func (service *CONTAINER_Service) UpdateContainer(b inventory.CONTAINER) inventory.CONTAINER {
	res := service.containerRepository.UpdateContainer(b)
	return res
}

func (service *CONTAINER_Service) DeleteContainer(b inventory.CONTAINER) {
	service.containerRepository.DeleteContainer(b)
}

func (service *CONTAINER_Service) IsAllowedToEdit(createUser string, cont_id uint64) bool {
	b := service.containerRepository.FindContainerByID(cont_id)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return createUser == id
}
