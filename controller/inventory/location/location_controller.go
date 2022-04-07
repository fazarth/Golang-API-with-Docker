package location

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

//LocationController interface is a contract what this controller can do
type LocationController interface {
	InsertLocation(context *gin.Context)
	ReadAllLocation(context *gin.Context)
	FindLocationByID(context *gin.Context)
	UpdateLocation(context *gin.Context)
	DeleteLocation(context *gin.Context)
}

type LOCATION_Controller struct {
	location_Service LocationService
	jwtService       user.JWTService
}

//NewGreigeController create a new instances of BoookController
func NewLocationController(LocationServ LocationService, jwtServ user.JWTService) LocationController {
	return &LOCATION_Controller{
		location_Service: LocationServ,
		jwtService:       jwtServ,
	}
}

// CreateLocations
// @Security bearerAuth
// @Description API untuk membuat locationbaru.
// @Summary Membuat locationbaru.
// @Tags Locations
// @Accept json
// @Produce json
// @Param location body global.MODULE true "Locations Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /location [post]
func (c *LOCATION_Controller) InsertLocation(context *gin.Context) {
	var locationCreate inventory.LOCATION
	errDTO := context.ShouldBind(&locationCreate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			locationCreate.CREATE_USER = convertedUserID
			locationCreate.UPDATE_USER = convertedUserID
		}
		result := c.location_Service.InsertLocation(locationCreate)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}

}

// GetLocation Get All Locations
// @Security bearerAuth
// @Description API untuk mengambil semua locationyang terdapat dalam database.
// @Summary Mengambil Semua Locations
// @Tags Locations
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /location [get]
func (c *LOCATION_Controller) ReadAllLocation(context *gin.Context) {
	var greiges []inventory.LOCATION = c.location_Service.ReadAllLocation()
	res := helper.BuildResponse(true, "OK", greiges)
	context.JSON(http.StatusOK, res)
}

// GetLocation by ID Locations
// @Security bearerAuth
// @Description API untuk mencari locationby ID yang terdapat dalam database.
// @Summary Mengambil Locations by ID
// @Tags Locations
// @Accept json
// @Produce json
// @Param id path string true "Locations ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /location/{id} [get]
func (c *LOCATION_Controller) FindLocationByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var greige inventory.LOCATION = c.location_Service.FindLocationByID(id)
	if (greige == inventory.LOCATION{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", greige)
		context.JSON(http.StatusOK, res)
	}
}

// UpdateLocation
// @Security bearerAuth
// @Description API untuk update location
// @Summary Update location
// @Tags Locations
// @Accept json
// @Produce json
// @Param location body global.MODULE true "Locations Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /location/{id} [put]
func (c *LOCATION_Controller) UpdateLocation(context *gin.Context) {
	var locationUpdateDTO inventory.LOCATION
	id, errDTO := strconv.ParseUint(context.Param("id"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("No param id was found", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	errDTO = context.ShouldBind(&locationUpdateDTO)
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
	locationUpdateDTO.LOCN_ID = id
	locationUpdateDTO.UPDATE_USER = idUser
	if errID == nil {
		response := helper.BuildErrorResponse("User Id Not Found", "User Id not found", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
	result := c.location_Service.UpdateLocation(locationUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

// DeleteLocationId
// @Security bearerAuth
// @Description API untuk delete location by ID
// @Summary Delete location
// @Tags Locations
// @Accept json
// @Produce json
// @Param id path string true "Location ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /location/{id} [delete]
func (c *LOCATION_Controller) DeleteLocation(context *gin.Context) {
	var locationDeleteDTO inventory.LOCATION

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	locationDeleteDTO.LOCN_ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.location_Service.IsAllowedToEdit(userID, locationDeleteDTO.LOCN_ID) {
		c.location_Service.DeleteLocation(locationDeleteDTO)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *LOCATION_Controller) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
