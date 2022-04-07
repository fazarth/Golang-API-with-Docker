package location

import (
	"backend/models/inventory"
	"fmt"
	"log"

	"github.com/mashingan/smapping"
)

type LocationService interface {
	InsertLocation(b inventory.LOCATION) inventory.LOCATION
	ReadAllLocation() []inventory.LOCATION
	FindLocationByID(location_ID uint64) inventory.LOCATION
	UpdateLocation(b inventory.LOCATION) inventory.LOCATION
	DeleteLocation(b inventory.LOCATION)
	IsAllowedToEdit(createUser string, location_ID uint64) bool
}

type LOCATION_Service struct {
	locationRepository LocationRepository
}

//NewLocationService .....
func NewLocationService(LocationRepo LocationRepository) LocationService {
	return &LOCATION_Service{
		locationRepository: LocationRepo,
	}
}

func (service *LOCATION_Service) InsertLocation(b inventory.LOCATION) inventory.LOCATION {
	Location := inventory.LOCATION{}
	err := smapping.FillStruct(&Location, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.locationRepository.InsertLocation(Location)
	return res
}

func (service *LOCATION_Service) ReadAllLocation() []inventory.LOCATION {
	return service.locationRepository.ReadAllLocation()
}

func (service *LOCATION_Service) FindLocationByID(location_ID uint64) inventory.LOCATION {
	return service.locationRepository.FindLocationByID(location_ID)
}

func (service *LOCATION_Service) UpdateLocation(b inventory.LOCATION) inventory.LOCATION {
	res := service.locationRepository.UpdateLocation(b)
	return res
}

func (service *LOCATION_Service) DeleteLocation(b inventory.LOCATION) {
	service.locationRepository.DeleteLocation(b)
}

func (service *LOCATION_Service) IsAllowedToEdit(createUser string, location_ID uint64) bool {
	b := service.locationRepository.FindLocationByID(location_ID)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return createUser == id
}
