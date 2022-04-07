package systemsetting

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
type SystemSettingController interface {
	GetAllSystemSetting(context *gin.Context)
	FindSystemSettingByID(context *gin.Context)
	InsertSystemSetting(context *gin.Context)
	UpdateSystemSetting(context *gin.Context)
	DeleteSystemSetting(context *gin.Context)
}

type systemsettingController struct {
	systemsettingsService SystemSettingService
	jwtService            user.JWTService
}

//NewModuleController create a new instances of ModuleController
func NewSystemSettingController(systemsettingServ SystemSettingService, jwtServ user.JWTService) SystemSettingController {
	return &systemsettingController{
		systemsettingsService: systemsettingServ,
		jwtService:            jwtServ,
	}
}

// CreateSystemSetting
// @Security bearerAuth
// @Description API untuk membuat systemsetting baru.
// @Summary Membuat systemsetting baru.
// @Tags SystemSetting
// @Accept json
// @Produce json
// @Param systemsetting body global.SYSTEMSETTING true "SystemSetting Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /systemsettings [post]
func (c *systemsettingController) InsertSystemSetting(context *gin.Context) {
	var systemsettingCreateDTO global.SYSTEMSETTING
	errDTO := context.ShouldBind(&systemsettingCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			systemsettingCreateDTO.CREATE_USER = convertedUserID
			systemsettingCreateDTO.UPDATE_USER = convertedUserID
		}
		result := c.systemsettingsService.InsertSystemSetting(systemsettingCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

// GetModule Get All SystemSetting
// @Security bearerAuth
// @Description API untuk mengambil semua systemsetting yang terdapat dalam database.
// @Summary Mengambil Semua SystemSetting
// @Tags SystemSetting
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /systemsettings [get]
func (c *systemsettingController) GetAllSystemSetting(context *gin.Context) {
	var systemsettings []global.SYSTEMSETTING = c.systemsettingsService.GetAllSystemSetting()
	res := helper.BuildResponse(true, "OK", systemsettings)
	context.JSON(http.StatusOK, res)
}

// GetModule by ID SystemSetting
// @Security bearerAuth
// @Description API untuk mencari systemsetting by ID yang terdapat dalam database.
// @Summary Mengambil SystemSetting by ID
// @Tags SystemSetting
// @Accept json
// @Produce json
// @Param id path string true "SystemSetting ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /systemsetting/{id} [get]
func (c *systemsettingController) FindSystemSettingByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var systemsetting global.SYSTEMSETTING = c.systemsettingsService.FindSystemSettingByID(id)
	if (systemsetting == global.SYSTEMSETTING{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", systemsetting)
		context.JSON(http.StatusOK, res)
	}
}

// UpdateModule
// @Security bearerAuth
// @Description API untuk update systemsetting.
// @Summary Update systemsetting.
// @Tags SystemSetting
// @Accept json
// @Produce json
// @Param systemsetting body global.SYSTEMSETTING true "SystemSetting Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /systemsetting/{id} [put]
func (c *systemsettingController) UpdateSystemSetting(context *gin.Context) {
	var systemsettingUpdateDTO global.SYSTEMSETTING
	id, errDTO := strconv.ParseUint(context.Param("id"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("No param id was found", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	errDTO = context.ShouldBind(&systemsettingUpdateDTO)
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
	systemsettingUpdateDTO.SYSTEM_SETTING_ID = id
	systemsettingUpdateDTO.UPDATE_USER = idUser
	if errID == nil {
		response := helper.BuildErrorResponse("User Id Not Found", "User Id not found", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
	result := c.systemsettingsService.UpdateSystemSetting(systemsettingUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

// DeleteModuleId
// @Security bearerAuth
// @Description API untuk delete systemsetting.
// @Summary Delete systemsetting.
// @Tags SystemSetting
// @Accept json
// @Produce json
// @Param id path string true "Module ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /systemsetting/{id} [delete]
func (c *systemsettingController) DeleteSystemSetting(context *gin.Context) {
	var systemsetting global.SYSTEMSETTING
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	systemsetting.SYSTEM_SETTING_ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.systemsettingsService.IsAllowedToEdit(userID, systemsetting.SYSTEM_SETTING_ID) {
		c.systemsettingsService.DeleteSystemSetting(systemsetting)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *systemsettingController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
