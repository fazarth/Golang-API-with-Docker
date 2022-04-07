package global

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
type GlobalController interface {
	GetAllGlobal(context *gin.Context)
	FindGlobalByID(context *gin.Context)
	InsertGlobal(context *gin.Context)
	UpdateGlobal(context *gin.Context)
	DeleteGlobal(context *gin.Context)
}

type globalController struct {
	globalService GlobalService
	jwtService    user.JWTService
}

//NewModuleController create a new instances of ModuleController
func NewGlobalController(globalServ GlobalService, jwtServ user.JWTService) GlobalController {
	return &globalController{
		globalService: globalServ,
		jwtService:    jwtServ,
	}
}

// CreateGlobal
// @Security bearerAuth
// @Description API untuk membuat global baru.
// @Summary Membuat global baru.
// @Tags Global
// @Accept json
// @Produce json
// @Param global body global.GLOBAL true "Global Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /global [post]
func (c *globalController) InsertGlobal(context *gin.Context) {
	var globalCreateDTO global.GLOBAL
	errDTO := context.ShouldBind(&globalCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			globalCreateDTO.CREATE_USER = convertedUserID
			globalCreateDTO.UPDATE_USER = convertedUserID
		}
		result := c.globalService.InsertGlobal(globalCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

// GetModule Get All Global
// @Security bearerAuth
// @Description API untuk mengambil semua global yang terdapat dalam database.
// @Summary Mengambil Semua Global
// @Tags Global
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /global [get]
func (c *globalController) GetAllGlobal(context *gin.Context) {
	var global []global.GLOBAL = c.globalService.GetAllGlobal()
	res := helper.BuildResponse(true, "OK", global)
	context.JSON(http.StatusOK, res)
}

// GetModule by ID Global
// @Security bearerAuth
// @Description API untuk mencari global by ID yang terdapat dalam database.
// @Summary Mengambil Global by ID
// @Tags Global
// @Accept json
// @Produce json
// @Param id path string true "Global ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /global/{id} [get]
func (c *globalController) FindGlobalByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var globals global.GLOBAL = c.globalService.FindGlobalByID(id)
	if (globals == global.GLOBAL{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", globals)
		context.JSON(http.StatusOK, res)
	}
}

// UpdateModule
// @Security bearerAuth
// @Description API untuk update global.
// @Summary Update global.
// @Tags Global
// @Accept json
// @Produce json
// @Param global body global.GLOBAL true "Global Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /global/{id} [put]
func (c *globalController) UpdateGlobal(context *gin.Context) {
	var globalUpdateDTO global.GLOBAL
	id, errDTO := strconv.ParseUint(context.Param("id"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("No param id was found", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	errDTO = context.ShouldBind(&globalUpdateDTO)
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
	globalUpdateDTO.GLOBAL_ID = id
	globalUpdateDTO.UPDATE_USER = idUser
	if errID == nil {
		response := helper.BuildErrorResponse("User Id Not Found", "User Id not found", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
	result := c.globalService.UpdateGlobal(globalUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

// DeleteModuleId
// @Security bearerAuth
// @Description API untuk delete global.
// @Summary Delete global.
// @Tags Global
// @Accept json
// @Produce json
// @Param id path string true "Module ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /global/{id} [delete]
func (c *globalController) DeleteGlobal(context *gin.Context) {
	var global global.GLOBAL
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	global.GLOBAL_ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.globalService.IsAllowedToEdit(userID, global.GLOBAL_ID) {
		c.globalService.DeleteGlobal(global)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *globalController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
