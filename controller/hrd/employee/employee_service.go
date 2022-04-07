package employee

import (
	"backend/models/hrd"
	"fmt"
	"log"

	"github.com/mashingan/smapping"
)

type EmployeeService interface {
	InsertEmployee(b hrd.EMPLOYEE) hrd.EMPLOYEE
	GetAllEmployee() []hrd.EMPLOYEE
	FindEmployeeByID(cont_id uint64) hrd.EMPLOYEE
	UpdateEmployee(b hrd.EMPLOYEE) hrd.EMPLOYEE
	DeleteEmployee(b hrd.EMPLOYEE)
	IsAllowedToEdit(createUser string, cont_id uint64) bool
}

type CONTAINER_Service struct {
	containerRepository EmployeeRepository
}

//NewEmployeeService .....
func NewEmployeeService(EmployeeRepo EmployeeRepository) EmployeeService {
	return &CONTAINER_Service{
		containerRepository: EmployeeRepo,
	}
}

func (service *CONTAINER_Service) InsertEmployee(b hrd.EMPLOYEE) hrd.EMPLOYEE {
	Employee := hrd.EMPLOYEE{}
	err := smapping.FillStruct(&Employee, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.containerRepository.InsertEmployee(Employee)
	return res
}

func (service *CONTAINER_Service) GetAllEmployee() []hrd.EMPLOYEE {
	return service.containerRepository.GetAllEmployee()
}

func (service *CONTAINER_Service) FindEmployeeByID(cont_id uint64) hrd.EMPLOYEE {
	return service.containerRepository.FindEmployeeByID(cont_id)
}

func (service *CONTAINER_Service) UpdateEmployee(b hrd.EMPLOYEE) hrd.EMPLOYEE {
	res := service.containerRepository.UpdateEmployee(b)
	return res
}

func (service *CONTAINER_Service) DeleteEmployee(b hrd.EMPLOYEE) {
	service.containerRepository.DeleteEmployee(b)
}

func (service *CONTAINER_Service) IsAllowedToEdit(createUser string, cont_id uint64) bool {
	b := service.containerRepository.FindEmployeeByID(cont_id)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return createUser == id
}
