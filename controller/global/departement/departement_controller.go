package departement

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
type DepartementController interface {
	GetAllDepartement(context *gin.Context)
	FindDepartementByID(context *gin.Context)
	InsertDepartement(context *gin.Context)
	UpdateDepartement(context *gin.Context)
	DeleteDepartement(context *gin.Context)
}

type moduleController struct {
	modulesService DepartementService
	jwtService     user.JWTService
}

//NewModuleController create a new instances of ModuleController
func NewDepartementController(moduleServ DepartementService, jwtServ user.JWTService) DepartementController {
	return &moduleController{
		modulesService: moduleServ,
		jwtService:     jwtServ,
	}
}

// CreateDepartement
// @Security bearerAuth
// @Description API untuk membuat module baru.
// @Summary Membuat module baru.
// @Tags Departement
// @Accept json
// @Produce json
// @Param module body global.DEPARTEMENT true "Departement Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /module [post]
func (c *moduleController) InsertDepartement(context *gin.Context) {
	var modulCreateDTO global.DEPARTEMENT
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
		result := c.modulesService.InsertDepartement(modulCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

// GetModule Get All Departement
// @Security bearerAuth
// @Description API untuk mengambil semua module yang terdapat dalam database.
// @Summary Mengambil Semua Departement
// @Tags Departement
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /module [get]
func (c *moduleController) GetAllDepartement(context *gin.Context) {
	var modules []global.DEPARTEMENT = c.modulesService.GetAllDepartement()
	res := helper.BuildResponse(true, "OK", modules)
	context.JSON(http.StatusOK, res)
}

// GetModule by ID Departement
// @Security bearerAuth
// @Description API untuk mencari module by ID yang terdapat dalam database.
// @Summary Mengambil Departement by ID
// @Tags Departement
// @Accept json
// @Produce json
// @Param id path string true "Departement ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /module/{id} [get]
func (c *moduleController) FindDepartementByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var module global.DEPARTEMENT = c.modulesService.FindDepartementByID(id)
	if (module == global.DEPARTEMENT{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", module)
		context.JSON(http.StatusOK, res)
	}
}

// UpdateModule
// @Security bearerAuth
// @Description API untuk update module.
// @Summary Update module.
// @Tags Departement
// @Accept json
// @Produce json
// @Param module body global.DEPARTEMENT true "Departement Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /module/{id} [put]
func (c *moduleController) UpdateDepartement(context *gin.Context) {
	var moduleUpdateDTO global.DEPARTEMENT
	id, errDTO := strconv.ParseUint(context.Param("id"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("No param id was found", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	errDTO = context.ShouldBind(&moduleUpdateDTO)
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
	moduleUpdateDTO.DEPT_ID = id
	moduleUpdateDTO.UPDATE_USER = idUser
	if errID == nil {
		response := helper.BuildErrorResponse("User Id Not Found", "User Id not found", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
	result := c.modulesService.UpdateDepartement(moduleUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

// DeleteModuleId
// @Security bearerAuth
// @Description API untuk delete module.
// @Summary Delete module.
// @Tags Departement
// @Accept json
// @Produce json
// @Param id path string true "Module ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /module/{id} [delete]
func (c *moduleController) DeleteDepartement(context *gin.Context) {
	var module global.DEPARTEMENT
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	module.DEPT_ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.modulesService.IsAllowedToEdit(userID, module.DEPT_ID) {
		c.modulesService.DeleteDepartement(module)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *moduleController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
