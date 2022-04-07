package oborhdr

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

//OborHDRController interface is a contract what this controller can do
type OborHDRController interface {
	InsertOborHDR(context *gin.Context)
	ReadAllOborHDR(context *gin.Context)
	FindOborHDRByID(context *gin.Context)
	UpdateOborHDR(context *gin.Context)
	DeleteOborHDR(context *gin.Context)
}

type OBORHDR_Controller struct {
	oborHDR_Service OborHDRService
	jwtService      user.JWTService
}

//NewGreigeController create a new instances of BoookController
func NewOborHDRController(OborHDRServ OborHDRService, jwtServ user.JWTService) OborHDRController {
	return &OBORHDR_Controller{
		oborHDR_Service: OborHDRServ,
		jwtService:      jwtServ,
	}
}

// CreateOborHDRs
// @Security bearerAuth
// @Description API untuk membuat oborHDRbaru.
// @Summary Membuat oborHDRbaru.
// @Tags OborHDRs
// @Accept json
// @Produce json
// @Param oborHDR body global.MODULE true "OborHDRs Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /oborHDR [post]
func (c *OBORHDR_Controller) InsertOborHDR(context *gin.Context) {
	var oborHDRCreate inventory.OBORHDR
	errDTO := context.ShouldBind(&oborHDRCreate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			oborHDRCreate.CREATE_USER = convertedUserID
			oborHDRCreate.UPDATE_USER = convertedUserID
		}
		result := c.oborHDR_Service.InsertOborHDR(oborHDRCreate)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}

}

// GetOborHDR Get All OborHDRs
// @Security bearerAuth
// @Description API untuk mengambil semua oborHDRyang terdapat dalam database.
// @Summary Mengambil Semua OborHDRs
// @Tags OborHDRs
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /oborHDR [get]
func (c *OBORHDR_Controller) ReadAllOborHDR(context *gin.Context) {
	var greiges []inventory.OBORHDR = c.oborHDR_Service.ReadAllOborHDR()
	res := helper.BuildResponse(true, "OK", greiges)
	context.JSON(http.StatusOK, res)
}

// GetOborHDR by ID OborHDRs
// @Security bearerAuth
// @Description API untuk mencari oborHDRby ID yang terdapat dalam database.
// @Summary Mengambil OborHDRs by ID
// @Tags OborHDRs
// @Accept json
// @Produce json
// @Param id path string true "OborHDRs ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /oborHDR/{id} [get]
func (c *OBORHDR_Controller) FindOborHDRByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var greige inventory.OBORHDR = c.oborHDR_Service.FindOborHDRByID(id)
	if (greige == inventory.OBORHDR{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", greige)
		context.JSON(http.StatusOK, res)
	}
}

// UpdateOborHDR
// @Security bearerAuth
// @Description API untuk update oborHDR
// @Summary Update oborHDR
// @Tags OborHDRs
// @Accept json
// @Produce json
// @Param oborHDR body global.MODULE true "OborHDRs Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /oborHDR/{id} [put]
func (c *OBORHDR_Controller) UpdateOborHDR(context *gin.Context) {
	var oborHDRUpdateDTO inventory.OBORHDR
	id, errDTO := strconv.ParseUint(context.Param("id"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("No param id was found", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	errDTO = context.ShouldBind(&oborHDRUpdateDTO)
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
	oborHDRUpdateDTO.ORDER_ID = id
	oborHDRUpdateDTO.UPDATE_USER = idUser
	if errID == nil {
		response := helper.BuildErrorResponse("User Id Not Found", "User Id not found", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
	result := c.oborHDR_Service.UpdateOborHDR(oborHDRUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

// DeleteOborHDRId
// @Security bearerAuth
// @Description API untuk delete oborHDR by ID
// @Summary Delete oborHDR
// @Tags OborHDRs
// @Accept json
// @Produce json
// @Param id path string true "OborHDR ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /oborHDR/{id} [delete]
func (c *OBORHDR_Controller) DeleteOborHDR(context *gin.Context) {
	var oborHDRDeleteDTO inventory.OBORHDR

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	oborHDRDeleteDTO.ORDER_ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.oborHDR_Service.IsAllowedToEdit(userID, oborHDRDeleteDTO.ORDER_ID) {
		c.oborHDR_Service.DeleteOborHDR(oborHDRDeleteDTO)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *OBORHDR_Controller) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
