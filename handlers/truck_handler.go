package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/laertkokona/crud-test/helpers"
	"github.com/laertkokona/crud-test/models"
	"github.com/laertkokona/crud-test/services"
	"net/http"
	"strconv"
)

// TruckHandler interface
type TruckHandler interface {
	CreateTruck(ctx *gin.Context)
	GetTruck(ctx *gin.Context)
	GetAllTrucks(ctx *gin.Context)
	UpdateTruck(ctx *gin.Context)
	DeleteTruck(ctx *gin.Context)
}

// truckHandler struct
type truckHandler struct {
	truckService services.TruckService
}

// NewTruckHandler returns a new instance of truckHandler
func NewTruckHandler(truckService services.TruckService) TruckHandler {
	return truckHandler{
		truckService: truckService,
	}
}

// CreateTruck method that takes a models.Truck object and saves it to the database
//
// CreateTruck godoc
// @Summary Create a new truck
// @Description create truck
// @Tags trucks
// @Accept  json
// @Produce  json
// @Param truck body models.Truck true "Truck object"
// @Success 200 {object} models.Truck
func (p truckHandler) CreateTruck(ctx *gin.Context) {
	var truck models.Truck
	if err := ctx.ShouldBindJSON(&truck); err != nil {
		helpers.FailedResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	truckDTO, status, err := p.truckService.CreateTruck(truck)
	if err != nil {
		helpers.FailedResponse(ctx, status, err.Error(), nil)
		return
	}
	helpers.SuccessResponse(ctx, truckDTO)
}

// GetTruck method that takes a truck id and returns the truck object
func (p truckHandler) GetTruck(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		helpers.FailedResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	truckDTO, status, err := p.truckService.GetTruck(intId)
	if err != nil {
		helpers.FailedResponse(ctx, status, err.Error(), nil)
		return
	}
	helpers.SuccessResponse(ctx, truckDTO)
}

// GetAllTrucks method that returns all trucks
func (p truckHandler) GetAllTrucks(ctx *gin.Context) {
	var pagination models.Pagination
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		helpers.FailedResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	trucksDTO, status, err := p.truckService.GetAllTrucks(pagination)
	if err != nil {
		helpers.FailedResponse(ctx, status, err.Error(), nil)
		return
	}
	helpers.SuccessResponse(ctx, trucksDTO)
}

// UpdateTruck method that takes a truck id and updates the truck in the database
func (p truckHandler) UpdateTruck(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		helpers.FailedResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	var truck models.Truck
	if err := ctx.ShouldBindJSON(&truck); err != nil {
		helpers.FailedResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	truckDTO, status, err := p.truckService.UpdateTruck(intId, truck)
	if err != nil {
		helpers.FailedResponse(ctx, status, err.Error(), nil)
		return
	}
	helpers.SuccessResponse(ctx, truckDTO)
}

// DeleteTruck method that takes a truck id and deletes the truck from the database
func (p truckHandler) DeleteTruck(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		helpers.FailedResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	truckDTO, status, err := p.truckService.DeleteTruck(intId)
	if err != nil {
		helpers.FailedResponse(ctx, status, err.Error(), nil)
		return
	}
	helpers.SuccessResponse(ctx, truckDTO)
}
