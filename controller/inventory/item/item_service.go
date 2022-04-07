package item

import (
	"backend/models/inventory"
	"fmt"
	"log"

	"github.com/mashingan/smapping"
)

type ItemService interface {
	InsertItem(b inventory.ITEM) inventory.ITEM
	ReadAllItem() []inventory.ITEM
	FindItemByID(item_id uint64) inventory.ITEM
	UpdateItem(b inventory.ITEM) inventory.ITEM
	DeleteItem(b inventory.ITEM)
	IsAllowedToEdit(createUser string, item_id uint64) bool
}

type ITEM_Service struct {
	containerRepository ItemRepository
}

//NewItemService .....
func NewItemService(ItemRepo ItemRepository) ItemService {
	return &ITEM_Service{
		containerRepository: ItemRepo,
	}
}

func (service *ITEM_Service) InsertItem(b inventory.ITEM) inventory.ITEM {
	Item := inventory.ITEM{}
	err := smapping.FillStruct(&Item, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.containerRepository.InsertItem(Item)
	return res
}

func (service *ITEM_Service) ReadAllItem() []inventory.ITEM {
	return service.containerRepository.ReadAllItem()
}

func (service *ITEM_Service) FindItemByID(item_id uint64) inventory.ITEM {
	return service.containerRepository.FindItemByID(item_id)
}

func (service *ITEM_Service) UpdateItem(b inventory.ITEM) inventory.ITEM {
	res := service.containerRepository.UpdateItem(b)
	return res
}

func (service *ITEM_Service) DeleteItem(b inventory.ITEM) {
	service.containerRepository.DeleteItem(b)
}

func (service *ITEM_Service) IsAllowedToEdit(createUser string, item_id uint64) bool {
	b := service.containerRepository.FindItemByID(item_id)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return createUser == id
}
