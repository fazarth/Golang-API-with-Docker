package permission

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
type PermissionController interface {
	GetAllPermission(context *gin.Context)
	FindPermissionByID(context *gin.Context)
	InsertPermission(context *gin.Context)
	UpdatePermission(context *gin.Context)
	DeletePermission(context *gin.Context)
}

type permissionController struct {
	permissionService PermissionService
	jwtService        user.JWTService
}

//NewModuleController create a new instances of ModuleController
func NewPermissionController(permissionServ PermissionService, jwtServ user.JWTService) PermissionController {
	return &permissionController{
		permissionService: permissionServ,
		jwtService:        jwtServ,
	}
}

// CreatePermission
// @Security bearerAuth
// @Description API untuk membuat permission baru.
// @Summary Membuat permission baru.
// @Tags Permission
// @Accept json
// @Produce json
// @Param permission body global.PERMISSION true "Permission Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /permission [post]
func (c *permissionController) InsertPermission(context *gin.Context) {
	var permissionCreateDTO global.PERMISSION
	errDTO := context.ShouldBind(&permissionCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			permissionCreateDTO.CREATE_USER = convertedUserID
			permissionCreateDTO.UPDATE_USER = convertedUserID
		}
		result := c.permissionService.InsertPermission(permissionCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

// GetModule Get All Permission
// @Security bearerAuth
// @Description API untuk mengambil semua permission yang terdapat dalam database.
// @Summary Mengambil Semua Permission
// @Tags Permission
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /permission [get]
func (c *permissionController) GetAllPermission(context *gin.Context) {
	var permission []global.PERMISSION = c.permissionService.GetAllPermission()
	res := helper.BuildResponse(true, "OK", permission)
	context.JSON(http.StatusOK, res)
}

// GetModule by ID Permission
// @Security bearerAuth
// @Description API untuk mencari permission by ID yang terdapat dalam database.
// @Summary Mengambil Permission by ID
// @Tags Permission
// @Accept json
// @Produce json
// @Param id path string true "Permission ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /permission/{id} [get]
func (c *permissionController) FindPermissionByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var permission global.PERMISSION = c.permissionService.FindPermissionByID(id)
	if (permission == global.PERMISSION{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", permission)
		context.JSON(http.StatusOK, res)
	}
}

// UpdateModule
// @Security bearerAuth
// @Description API untuk update permission.
// @Summary Update permission.
// @Tags Permission
// @Accept json
// @Produce json
// @Param permission body global.PERMISSION true "Permission Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /permission/{id} [put]
func (c *permissionController) UpdatePermission(context *gin.Context) {
	var permissionUpdateDTO global.PERMISSION
	id, errDTO := strconv.ParseUint(context.Param("id"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("No param id was found", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	errDTO = context.ShouldBind(&permissionUpdateDTO)
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
	permissionUpdateDTO.USER_PERMISSION_ID = id
	permissionUpdateDTO.UPDATE_USER = idUser
	if errID == nil {
		response := helper.BuildErrorResponse("User Id Not Found", "User Id not found", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
	result := c.permissionService.UpdatePermission(permissionUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

// DeleteModuleId
// @Security bearerAuth
// @Description API untuk delete permission.
// @Summary Delete permission.
// @Tags Permission
// @Accept json
// @Produce json
// @Param id path string true "Module ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /permission/{id} [delete]
func (c *permissionController) DeletePermission(context *gin.Context) {
	var permission global.PERMISSION
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	permission.USER_PERMISSION_ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.permissionService.IsAllowedToEdit(userID, permission.USER_PERMISSION_ID) {
		c.permissionService.DeletePermission(permission)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *permissionController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
