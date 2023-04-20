package services

import (
	"errors"
	"github.com/laertkokona/crud-test/models"
	"github.com/peteprogrammer/go-automapper"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
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

// mockTruckRepo is a mock implementation of the repositories.TruckRepo interface
type mockTruckRepo struct {
	// findAll is a mock function with given fields: pagination
	findAll func(pagination models.Pagination) ([]models.Truck, error)
	// findByID is a mock function with given fields: id
	findByID func(id int) (models.Truck, error)
	// save is a mock function with given fields: truck
	save func(truck models.Truck) (models.Truck, error)
	// update is a mock function with given fields: truck
	update func(truck models.Truck) (models.Truck, error)
	// delete is a mock function with given fields: truck
	delete func(truck models.Truck) error
	// deleteById is a mock function with given fields: id
	deleteById func(id int) (models.Truck, error)
}

// FindAll is a mock function with given fields: pagination
func (_m *mockTruckRepo) FindAll(pagination models.Pagination) ([]models.Truck, error) {
	return _m.findAll(pagination)
}

// FindByID is a mock function with given fields: id
func (_m *mockTruckRepo) FindByID(id int) (models.Truck, error) {
	return _m.findByID(id)
}

// Save is a mock function with given fields: truck
func (_m *mockTruckRepo) Save(truck models.Truck) (models.Truck, error) {
	return _m.save(truck)
}

// Update is a mock function with given fields: truck
func (_m *mockTruckRepo) Update(truck models.Truck) (models.Truck, error) {
	return _m.update(truck)
}

// Delete is a mock function with given fields: truck
func (_m *mockTruckRepo) Delete(truck models.Truck) error {
	return _m.delete(truck)
}

// DeleteById is a mock function with given fields: id
func (_m *mockTruckRepo) DeleteById(id int) (models.Truck, error) {
	return _m.deleteById(id)
}

// newMockTruckRepo returns a new instance of the mockTruckRepo
func newMockTruckRepo() *mockTruckRepo {
	return &mockTruckRepo{
		findAll: func(pagination models.Pagination) ([]models.Truck, error) {
			return mockTrucks, nil
		},
		findByID: func(id int) (models.Truck, error) {
			return mockTrucks[id-1], nil
		},
		save: func(truck models.Truck) (models.Truck, error) {
			return truck, nil
		},
		update: func(truck models.Truck) (models.Truck, error) {
			truck.ID = 1
			return truck, nil
		},
		delete: func(truck models.Truck) error {
			return nil
		},
		deleteById: func(id int) (models.Truck, error) {
			return mockTrucks[id-1], nil
		},
	}
}

// newMockTruckErrorRepo returns a new instance of the mockTruckRepo with error
func newMockTruckErrorRepo() *mockTruckRepo {
	return &mockTruckRepo{
		findAll: func(pagination models.Pagination) ([]models.Truck, error) {
			return nil, errors.New("error")
		},
		findByID: func(id int) (models.Truck, error) {
			return models.Truck{}, errors.New("error")
		},
		save: func(truck models.Truck) (models.Truck, error) {
			return models.Truck{}, errors.New("error")
		},
		update: func(truck models.Truck) (models.Truck, error) {
			return models.Truck{}, errors.New("error")
		},
		delete: func(truck models.Truck) error {
			return errors.New("error")
		},
		deleteById: func(id int) (models.Truck, error) {
			return models.Truck{}, errors.New("error")
		},
	}
}

// newMockTruckSpecificErrorRepo returns a new instance of the mockTruckRepo with error
func newMockTruckSpecificErrorRepo() *mockTruckRepo {
	return &mockTruckRepo{
		findAll: func(pagination models.Pagination) ([]models.Truck, error) {
			return mockTrucks, nil
		},
		findByID: func(id int) (models.Truck, error) {
			return mockTrucks[id-1], nil
		},
		save: func(truck models.Truck) (models.Truck, error) {
			return models.Truck{}, errors.New("error")
		},
		update: func(truck models.Truck) (models.Truck, error) {
			return models.Truck{}, errors.New("error")
		},
		delete: func(truck models.Truck) error {
			return errors.New("error")
		},
		deleteById: func(id int) (models.Truck, error) {
			return models.Truck{}, errors.New("error")
		},
	}
}

// TestNewTruckService tests services.NewTruckService
func TestNewTruckService(t *testing.T) {
	mockRepo := newMockTruckRepo()
	mockService := NewTruckService(mockRepo)

	assert.NotNil(t, mockService)
	assert.IsType(t, truckService{}, mockService)
}

// TestCreateTruck tests services.CreateTruck using mockTruckRepo and gin
func TestCreateTruck(t *testing.T) {
	mockRepo := newMockTruckRepo()
	mockService := NewTruckService(mockRepo)

	mockTruck := models.Truck{
		LicensePlate:  "AA444",
		ChassisNumber: "AA444AA",
	}

	truckDTO, status, err := mockService.CreateTruck(mockTruck)
	var mockTruckDTO models.TruckDTO
	automapper.Map(mockTruck, &mockTruckDTO)
	assert.NoError(t, err, "should not return error")
	assert.Equal(t, http.StatusOK, status, "should return status ok")
	assert.Equal(t, mockTruckDTO, truckDTO, "should return truck")
}

// TestCreateTruck_SaveError tests services.CreateTruck using mockTruckRepo and gin
func TestCreateTruck_SaveError(t *testing.T) {
	mockRepo := newMockTruckErrorRepo()
	mockService := NewTruckService(mockRepo)

	mockTruck := models.Truck{
		LicensePlate:  "AA444",
		ChassisNumber: "AA444AA",
	}

	truckDTO, status, err := mockService.CreateTruck(mockTruck)
	assert.Error(t, err, "should return error")
	assert.Equal(t, http.StatusInternalServerError, status, "should return status internal server error")
	assert.Equal(t, models.TruckDTO{}, truckDTO, "should return empty truck")
}

// TestGetTruck tests services.GetTruck using mockTruckRepo and gin
func TestGetTruck(t *testing.T) {
	mockRepo := newMockTruckRepo()
	mockService := NewTruckService(mockRepo)

	truckDTO, status, err := mockService.GetTruck(1)
	var mockTruckDTO models.TruckDTO
	automapper.Map(&mockTrucks[0], &mockTruckDTO)
	assert.NoError(t, err, "should not return error")
	assert.Equal(t, http.StatusOK, status, "should return status ok")
	assert.Equal(t, mockTruckDTO, truckDTO, "should return truck")
}

// TestGetTruck_FindError tests services.GetTruck using mockTruckRepo and gin
func TestGetTruck_FindError(t *testing.T) {
	mockRepo := newMockTruckErrorRepo()
	mockService := NewTruckService(mockRepo)

	truckDTO, status, err := mockService.GetTruck(1)
	assert.Error(t, err, "should return error")
	assert.Equal(t, http.StatusNotFound, status, "should return status not found error")
	assert.Equal(t, models.TruckDTO{}, truckDTO, "should return empty truck")
}

// TestGetAllTrucks tests services.GetAllTrucks using mockTruckRepo and gin
func TestGetAllTrucks(t *testing.T) {
	mockRepo := newMockTruckRepo()
	mockService := NewTruckService(mockRepo)

	pagination := models.Pagination{
		Page:  1,
		Limit: 10,
	}
	trucksDTO, status, err := mockService.GetAllTrucks(pagination)
	var mockTrucksDTO []models.TruckDTO
	automapper.Map(mockTrucks, &mockTrucksDTO)
	assert.NoError(t, err, "should not return error")
	assert.Equal(t, http.StatusOK, status, "should return status ok")
	assert.Equal(t, mockTrucksDTO, trucksDTO, "should return trucks")
}

// TestGetAllTrucks_FindError tests services.GetAllTrucks using mockTruckRepo and gin
func TestGetAllTrucks_FindError(t *testing.T) {
	mockRepo := newMockTruckErrorRepo()
	mockService := NewTruckService(mockRepo)

	pagination := models.Pagination{
		Page:  1,
		Limit: 10,
	}
	trucksDTO, status, err := mockService.GetAllTrucks(pagination)
	assert.Error(t, err, "should return error")
	assert.Equal(t, http.StatusInternalServerError, status, "should return status not found error")
	assert.Equal(t, []models.TruckDTO{}, trucksDTO, "should return empty trucks")
}

// TestUpdateTruck tests services.UpdateTruck using mockTruckRepo and gin
func TestUpdateTruck(t *testing.T) {
	mockRepo := newMockTruckRepo()
	mockService := NewTruckService(mockRepo)

	mockTruck := models.Truck{
		LicensePlate:  "AA444",
		ChassisNumber: "AA444AA",
	}

	truckDTO, status, err := mockService.UpdateTruck(1, mockTruck)
	var mockTruckDTO models.TruckDTO
	mockTruck.ID = 1
	automapper.Map(mockTruck, &mockTruckDTO)
	assert.NoError(t, err, "should not return error")
	assert.Equal(t, http.StatusOK, status, "should return status ok")
	assert.Equal(t, mockTruckDTO, truckDTO, "should return truck")
}

// TestUpdateTruck_FindError tests services.UpdateTruck using mockTruckRepo and gin
func TestUpdateTruck_FindError(t *testing.T) {
	mockRepo := newMockTruckErrorRepo()
	mockService := NewTruckService(mockRepo)

	mockTruck := models.Truck{
		LicensePlate:  "AA444",
		ChassisNumber: "AA444AA",
	}

	truckDTO, status, err := mockService.UpdateTruck(1, mockTruck)
	assert.Error(t, err, "should return error")
	assert.Equal(t, http.StatusNotFound, status, "should return status not found error")
	assert.Equal(t, models.TruckDTO{}, truckDTO, "should return empty truck")
}

// TestUpdateTruck_UpdateError tests services.UpdateTruck using mockTruckRepo and gin
func TestUpdateTruck_UpdateError(t *testing.T) {
	mockRepo := newMockTruckSpecificErrorRepo()
	mockService := NewTruckService(mockRepo)

	mockTruck := models.Truck{
		LicensePlate:  "AA444",
		ChassisNumber: "AA444AA",
	}

	truckDTO, status, err := mockService.UpdateTruck(1, mockTruck)
	assert.Error(t, err, "should return error")
	assert.Equal(t, http.StatusInternalServerError, status, "should return status internal server error")
	assert.Equal(t, models.TruckDTO{}, truckDTO, "should return empty truck")
}

// TestDeleteTruck tests services.DeleteTruck using mockTruckRepo and gin
func TestDeleteTruck(t *testing.T) {
	mockRepo := newMockTruckRepo()
	mockService := NewTruckService(mockRepo)

	truckDTO, status, err := mockService.DeleteTruck(1)
	var mockTruckDTO models.TruckDTO
	automapper.Map(mockTrucks[0], &mockTruckDTO)
	assert.NoError(t, err, "should not return error")
	assert.Equal(t, http.StatusOK, status, "should return status ok")
	assert.Equal(t, mockTruckDTO, truckDTO, "should return truck")
}

// TestDeleteTruck_FindError tests services.DeleteTruck using mockTruckRepo and gin
func TestDeleteTruck_FindError(t *testing.T) {
	mockRepo := newMockTruckErrorRepo()
	mockService := NewTruckService(mockRepo)

	truckDTO, status, err := mockService.DeleteTruck(1)
	assert.Error(t, err, "should return error")
	assert.Equal(t, http.StatusNotFound, status, "should return status not found error")
	assert.Equal(t, models.TruckDTO{}, truckDTO, "should return empty truck")
}

// TestDeleteTruck_DeleteError tests services.DeleteTruck using mockTruckRepo and gin
func TestDeleteTruck_DeleteError(t *testing.T) {
	mockRepo := newMockTruckSpecificErrorRepo()
	mockService := NewTruckService(mockRepo)

	truckDTO, status, err := mockService.DeleteTruck(1)
	assert.Error(t, err, "should return error")
	assert.Equal(t, http.StatusInternalServerError, status, "should return status internal server error")
	assert.Equal(t, models.TruckDTO{}, truckDTO, "should return empty truck")
}
