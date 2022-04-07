package obordtl

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

//OborDTLController interface is a contract what this controller can do
type OborDTLController interface {
	InsertOborDTL(context *gin.Context)
	ReadAllOborDTL(context *gin.Context)
	FindOborDTLByID(context *gin.Context)
	UpdateOborDTL(context *gin.Context)
	DeleteOborDTL(context *gin.Context)
}

type OBORDTL_Controller struct {
	oborDTL_Service OborDTLService
	jwtService      user.JWTService
}

//NewGreigeController create a new instances of BoookController
func NewOborDTLController(OborDTLServ OborDTLService, jwtServ user.JWTService) OborDTLController {
	return &OBORDTL_Controller{
		oborDTL_Service: OborDTLServ,
		jwtService:      jwtServ,
	}
}

// CreateOborDTLs
// @Security bearerAuth
// @Description API untuk membuat oborDTLbaru.
// @Summary Membuat oborDTLbaru.
// @Tags OborDTLs
// @Accept json
// @Produce json
// @Param oborDTL body global.MODULE true "OborDTLs Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /oborDTL [post]
func (c *OBORDTL_Controller) InsertOborDTL(context *gin.Context) {
	var oborDTLCreate inventory.OBORDTL
	errDTO := context.ShouldBind(&oborDTLCreate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			oborDTLCreate.CREATE_USER = convertedUserID
			oborDTLCreate.UPDATE_USER = convertedUserID
		}
		result := c.oborDTL_Service.InsertOborDTL(oborDTLCreate)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}

}

// GetOborDTL Get All OborDTLs
// @Security bearerAuth
// @Description API untuk mengambil semua oborDTLyang terdapat dalam database.
// @Summary Mengambil Semua OborDTLs
// @Tags OborDTLs
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /oborDTL [get]
func (c *OBORDTL_Controller) ReadAllOborDTL(context *gin.Context) {
	var greiges []inventory.OBORDTL = c.oborDTL_Service.ReadAllOborDTL()
	res := helper.BuildResponse(true, "OK", greiges)
	context.JSON(http.StatusOK, res)
}

// GetOborDTL by ID OborDTLs
// @Security bearerAuth
// @Description API untuk mencari oborDTLby ID yang terdapat dalam database.
// @Summary Mengambil OborDTLs by ID
// @Tags OborDTLs
// @Accept json
// @Produce json
// @Param id path string true "OborDTLs ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /oborDTL/{id} [get]
func (c *OBORDTL_Controller) FindOborDTLByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var greige inventory.OBORDTL = c.oborDTL_Service.FindOborDTLByID(id)
	if (greige == inventory.OBORDTL{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", greige)
		context.JSON(http.StatusOK, res)
	}
}

// UpdateOborDTL
// @Security bearerAuth
// @Description API untuk update oborDTL
// @Summary Update oborDTL
// @Tags OborDTLs
// @Accept json
// @Produce json
// @Param oborDTL body global.MODULE true "OborDTLs Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /oborDTL/{id} [put]
func (c *OBORDTL_Controller) UpdateOborDTL(context *gin.Context) {
	var oborDTLUpdateDTO inventory.OBORDTL
	id, errDTO := strconv.ParseUint(context.Param("id"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("No param id was found", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	errDTO = context.ShouldBind(&oborDTLUpdateDTO)
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
	oborDTLUpdateDTO.TASK_ID = id
	oborDTLUpdateDTO.UPDATE_USER = idUser
	if errID == nil {
		response := helper.BuildErrorResponse("User Id Not Found", "User Id not found", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
	result := c.oborDTL_Service.UpdateOborDTL(oborDTLUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

// DeleteOborDTLId
// @Security bearerAuth
// @Description API untuk delete oborDTL by ID
// @Summary Delete oborDTL
// @Tags OborDTLs
// @Accept json
// @Produce json
// @Param id path string true "OborDTL ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /oborDTL/{id} [delete]
func (c *OBORDTL_Controller) DeleteOborDTL(context *gin.Context) {
	var oborDTLDeleteDTO inventory.OBORDTL

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	oborDTLDeleteDTO.TASK_ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.oborDTL_Service.IsAllowedToEdit(userID, oborDTLDeleteDTO.TASK_ID) {
		c.oborDTL_Service.DeleteOborDTL(oborDTLDeleteDTO)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *OBORDTL_Controller) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
