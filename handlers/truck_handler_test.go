package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/laertkokona/crud-test/models"
	"github.com/peteprogrammer/go-automapper"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mockTruckModels = []gorm.Model{
	{
		ID: uint(1),
	},
	{
		ID: uint(2),
	},
	{
		ID: uint(3),
	},
}

var mockTruck = models.Truck{
	ChassisNumber: "AA444",
	LicensePlate:  "AA444AA",
}

var mockTrucks = []models.Truck{
	{
		Model:         mockTruckModels[0],
		ChassisNumber: "AA111",
		LicensePlate:  "AA111AA",
	},
	{
		Model:         mockTruckModels[1],
		ChassisNumber: "AA222",
		LicensePlate:  "AA222AA",
	},
	{
		Model:         mockTruckModels[2],
		ChassisNumber: "AA333",
		LicensePlate:  "AA333AA",
	},
}

// mockTruckService is a mock implementation of the TruckService interface
type mockTruckService struct {
	createTruck  func(truck models.Truck) (models.TruckDTO, int, error)
	getTruck     func(id int) (models.TruckDTO, int, error)
	getAllTrucks func(pagination models.Pagination) ([]models.TruckDTO, int, error)
	updateTruck  func(id int, truck models.Truck) (models.TruckDTO, int, error)
	deleteTruck  func(id int) (models.TruckDTO, int, error)
}

// CreateTruck is a mock implementation of the CreateTruck method
func (m mockTruckService) CreateTruck(truck models.Truck) (models.TruckDTO, int, error) {
	return m.createTruck(truck)
}

// GetTruck is a mock implementation of the GetTruck method
func (m mockTruckService) GetTruck(id int) (models.TruckDTO, int, error) {
	return m.getTruck(id)
}

// GetAllTrucks is a mock implementation of the GetAllTrucks method
func (m mockTruckService) GetAllTrucks(pagination models.Pagination) ([]models.TruckDTO, int, error) {
	return m.getAllTrucks(pagination)
}

// UpdateTruck is a mock implementation of the UpdateTruck method
func (m mockTruckService) UpdateTruck(id int, truck models.Truck) (models.TruckDTO, int, error) {
	return m.updateTruck(id, truck)
}

// DeleteTruck is a mock implementation of the DeleteTruck method
func (m mockTruckService) DeleteTruck(id int) (models.TruckDTO, int, error) {
	return m.deleteTruck(id)
}

// newMockTruckService returns a new instance of mockTruckService
func newMockTruckService() *mockTruckService {
	return &mockTruckService{
		createTruck: func(truck models.Truck) (models.TruckDTO, int, error) {
			var truckDTO models.TruckDTO
			automapper.Map(truck, &truckDTO)
			return truckDTO, http.StatusOK, nil
		},
		getTruck: func(id int) (models.TruckDTO, int, error) {
			var truckDTO models.TruckDTO
			automapper.Map(mockTrucks[id-1], &truckDTO)
			return truckDTO, http.StatusOK, nil
		},
		getAllTrucks: func(pagination models.Pagination) ([]models.TruckDTO, int, error) {
			var trucksDTO []models.TruckDTO
			automapper.Map(mockTrucks, &trucksDTO)
			return trucksDTO, http.StatusOK, nil
		},
		updateTruck: func(id int, truck models.Truck) (models.TruckDTO, int, error) {
			var truckDTO models.TruckDTO
			automapper.Map(truck, &truckDTO)
			return truckDTO, http.StatusOK, nil
		},
		deleteTruck: func(id int) (models.TruckDTO, int, error) {
			var truckDTO models.TruckDTO
			automapper.Map(mockTrucks[id-1], &truckDTO)
			return truckDTO, http.StatusOK, nil
		},
	}
}

// newMockTruckErrorService returns a new instance of mockTruckService with errors
func newMockTruckErrorService() *mockTruckService {
	return &mockTruckService{
		createTruck: func(truck models.Truck) (models.TruckDTO, int, error) {
			return models.TruckDTO{}, http.StatusInternalServerError, nil
		},
		getTruck: func(id int) (models.TruckDTO, int, error) {
			return models.TruckDTO{}, http.StatusInternalServerError, nil
		},
		getAllTrucks: func(pagination models.Pagination) ([]models.TruckDTO, int, error) {
			return []models.TruckDTO{}, http.StatusInternalServerError, nil
		},
		updateTruck: func(id int, truck models.Truck) (models.TruckDTO, int, error) {
			return models.TruckDTO{}, http.StatusInternalServerError, nil
		},
		deleteTruck: func(id int) (models.TruckDTO, int, error) {
			return models.TruckDTO{}, http.StatusInternalServerError, nil
		},
	}
}

// TestCreateTruck tests the CreateTruck method
func TestCreateTruck(t *testing.T) {
	mockTruckService := newMockTruckService()
	truckHandler := NewTruckHandler(mockTruckService)

	mockTruckString, err := json.Marshal(mockTruck)
	assert.NoError(t, err, "Error marshalling mockTruck")

	r := gin.Default()
	r.POST("/trucks", truckHandler.CreateTruck)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/trucks", bytes.NewBuffer(mockTruckString))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")

}
