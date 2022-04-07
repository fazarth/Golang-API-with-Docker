package partner

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
type PartnerController interface {
	GetAllPartner(context *gin.Context)
	FindPartnerByID(context *gin.Context)
	InsertPartner(context *gin.Context)
	UpdatePartner(context *gin.Context)
	DeletePartner(context *gin.Context)
}

type partnerController struct {
	partnerService PartnerService
	jwtService     user.JWTService
}

//NewModuleController create a new instances of ModuleController
func NewPartnerController(partnerServ PartnerService, jwtServ user.JWTService) PartnerController {
	return &partnerController{
		partnerService: partnerServ,
		jwtService:     jwtServ,
	}
}

// CreatePartner
// @Security bearerAuth
// @Description API untuk membuat partner baru.
// @Summary Membuat partner baru.
// @Tags Partner
// @Accept json
// @Produce json
// @Param partner body global.PARTNER true "Partner Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /partner [post]
func (c *partnerController) InsertPartner(context *gin.Context) {
	var partnerCreateDTO global.PARTNER
	errDTO := context.ShouldBind(&partnerCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			partnerCreateDTO.CREATE_USER = convertedUserID
			partnerCreateDTO.UPDATE_USER = convertedUserID
		}
		result := c.partnerService.InsertPartner(partnerCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

// GetModule Get All Partner
// @Security bearerAuth
// @Description API untuk mengambil semua partner yang terdapat dalam database.
// @Summary Mengambil Semua Partner
// @Tags Partner
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /partner [get]
func (c *partnerController) GetAllPartner(context *gin.Context) {
	var partner []global.PARTNER = c.partnerService.GetAllPartner()
	res := helper.BuildResponse(true, "OK", partner)
	context.JSON(http.StatusOK, res)
}

// GetModule by ID Partner
// @Security bearerAuth
// @Description API untuk mencari partner by ID yang terdapat dalam database.
// @Summary Mengambil Partner by ID
// @Tags Partner
// @Accept json
// @Produce json
// @Param id path string true "Partner ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /partner/{id} [get]
func (c *partnerController) FindPartnerByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var partner global.PARTNER = c.partnerService.FindPartnerByID(id)
	if (partner == global.PARTNER{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", partner)
		context.JSON(http.StatusOK, res)
	}
}

// UpdateModule
// @Security bearerAuth
// @Description API untuk update partner.
// @Summary Update partner.
// @Tags Partner
// @Accept json
// @Produce json
// @Param partner body global.PARTNER true "Partner Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /partner/{id} [put]
func (c *partnerController) UpdatePartner(context *gin.Context) {
	var partnerUpdateDTO global.PARTNER
	id, errDTO := strconv.ParseUint(context.Param("id"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("No param id was found", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	errDTO = context.ShouldBind(&partnerUpdateDTO)
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
	partnerUpdateDTO.PARTNER_ID = id
	partnerUpdateDTO.UPDATE_USER = idUser
	if errID == nil {
		response := helper.BuildErrorResponse("User Id Not Found", "User Id not found", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
	result := c.partnerService.UpdatePartner(partnerUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

// DeleteModuleId
// @Security bearerAuth
// @Description API untuk delete partner.
// @Summary Delete partner.
// @Tags Partner
// @Accept json
// @Produce json
// @Param id path string true "Module ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /partner/{id} [delete]
func (c *partnerController) DeletePartner(context *gin.Context) {
	var partner global.PARTNER
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	partner.PARTNER_ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.partnerService.IsAllowedToEdit(userID, partner.PARTNER_ID) {
		c.partnerService.DeletePartner(partner)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *partnerController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
