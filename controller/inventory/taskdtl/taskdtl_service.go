package taskdtl

import (
	"backend/models/inventory"
	"fmt"
	"log"

	"github.com/mashingan/smapping"
)

type TaskDTLService interface {
	InsertTaskDTL(b inventory.TASKDTL) inventory.TASKDTL
	ReadAllTaskDTL() []inventory.TASKDTL
	FindTaskDTLByID(task_id uint64) inventory.TASKDTL
	UpdateTaskDTL(b inventory.TASKDTL) inventory.TASKDTL
	DeleteTaskDTL(b inventory.TASKDTL)
	IsAllowedToEdit(createUser string, task_id uint64) bool
}

type TASKDTL_Service struct {
	taskDTLRepository TaskDTLRepository
}

//NewTaskDTLService .....
func NewTaskDTLService(TaskDTLRepo TaskDTLRepository) TaskDTLService {
	return &TASKDTL_Service{
		taskDTLRepository: TaskDTLRepo,
	}
}

func (service *TASKDTL_Service) InsertTaskDTL(b inventory.TASKDTL) inventory.TASKDTL {
	TaskDTL := inventory.TASKDTL{}
	err := smapping.FillStruct(&TaskDTL, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.taskDTLRepository.InsertTaskDTL(TaskDTL)
	return res
}

func (service *TASKDTL_Service) ReadAllTaskDTL() []inventory.TASKDTL {
	return service.taskDTLRepository.ReadAllTaskDTL()
}

func (service *TASKDTL_Service) FindTaskDTLByID(task_id uint64) inventory.TASKDTL {
	return service.taskDTLRepository.FindTaskDTLByID(task_id)
}

func (service *TASKDTL_Service) UpdateTaskDTL(b inventory.TASKDTL) inventory.TASKDTL {
	res := service.taskDTLRepository.UpdateTaskDTL(b)
	return res
}

func (service *TASKDTL_Service) DeleteTaskDTL(b inventory.TASKDTL) {
	service.taskDTLRepository.DeleteTaskDTL(b)
}

func (service *TASKDTL_Service) IsAllowedToEdit(createUser string, task_id uint64) bool {
	b := service.taskDTLRepository.FindTaskDTLByID(task_id)
	id := fmt.Sprintf("%v", b.CREATE_USER)
	return createUser == id
}
