package usergroup

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
type UserGroupController interface {
	GetAllUserGroup(context *gin.Context)
	FindUserGroupByID(context *gin.Context)
	InsertUserGroup(context *gin.Context)
	UpdateUserGroup(context *gin.Context)
	DeleteUserGroup(context *gin.Context)
}

type userGroupController struct {
	userGroupService UserGroupService
	jwtService       user.JWTService
}

//NewModuleController create a new instances of ModuleController
func NewUserGroupController(userGroupServ UserGroupService, jwtServ user.JWTService) UserGroupController {
	return &userGroupController{
		userGroupService: userGroupServ,
		jwtService:       jwtServ,
	}
}

// CreateUserGroup
// @Security bearerAuth
// @Description API untuk membuat userGroup baru.
// @Summary Membuat userGroup baru.
// @Tags UserGroup
// @Accept json
// @Produce json
// @Param userGroup body global.USERGROUP true "UserGroup Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /userGroup [post]
func (c *userGroupController) InsertUserGroup(context *gin.Context) {
	var modulCreateDTO global.USERGROUP
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
		result := c.userGroupService.InsertUserGroup(modulCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

// GetModule Get All UserGroup
// @Security bearerAuth
// @Description API untuk mengambil semua userGroup yang terdapat dalam database.
// @Summary Mengambil Semua UserGroup
// @Tags UserGroup
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /userGroup [get]
func (c *userGroupController) GetAllUserGroup(context *gin.Context) {
	var userGroup []global.USERGROUP = c.userGroupService.GetAllUserGroup()
	res := helper.BuildResponse(true, "OK", userGroup)
	context.JSON(http.StatusOK, res)
}

// GetModule by ID UserGroup
// @Security bearerAuth
// @Description API untuk mencari userGroup by ID yang terdapat dalam database.
// @Summary Mengambil UserGroup by ID
// @Tags UserGroup
// @Accept json
// @Produce json
// @Param id path string true "UserGroup ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /userGroup/{id} [get]
func (c *userGroupController) FindUserGroupByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var userGroup global.USERGROUP = c.userGroupService.FindUserGroupByID(id)
	if (userGroup == global.USERGROUP{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", userGroup)
		context.JSON(http.StatusOK, res)
	}
}

// UpdateModule
// @Security bearerAuth
// @Description API untuk update userGroup.
// @Summary Update userGroup.
// @Tags UserGroup
// @Accept json
// @Produce json
// @Param userGroup body global.USERGROUP true "UserGroup Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /userGroup/{id} [put]
func (c *userGroupController) UpdateUserGroup(context *gin.Context) {
	var userGroupUpdateDTO global.USERGROUP
	id, errDTO := strconv.ParseUint(context.Param("id"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("No param id was found", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	errDTO = context.ShouldBind(&userGroupUpdateDTO)
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
	userGroupUpdateDTO.USER_GROUP_ID = id
	userGroupUpdateDTO.UPDATE_USER = idUser
	if errID == nil {
		response := helper.BuildErrorResponse("User Id Not Found", "User Id not found", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
	result := c.userGroupService.UpdateUserGroup(userGroupUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

// DeleteModuleId
// @Security bearerAuth
// @Description API untuk delete userGroup.
// @Summary Delete userGroup.
// @Tags UserGroup
// @Accept json
// @Produce json
// @Param id path string true "Module ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /userGroup/{id} [delete]
func (c *userGroupController) DeleteUserGroup(context *gin.Context) {
	var userGroup global.USERGROUP
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	userGroup.USER_GROUP_ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.userGroupService.IsAllowedToEdit(userID, userGroup.USER_GROUP_ID) {
		c.userGroupService.DeleteUserGroup(userGroup)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *userGroupController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
