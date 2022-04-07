package employee

import (
	"net/http"
	"strconv"

	"backend/controller/global/user"
	"backend/helper"
	"backend/models/hrd"

	"github.com/gin-gonic/gin"
)

//EmployeeController interface is a contract what this controller can do
type EmployeeController interface {
	InsertEmployee(context *gin.Context)
	GetAllEmployee(context *gin.Context)
	FindEmployeeByID(context *gin.Context)
	UpdateEmployee(context *gin.Context)
	DeleteEmployee(context *gin.Context)
}

type EMPLOYEE_Controller struct {
	Employee_Service EmployeeService
	jwtService       user.JWTService
}

func NewEmployeeController(EmployeeServ EmployeeService, jwtServ user.JWTService) EmployeeController {
	return &EMPLOYEE_Controller{
		Employee_Service: EmployeeServ,
		jwtService:       jwtServ,
	}
}

// CreateEmployees
// @Security bearerAuth
// @Description API untuk membuat employee baru.
// @Summary Membuat employee baru.
// @Tags Employees
// @Accept json
// @Produce json
// @Param employee body global.MODULE true "Employees Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /employee [post]
func (c *EMPLOYEE_Controller) InsertEmployee(context *gin.Context) {
	var EmployeeCreate hrd.EMPLOYEE
	errDTO := context.ShouldBind(&EmployeeCreate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	result := c.Employee_Service.InsertEmployee(EmployeeCreate)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)

}

// GetEmployees Get All Employees
// @Security bearerAuth
// @Description API untuk mengambil semua employee yang terdapat dalam database.
// @Summary Mengambil Semua Employees
// @Tags Employees
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /employee [get]
func (c *EMPLOYEE_Controller) GetAllEmployee(context *gin.Context) {
	var greiges []hrd.EMPLOYEE = c.Employee_Service.GetAllEmployee()
	res := helper.BuildResponse(true, "OK", greiges)
	context.JSON(http.StatusOK, res)
}

// GetEmployee by ID Employees
// @Security bearerAuth
// @Description API untuk mencari employee by ID yang terdapat dalam database.
// @Summary Mengambil Employees by ID
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path string true "Employees ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /employee/{id} [get]
func (c *EMPLOYEE_Controller) FindEmployeeByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var greige hrd.EMPLOYEE = c.Employee_Service.FindEmployeeByID(id)
	if (greige == hrd.EMPLOYEE{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", greige)
		context.JSON(http.StatusOK, res)
	}
}

// UpdateEmployee
// @Security bearerAuth
// @Description API untuk update employee.
// @Summary Update employee.
// @Tags Employees
// @Accept json
// @Produce json
// @Param employee body global.MODULE true "Employees Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /employee/{id} [put]
func (c *EMPLOYEE_Controller) UpdateEmployee(context *gin.Context) {
	var greigeUpdateDTO hrd.EMPLOYEE

	errDTO := context.ShouldBind(&greigeUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	result := c.Employee_Service.UpdateEmployee(greigeUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

// DeleteEmployeeId
// @Security bearerAuth
// @Description API untuk delete employee.
// @Summary Delete employee.
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /employee/{id} [delete]
func (c *EMPLOYEE_Controller) DeleteEmployee(context *gin.Context) {
	var greige hrd.EMPLOYEE
	c.Employee_Service.DeleteEmployee(greige)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
