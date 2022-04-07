package company

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
type CompanyController interface {
	GetAllCompany(context *gin.Context)
	FindCompanyByID(context *gin.Context)
	InsertCompany(context *gin.Context)
	UpdateCompany(context *gin.Context)
	DeleteCompany(context *gin.Context)
}

type companyController struct {
	companysService CompanyService
	jwtService      user.JWTService
}

//NewModuleController create a new instances of ModuleController
func NewCompanyController(companyServ CompanyService, jwtServ user.JWTService) CompanyController {
	return &companyController{
		companysService: companyServ,
		jwtService:      jwtServ,
	}
}

// CreateCompany
// @Security bearerAuth
// @Description API untuk membuat company baru.
// @Summary Membuat company baru.
// @Tags Company
// @Accept json
// @Produce json
// @Param company body global.COMPANY true "Company Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /company [post]
func (c *companyController) InsertCompany(context *gin.Context) {
	var companyCreateDTO global.COMPANY
	errDTO := context.ShouldBind(&companyCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			companyCreateDTO.CREATE_USER = convertedUserID
			companyCreateDTO.UPDATE_USER = convertedUserID
		}
		result := c.companysService.InsertCompany(companyCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

// GetModule Get All Company
// @Security bearerAuth
// @Description API untuk mengambil semua company yang terdapat dalam database.
// @Summary Mengambil Semua Company
// @Tags Company
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /company [get]
func (c *companyController) GetAllCompany(context *gin.Context) {
	var companys []global.COMPANY = c.companysService.GetAllCompany()
	res := helper.BuildResponse(true, "OK", companys)
	context.JSON(http.StatusOK, res)
}

// GetModule by ID Company
// @Security bearerAuth
// @Description API untuk mencari company by ID yang terdapat dalam database.
// @Summary Mengambil Company by ID
// @Tags Company
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /company/{id} [get]
func (c *companyController) FindCompanyByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var company global.COMPANY = c.companysService.FindCompanyByID(id)
	if (company == global.COMPANY{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", company)
		context.JSON(http.StatusOK, res)
	}
}

// UpdateModule
// @Security bearerAuth
// @Description API untuk update company.
// @Summary Update company.
// @Tags Company
// @Accept json
// @Produce json
// @Param company body global.COMPANY true "Company Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /company/{id} [put]
func (c *companyController) UpdateCompany(context *gin.Context) {
	var companyUpdateDTO global.COMPANY
	id, errDTO := strconv.ParseUint(context.Param("id"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("No param id was found", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	errDTO = context.ShouldBind(&companyUpdateDTO)
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
	companyUpdateDTO.COMPANY_ID = id
	companyUpdateDTO.UPDATE_USER = idUser
	if errID == nil {
		response := helper.BuildErrorResponse("User Id Not Found", "User Id not found", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
	result := c.companysService.UpdateCompany(companyUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

// DeleteModuleId
// @Security bearerAuth
// @Description API untuk delete company.
// @Summary Delete company.
// @Tags Company
// @Accept json
// @Produce json
// @Param id path string true "Module ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /company/{id} [delete]
func (c *companyController) DeleteCompany(context *gin.Context) {
	var company global.COMPANY
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	company.COMPANY_ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.companysService.IsAllowedToEdit(userID, company.COMPANY_ID) {
		c.companysService.DeleteCompany(company)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *companyController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
