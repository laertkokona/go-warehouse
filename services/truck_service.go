package services

import (
	"github.com/laertkokona/crud-test/models"
	"github.com/laertkokona/crud-test/repositories"
	"github.com/laertkokona/crud-test/utils"
	"github.com/peteprogrammer/go-automapper"
	"net/http"
)

// TruckService interface using gin context
type TruckService interface {
	CreateTruck(truck models.Truck) (models.TruckDTO, int, error)
	GetTruck(id int) (models.TruckDTO, int, error)
	GetAllTrucks(pagination models.Pagination) ([]models.TruckDTO, int, error)
	UpdateTruck(id int, truck models.Truck) (models.TruckDTO, int, error)
	DeleteTruck(id int) (models.TruckDTO, int, error)
}

// truckService struct
type truckService struct {
	TruckRepo repositories.TruckRepo
}

// NewTruckService returns a new instance of truckService
func NewTruckService(truckRepo repositories.TruckRepo) TruckService {
	return truckService{
		TruckRepo: truckRepo,
	}
}

// CreateTruck method that takes a models.Truck object and saves it to the database
func (p truckService) CreateTruck(truck models.Truck) (models.TruckDTO, int, error) {
	truck, err := p.TruckRepo.Save(truck)
	if err != nil {
		return models.TruckDTO{}, http.StatusInternalServerError, err
	}
	var truckDTO models.TruckDTO
	automapper.Map(truck, &truckDTO)
	return truckDTO, http.StatusOK, nil
}

// GetTruck method that takes a truck id and returns the truck object
func (p truckService) GetTruck(id int) (models.TruckDTO, int, error) {
	truck, err := p.TruckRepo.FindByID(id)
	if err != nil {
		return models.TruckDTO{}, http.StatusNotFound, err
	}
	var truckDTO models.TruckDTO
	automapper.Map(truck, &truckDTO)
	return truckDTO, http.StatusOK, nil
}

// GetAllTrucks method that returns all trucks
func (p truckService) GetAllTrucks(pagination models.Pagination) ([]models.TruckDTO, int, error) {
	trucks, err := p.TruckRepo.FindAll(pagination)
	if err != nil {
		return []models.TruckDTO{}, http.StatusInternalServerError, err
	}
	var trucksDTO []models.TruckDTO
	automapper.Map(trucks, &trucksDTO)
	return trucksDTO, http.StatusOK, nil
}

// UpdateTruck method that takes a truck id and updates the truck in the database
func (p truckService) UpdateTruck(id int, truck models.Truck) (models.TruckDTO, int, error) {
	truckDb, err := p.TruckRepo.FindByID(id)
	if err != nil {
		return models.TruckDTO{}, http.StatusNotFound, err
	}
	utils.CopyNonEmptyFields(&truckDb, &truck)
	truckDb, err = p.TruckRepo.Update(truckDb)
	if err != nil {
		return models.TruckDTO{}, http.StatusInternalServerError, err
	}
	var truckDTO models.TruckDTO
	automapper.Map(truckDb, &truckDTO)
	return truckDTO, http.StatusOK, nil
}

// DeleteTruck method that takes a truck id and deletes the truck from the database
func (p truckService) DeleteTruck(id int) (models.TruckDTO, int, error) {
	truck, err := p.TruckRepo.FindByID(id)
	if err != nil {
		return models.TruckDTO{}, http.StatusNotFound, err
	}
	err = p.TruckRepo.Delete(truck)
	if err != nil {
		return models.TruckDTO{}, http.StatusInternalServerError, err
	}
	var truckDTO models.TruckDTO
	automapper.Map(truck, &truckDTO)
	return truckDTO, http.StatusOK, nil
}
