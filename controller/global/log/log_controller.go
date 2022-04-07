package log

import (
	"fmt"
	"net/http"
	"strconv"

	"backend/controller/global/user"
	"backend/helper"
	"backend/models/global"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//ModuleController interface is a contract what this controller can do
type LogController interface {
	GetAllLog(context *gin.Context)
	FindLogByID(context *gin.Context)
	InsertLog(context *gin.Context)
	UpdateLog(context *gin.Context)
	DeleteLog(context *gin.Context)
}

type logController struct {
	logService LogService
	jwtService user.JWTService
}

//NewModuleController create a new instances of ModuleController
func NewLogController(logServ LogService, jwtServ user.JWTService) LogController {
	return &logController{
		logService: logServ,
		jwtService: jwtServ,
	}
}

// CreateLog
// @Security bearerAuth
// @Description API untuk membuat log baru.
// @Summary Membuat log baru.
// @Tags Log
// @Accept json
// @Produce json
// @Param log body global.LOG true "Log Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /log [post]
func (c *logController) InsertLog(context *gin.Context) {
	var modulCreateDTO global.LOG
	errDTO := context.ShouldBind(&modulCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			modulCreateDTO.CREATE_USER = convertedUserID
			modulCreateDTO.UPDATE_USER = convertedUserID
		}
		result := c.logService.InsertLog(modulCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

// GetModule Get All Log
// @Security bearerAuth
// @Description API untuk mengambil semua log yang terdapat dalam database.
// @Summary Mengambil Semua Log
// @Tags Log
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /log [get]
func (c *logController) GetAllLog(context *gin.Context) {
	var log []global.LOG = c.logService.GetAllLog()
	res := helper.BuildResponse(true, "OK", log)
	context.JSON(http.StatusOK, res)
}

// GetModule by ID Log
// @Security bearerAuth
// @Description API untuk mencari log by ID yang terdapat dalam database.
// @Summary Mengambil Log by ID
// @Tags Log
// @Accept json
// @Produce json
// @Param id path string true "Log ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /log/{id} [get]
func (c *logController) FindLogByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var log global.LOG = c.logService.FindLogByID(id)
	if (log == global.LOG{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", log)
		context.JSON(http.StatusOK, res)
	}
}

// UpdateModule
// @Security bearerAuth
// @Description API untuk update log.
// @Summary Update log.
// @Tags Log
// @Accept json
// @Produce json
// @Param log body global.LOG true "Log Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /log/{id} [put]
func (c *logController) UpdateLog(context *gin.Context) {
	var logUpdateDTO global.LOG
	id, errDTO := strconv.ParseUint(context.Param("id"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("No param id was found", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	errDTO = context.ShouldBind(&logUpdateDTO)
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
	logUpdateDTO.USER_LOGS_ID = id
	logUpdateDTO.UPDATE_USER = idUser
	if errID == nil {
		response := helper.BuildErrorResponse("User Id Not Found", "User Id not found", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
	result := c.logService.UpdateLog(logUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

// DeleteModuleId
// @Security bearerAuth
// @Description API untuk delete log.
// @Summary Delete log.
// @Tags Log
// @Accept json
// @Produce json
// @Param id path string true "Module ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /log/{id} [delete]
func (c *logController) DeleteLog(context *gin.Context) {
	var log global.LOG
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	log.USER_LOGS_ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.logService.IsAllowedToEdit(userID, log.USER_LOGS_ID) {
		c.logService.DeleteLog(log)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *logController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
