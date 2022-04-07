package item

import (
	"net/http"
	"strconv"

	"backend/controller/global/user"
	"backend/helper"
	"backend/models/inventory"

	"github.com/gin-gonic/gin"
)

//ItemController interface is a contract what this controller can do
type ItemController interface {
	InsertItem(context *gin.Context)
	ReadAllItem(context *gin.Context)
	FindItemByID(context *gin.Context)
	UpdateItem(context *gin.Context)
	DeleteItem(context *gin.Context)
}

type ITEM_Controller struct {
	Item_Service ItemService
	jwtService   user.JWTService
}

//NewGreigeController create a new instances of BoookController
func NewItemController(ItemServ ItemService, jwtServ user.JWTService) ItemController {
	return &ITEM_Controller{
		Item_Service: ItemServ,
		jwtService:   jwtServ,
	}
}

func (c *ITEM_Controller) InsertItem(context *gin.Context) {
	var ItemCreate inventory.ITEM
	errDTO := context.ShouldBind(&ItemCreate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	result := c.Item_Service.InsertItem(ItemCreate)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)

}

func (c *ITEM_Controller) ReadAllItem(context *gin.Context) {
	var greiges []inventory.ITEM = c.Item_Service.ReadAllItem()
	res := helper.BuildResponse(true, "OK", greiges)
	context.JSON(http.StatusOK, res)
}

func (c *ITEM_Controller) FindItemByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var greige inventory.ITEM = c.Item_Service.FindItemByID(id)
	if (greige == inventory.ITEM{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", greige)
		context.JSON(http.StatusOK, res)
	}
}

func (c *ITEM_Controller) UpdateItem(context *gin.Context) {
	var greigeUpdateDTO inventory.ITEM

	errDTO := context.ShouldBind(&greigeUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	result := c.Item_Service.UpdateItem(greigeUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

func (c *ITEM_Controller) DeleteItem(context *gin.Context) {
	var greige inventory.ITEM
	c.Item_Service.DeleteItem(greige)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
