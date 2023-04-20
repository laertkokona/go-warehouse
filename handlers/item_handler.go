package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/laertkokona/crud-test/models"
	"github.com/laertkokona/crud-test/services"
	"net/http"
	"strconv"
)

// ItemHandler interface
type ItemHandler interface {
	CreateItem(ctx *gin.Context)
	GetItem(ctx *gin.Context)
	GetAllItems(ctx *gin.Context)
	UpdateItem(ctx *gin.Context)
	DeleteItem(ctx *gin.Context)
}

// itemHandler struct
type itemHandler struct {
	itemService services.ItemService
}

// NewItemHandler returns a new instance of itemHandler
func NewItemHandler(itemService services.ItemService) ItemHandler {
	return itemHandler{
		itemService: itemService,
	}
}

// CreateItem method that takes a models.Item object and saves it to the database
func (p itemHandler) CreateItem(ctx *gin.Context) {
	// get the item object from the request body
	// call the item service to save the item
	// return the item object
	var item models.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	itemDTO, status, err := p.itemService.CreateItem(item)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, itemDTO)
}

// GetItem method that takes an item id and returns the item object
func (p itemHandler) GetItem(ctx *gin.Context) {
	// get the item id from the request params
	// call the item service to get the item
	// return the item object
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	itemDTO, status, err := p.itemService.GetItem(intId)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, itemDTO)
}

// GetAllItems method that returns all items
func (p itemHandler) GetAllItems(ctx *gin.Context) {
	// call the item service to get all items
	// return the items object
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		intPage = 0
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//return
	}
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		intLimit = 0
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//return
	}
	pagination := models.Pagination{
		Page:  intPage,
		Limit: intLimit,
	}
	itemsDTO, status, err := p.itemService.GetAllItems(pagination)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, itemsDTO)
}

// UpdateItem method that takes an item id and updates the item object
func (p itemHandler) UpdateItem(ctx *gin.Context) {
	// get the item id from the request params
	// get the item object from the request body
	// call the item service to update the item
	// return the item object
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var item models.Item
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	itemDTO, status, err := p.itemService.UpdateItem(intId, item)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, itemDTO)
}

// DeleteItem method that takes an item id and deletes the item object
func (p itemHandler) DeleteItem(ctx *gin.Context) {
	// get the item id from the request params
	// call the item service to delete the item
	// return the item object
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	itemDTO, status, err := p.itemService.DeleteItem(intId)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, itemDTO)
}
