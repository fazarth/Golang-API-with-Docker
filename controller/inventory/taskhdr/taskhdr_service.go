package taskhdr

import (
	"backend/models/inventory"
	"fmt"
	"log"

	"github.com/mashingan/smapping"
)

type TaskHDRService interface {
	InsertTaskHDR(b inventory.TASKHDR) inventory.TASKHDR
	ReadAllTaskHDR() []inventory.TASKHDR
	FindTaskHDRByID(task_id uint64) inventory.TASKHDR
	UpdateTaskHDR(b inventory.TASKHDR) inventory.TASKHDR
	DeleteTaskHDR(b inventory.TASKHDR)
	IsAllowedToEdit(createUser string, task_id uint64) bool
}

type TASKHDR_Service struct {
	taskHDRRepository TaskHDRRepository
}

//NewTaskHDRService .....
func NewTaskHDRService(TaskHDRRepo TaskHDRRepository) TaskHDRService {
	return &TASKHDR_Service{
		taskHDRRepository: TaskHDRRepo,
	}
}

func (service *TASKHDR_Service) InsertTaskHDR(b inventory.TASKHDR) inventory.TASKHDR {
	TaskHDR := inventory.TASKHDR{}
	err := smapping.FillStruct(&TaskHDR, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.taskHDRRepository.InsertTaskHDR(TaskHDR)
	return res
}

func (service *TASKHDR_Service) ReadAllTaskHDR() []inventory.TASKHDR {
	return service.taskHDRRepository.ReadAllTaskHDR()
}

func (service *TASKHDR_Service) FindTaskHDRByID(task_id uint64) inventory.TASKHDR {
	return service.taskHDRRepository.FindTaskHDRByID(task_id)
}

func (service *TASKHDR_Service) UpdateTaskHDR(b inventory.TASKHDR) inventory.TASKHDR {
	res := service.taskHDRRepository.UpdateTaskHDR(b)
	return res
}

func (service *TASKHDR_Service) DeleteTaskHDR(b inventory.TASKHDR) {
	service.taskHDRRepository.DeleteTaskHDR(b)
}

func (service *TASKHDR_Service) IsAllowedToEdit(createUser string, task_id uint64) bool {
	b := service.taskHDRRepository.FindTaskHDRByID(task_id)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return createUser == id
}
