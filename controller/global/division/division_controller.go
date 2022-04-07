package division

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
type DivisionController interface {
	GetAllDivision(context *gin.Context)
	FindDivisionByID(context *gin.Context)
	InsertDivision(context *gin.Context)
	UpdateDivision(context *gin.Context)
	DeleteDivision(context *gin.Context)
}

type divisionController struct {
	divisionsService DivisionService
	jwtService       user.JWTService
}

//NewModuleController create a new instances of ModuleController
func NewDivisionController(divisionServ DivisionService, jwtServ user.JWTService) DivisionController {
	return &divisionController{
		divisionsService: divisionServ,
		jwtService:       jwtServ,
	}
}

// CreateDivision
// @Security bearerAuth
// @Description API untuk membuat division baru.
// @Summary Membuat division baru.
// @Tags Division
// @Accept json
// @Produce json
// @Param division body global.DIVISION true "Division Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /division [post]
func (c *divisionController) InsertDivision(context *gin.Context) {
	var modulCreateDTO global.DIVISION
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
		result := c.divisionsService.InsertDivision(modulCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

// GetModule Get All Division
// @Security bearerAuth
// @Description API untuk mengambil semua division yang terdapat dalam database.
// @Summary Mengambil Semua Division
// @Tags Division
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /division [get]
func (c *divisionController) GetAllDivision(context *gin.Context) {
	var divisions []global.DIVISION = c.divisionsService.GetAllDivision()
	res := helper.BuildResponse(true, "OK", divisions)
	context.JSON(http.StatusOK, res)
}

// GetModule by ID Division
// @Security bearerAuth
// @Description API untuk mencari division by ID yang terdapat dalam database.
// @Summary Mengambil Division by ID
// @Tags Division
// @Accept json
// @Produce json
// @Param id path string true "Division ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /division/{id} [get]
func (c *divisionController) FindDivisionByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var division global.DIVISION = c.divisionsService.FindDivisionByID(id)
	if (division == global.DIVISION{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", division)
		context.JSON(http.StatusOK, res)
	}
}

// UpdateModule
// @Security bearerAuth
// @Description API untuk update division.
// @Summary Update division.
// @Tags Division
// @Accept json
// @Produce json
// @Param division body global.DIVISION true "Division Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /division/{id} [put]
func (c *divisionController) UpdateDivision(context *gin.Context) {
	var divisionUpdateDTO global.DIVISION
	id, errDTO := strconv.ParseUint(context.Param("id"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("No param id was found", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	errDTO = context.ShouldBind(&divisionUpdateDTO)
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
	divisionUpdateDTO.DIV_ID = id
	divisionUpdateDTO.UPDATE_USER = idUser
	if errID == nil {
		response := helper.BuildErrorResponse("User Id Not Found", "User Id not found", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
	result := c.divisionsService.UpdateDivision(divisionUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

// DeleteModuleId
// @Security bearerAuth
// @Description API untuk delete division.
// @Summary Delete division.
// @Tags Division
// @Accept json
// @Produce json
// @Param id path string true "Module ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /division/{id} [delete]
func (c *divisionController) DeleteDivision(context *gin.Context) {
	var division global.DIVISION
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	division.DIV_ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.divisionsService.IsAllowedToEdit(userID, division.DIV_ID) {
		c.divisionsService.DeleteDivision(division)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *divisionController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
