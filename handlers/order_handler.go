package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/laertkokona/crud-test/models"
	"github.com/laertkokona/crud-test/services"
	"net/http"
	"strconv"
)

// OrderHandler is the handler for the order resource
type OrderHandler interface {
	CreateOrder(ctx *gin.Context)
	GetOrder(ctx *gin.Context)
	GetAllOrders(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
	DeleteOrder(ctx *gin.Context)
}

// orderHandler is the handler for the order resource
type orderHandler struct {
	orderService services.OrderService
}

// NewOrderHandler returns a new instance of orderHandler
func NewOrderHandler(orderService services.OrderService) OrderHandler {
	return orderHandler{
		orderService: orderService,
	}
}

// CreateOrder method that takes a models.Order object and saves it to the database
func (p orderHandler) CreateOrder(ctx *gin.Context) {
	// get the order object from the request body
	// call the order service to save the order
	// return the order object
	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order, status, err := p.orderService.CreateOrder(order)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, order)
}

// GetOrder method that takes an order id and returns the order object
func (p orderHandler) GetOrder(ctx *gin.Context) {
	// get the order id from the request params
	// call the order service to get the order
	// return the order object
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order, status, err := p.orderService.GetOrder(intId)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, order)
}

// GetAllOrders method that returns all the orders
func (p orderHandler) GetAllOrders(ctx *gin.Context) {
	// call the order service to get all the orders
	// return the orders
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
	orders, status, err := p.orderService.GetAllOrders(pagination)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, orders)
}

// UpdateOrder method that takes an order id and a models.Order object and updates the order
func (p orderHandler) UpdateOrder(ctx *gin.Context) {
	// get the order id from the request params
	// get the order object from the request body
	// call the order service to update the order
	// return the order object
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order, status, err := p.orderService.UpdateOrder(intId, order)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, order)
}

// DeleteOrder method that takes an order id and deletes the order
func (p orderHandler) DeleteOrder(ctx *gin.Context) {
	// get the order id from the request params
	// call the order service to delete the order
	// return the order object
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order, status, err := p.orderService.DeleteOrder(intId)
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(status, order)
}
