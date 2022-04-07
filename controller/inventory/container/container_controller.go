package container

import (
	"net/http"
	"strconv"

	"backend/controller/global/user"
	"backend/helper"
	"backend/models/inventory"

	"github.com/gin-gonic/gin"
)

//ContainerController interface is a contract what this controller can do
type ContainerController interface {
	InsertContainer(context *gin.Context)
	ReadAllContainer(context *gin.Context)
	FindContainerByID(context *gin.Context)
	UpdateContainer(context *gin.Context)
	DeleteContainer(context *gin.Context)
}

type CONTAINER_Controller struct {
	Container_Service ContainerService
	jwtService        user.JWTService
}

//NewContainerController create a new instances of ContainerkController
func NewContainerController(ContainerServ ContainerService, jwtServ user.JWTService) ContainerController {
	return &CONTAINER_Controller{
		Container_Service: ContainerServ,
		jwtService:        jwtServ,
	}
}

func (c *CONTAINER_Controller) InsertContainer(context *gin.Context) {
	var ContainerCreate inventory.CONTAINER
	errDTO := context.ShouldBind(&ContainerCreate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	result := c.Container_Service.InsertContainer(ContainerCreate)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)

}

func (c *CONTAINER_Controller) ReadAllContainer(context *gin.Context) {
	var greiges []inventory.CONTAINER = c.Container_Service.ReadAllContainer()
	res := helper.BuildResponse(true, "OK", greiges)
	context.JSON(http.StatusOK, res)
}

func (c *CONTAINER_Controller) FindContainerByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var greige inventory.CONTAINER = c.Container_Service.FindContainerByID(id)
	if (greige == inventory.CONTAINER{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", greige)
		context.JSON(http.StatusOK, res)
	}
}

func (c *CONTAINER_Controller) UpdateContainer(context *gin.Context) {
	var greigeUpdateDTO inventory.CONTAINER

	errDTO := context.ShouldBind(&greigeUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	result := c.Container_Service.UpdateContainer(greigeUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

func (c *CONTAINER_Controller) DeleteContainer(context *gin.Context) {
	var greige inventory.CONTAINER
	c.Container_Service.DeleteContainer(greige)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
