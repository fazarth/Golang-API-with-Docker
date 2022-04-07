package taskdtl

import (
	"fmt"
	"net/http"
	"strconv"

	"backend/controller/global/user"
	"backend/helper"
	"backend/models/inventory"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//TaskDTLController interface is a contract what this controller can do
type TaskDTLController interface {
	InsertTaskDTL(context *gin.Context)
	ReadAllTaskDTL(context *gin.Context)
	FindTaskDTLByID(context *gin.Context)
	UpdateTaskDTL(context *gin.Context)
	DeleteTaskDTL(context *gin.Context)
}

type TASKDTL_Controller struct {
	taskDTL_Service TaskDTLService
	jwtService      user.JWTService
}

//NewGreigeController create a new instances of BoookController
func NewTaskDTLController(TaskDTLServ TaskDTLService, jwtServ user.JWTService) TaskDTLController {
	return &TASKDTL_Controller{
		taskDTL_Service: TaskDTLServ,
		jwtService:      jwtServ,
	}
}

// CreateTaskDTLs
// @Security bearerAuth
// @Description API untuk membuat taskDTLbaru.
// @Summary Membuat taskDTLbaru.
// @Tags TaskDTLs
// @Accept json
// @Produce json
// @Param taskDTL body global.MODULE true "TaskDTLs Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /taskDTL [post]
func (c *TASKDTL_Controller) InsertTaskDTL(context *gin.Context) {
	var taskDTLCreate inventory.TASKDTL
	errDTO := context.ShouldBind(&taskDTLCreate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			taskDTLCreate.CREATE_USER = convertedUserID
			taskDTLCreate.UPDATE_USER = convertedUserID
		}
		result := c.taskDTL_Service.InsertTaskDTL(taskDTLCreate)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}

}

// GetTaskDTL Get All TaskDTLs
// @Security bearerAuth
// @Description API untuk mengambil semua taskDTLyang terdapat dalam database.
// @Summary Mengambil Semua TaskDTLs
// @Tags TaskDTLs
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /taskDTL [get]
func (c *TASKDTL_Controller) ReadAllTaskDTL(context *gin.Context) {
	var greiges []inventory.TASKDTL = c.taskDTL_Service.ReadAllTaskDTL()
	res := helper.BuildResponse(true, "OK", greiges)
	context.JSON(http.StatusOK, res)
}

// GetTaskDTL by ID TaskDTLs
// @Security bearerAuth
// @Description API untuk mencari taskDTLby ID yang terdapat dalam database.
// @Summary Mengambil TaskDTLs by ID
// @Tags TaskDTLs
// @Accept json
// @Produce json
// @Param id path string true "TaskDTLs ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /taskDTL/{id} [get]
func (c *TASKDTL_Controller) FindTaskDTLByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var greige inventory.TASKDTL = c.taskDTL_Service.FindTaskDTLByID(id)
	if (greige == inventory.TASKDTL{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", greige)
		context.JSON(http.StatusOK, res)
	}
}

// UpdateTaskDTL
// @Security bearerAuth
// @Description API untuk update taskDTL
// @Summary Update taskDTL
// @Tags TaskDTLs
// @Accept json
// @Produce json
// @Param taskDTL body global.MODULE true "TaskDTLs Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /taskDTL/{id} [put]
func (c *TASKDTL_Controller) UpdateTaskDTL(context *gin.Context) {
	var taskDTLUpdateDTO inventory.TASKDTL
	id, errDTO := strconv.ParseUint(context.Param("id"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("No param id was found", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	errDTO = context.ShouldBind(&taskDTLUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	idUser, errID := strconv.ParseUint(userID, 10, 64)
	taskDTLUpdateDTO.TASK_ID = id
	taskDTLUpdateDTO.UPDATE_USER = idUser
	if errID == nil {
		response := helper.BuildErrorResponse("User Id Not Found", "User Id not found", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
	result := c.taskDTL_Service.UpdateTaskDTL(taskDTLUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

// DeleteTaskDTLId
// @Security bearerAuth
// @Description API untuk delete taskDTL by ID
// @Summary Delete taskDTL
// @Tags TaskDTLs
// @Accept json
// @Produce json
// @Param id path string true "TaskDTL ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /taskDTL/{id} [delete]
func (c *TASKDTL_Controller) DeleteTaskDTL(context *gin.Context) {
	var taskDTLDeleteDTO inventory.TASKDTL

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	taskDTLDeleteDTO.TASK_ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.taskDTL_Service.IsAllowedToEdit(userID, taskDTLDeleteDTO.TASK_ID) {
		c.taskDTL_Service.DeleteTaskDTL(taskDTLDeleteDTO)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *TASKDTL_Controller) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
