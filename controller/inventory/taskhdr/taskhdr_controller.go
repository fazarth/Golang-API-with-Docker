package taskhdr

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

//TaskHDRController interface is a contract what this controller can do
type TaskHDRController interface {
	InsertTaskHDR(context *gin.Context)
	ReadAllTaskHDR(context *gin.Context)
	FindTaskHDRByID(context *gin.Context)
	UpdateTaskHDR(context *gin.Context)
	DeleteTaskHDR(context *gin.Context)
}

type TASKHDR_Controller struct {
	taskHDR_Service TaskHDRService
	jwtService      user.JWTService
}

//NewGreigeController create a new instances of BoookController
func NewTaskHDRController(TaskHDRServ TaskHDRService, jwtServ user.JWTService) TaskHDRController {
	return &TASKHDR_Controller{
		taskHDR_Service: TaskHDRServ,
		jwtService:      jwtServ,
	}
}

// CreateTaskHDRs
// @Security bearerAuth
// @Description API untuk membuat taskHDRbaru.
// @Summary Membuat taskHDRbaru.
// @Tags TaskHDRs
// @Accept json
// @Produce json
// @Param taskHDR body global.MODULE true "TaskHDRs Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /taskHDR [post]
func (c *TASKHDR_Controller) InsertTaskHDR(context *gin.Context) {
	var taskHDRCreate inventory.TASKHDR
	errDTO := context.ShouldBind(&taskHDRCreate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			taskHDRCreate.CREATE_USER = convertedUserID
			taskHDRCreate.UPDATE_USER = convertedUserID
		}
		result := c.taskHDR_Service.InsertTaskHDR(taskHDRCreate)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}

}

// GetTaskHDR Get All TaskHDRs
// @Security bearerAuth
// @Description API untuk mengambil semua taskHDRyang terdapat dalam database.
// @Summary Mengambil Semua TaskHDRs
// @Tags TaskHDRs
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /taskHDR [get]
func (c *TASKHDR_Controller) ReadAllTaskHDR(context *gin.Context) {
	var greiges []inventory.TASKHDR = c.taskHDR_Service.ReadAllTaskHDR()
	res := helper.BuildResponse(true, "OK", greiges)
	context.JSON(http.StatusOK, res)
}

// GetTaskHDR by ID TaskHDRs
// @Security bearerAuth
// @Description API untuk mencari taskHDRby ID yang terdapat dalam database.
// @Summary Mengambil TaskHDRs by ID
// @Tags TaskHDRs
// @Accept json
// @Produce json
// @Param id path string true "TaskHDRs ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /taskHDR/{id} [get]
func (c *TASKHDR_Controller) FindTaskHDRByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var greige inventory.TASKHDR = c.taskHDR_Service.FindTaskHDRByID(id)
	if (greige == inventory.TASKHDR{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", greige)
		context.JSON(http.StatusOK, res)
	}
}

// UpdateTaskHDR
// @Security bearerAuth
// @Description API untuk update taskHDR
// @Summary Update taskHDR
// @Tags TaskHDRs
// @Accept json
// @Produce json
// @Param taskHDR body global.MODULE true "TaskHDRs Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /taskHDR/{id} [put]
func (c *TASKHDR_Controller) UpdateTaskHDR(context *gin.Context) {
	var taskHDRUpdateDTO inventory.TASKHDR
	id, errDTO := strconv.ParseUint(context.Param("id"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("No param id was found", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	errDTO = context.ShouldBind(&taskHDRUpdateDTO)
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
	taskHDRUpdateDTO.TASK_ID = id
	taskHDRUpdateDTO.UPDATE_USER = idUser
	if errID == nil {
		response := helper.BuildErrorResponse("User Id Not Found", "User Id not found", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
	result := c.taskHDR_Service.UpdateTaskHDR(taskHDRUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

// DeleteTaskHDRId
// @Security bearerAuth
// @Description API untuk delete taskHDR by ID
// @Summary Delete taskHDR
// @Tags TaskHDRs
// @Accept json
// @Produce json
// @Param id path string true "TaskHDR ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /taskHDR/{id} [delete]
func (c *TASKHDR_Controller) DeleteTaskHDR(context *gin.Context) {
	var taskHDRDeleteDTO inventory.TASKHDR

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	taskHDRDeleteDTO.TASK_ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.taskHDR_Service.IsAllowedToEdit(userID, taskHDRDeleteDTO.TASK_ID) {
		c.taskHDR_Service.DeleteTaskHDR(taskHDRDeleteDTO)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *TASKHDR_Controller) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
