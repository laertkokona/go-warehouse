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

var itmModels = []gorm.Model{
	{
		ID: uint(1),
	},
	{
		ID: uint(2),
	},
	{
		ID: uint(3),
	},
	{
		ID: uint(4),
	},
	{
		ID: uint(5),
	},
}
var mockItems = []models.Item{
	{
		Model:             itmModels[0],
		Name:              "Item 1",
		Description:       "Description Item 1",
		Code:              "itm1",
		TotalQuantity:     100,
		AvailableQuantity: 100,
		Price:             9.99,
		Category:          "Category Test",
	},
	{
		Model:             itmModels[1],
		Name:              "Item 2",
		Description:       "Description Item 2",
		Code:              "itm2",
		TotalQuantity:     200,
		AvailableQuantity: 200,
		Price:             19.99,
		Category:          "Category Test",
	},
	{
		Model:             itmModels[2],
		Name:              "Item 3",
		Description:       "Description Item 3",
		Code:              "itm3",
		TotalQuantity:     300,
		AvailableQuantity: 300,
		Price:             29.99,
		Category:          "Category Test",
	},
	{
		Model:             itmModels[3],
		Name:              "Item 4",
		Description:       "Description Item 4",
		Code:              "itm4",
		TotalQuantity:     400,
		AvailableQuantity: 400,
		Price:             39.99,
		Category:          "Category Test",
	},
	{
		Model:             itmModels[4],
		Name:              "Item 5",
		Description:       "Description Item 5",
		Code:              "itm5",
		TotalQuantity:     500,
		AvailableQuantity: 500,
		Price:             49.99,
		Category:          "Category Test",
	},
}

// mockItemRepo is a mock implementation of the repositories.ItemRepo interface
type mockItemRepo struct {
	// findAll is a mock function with given fields: pagination
	findAll func(pagination models.Pagination) ([]models.Item, error)
	// findByID is a mock function with given fields: id
	findByID func(id int) (models.Item, error)
	// findByName is a mock function with given fields: name
	findByName func(name string) (models.Item, error)
	// save is a mock function with given fields: item
	save func(item models.Item) (models.Item, error)
	// update is a mock function with given fields: item
	update func(item models.Item) (models.Item, error)
	// delete is a mock function with given fields: item
	delete func(item models.Item) error
	// deleteById is a mock function with given fields: id
	deleteById func(id int) (models.Item, error)
}

// FindAll is a mock function with given fields: pagination
func (_m *mockItemRepo) FindAll(pagination models.Pagination) ([]models.Item, error) {
	return _m.findAll(pagination)
}

// FindByID is a mock function with given fields: id
func (_m *mockItemRepo) FindByID(id int) (models.Item, error) {
	return _m.findByID(id)
}

// FindByName is a mock function with given fields: name
func (_m *mockItemRepo) FindByName(name string) (models.Item, error) {
	return _m.findByName(name)
}

// Save is a mock function with given fields: item
func (_m *mockItemRepo) Save(item models.Item) (models.Item, error) {
	return _m.save(item)
}

// Update is a mock function with given fields: item
func (_m *mockItemRepo) Update(item models.Item) (models.Item, error) {
	return _m.update(item)
}

// Delete is a mock function with given fields: item
func (_m *mockItemRepo) Delete(item models.Item) error {
	return _m.delete(item)
}

// DeleteById is a mock function with given fields: id
func (_m *mockItemRepo) DeleteById(id int) (models.Item, error) {
	return _m.deleteById(id)
}

// newMockItemRepo returns a new instance of the mockItemRepo
func newMockItemRepo() *mockItemRepo {
	return &mockItemRepo{
		findAll: func(pagination models.Pagination) ([]models.Item, error) {
			return mockItems, nil
		},
		findByID: func(id int) (models.Item, error) {
			return mockItems[id-1], nil
		},
		findByName: func(name string) (models.Item, error) {
			var itm models.Item
			for _, item := range mockItems {
				if item.Name == name {
					itm = item
				}
			}
			return itm, nil
		},
		save: func(item models.Item) (models.Item, error) {
			return item, nil
		},
		update: func(item models.Item) (models.Item, error) {
			item.ID = 1
			return item, nil
		},
		delete: func(item models.Item) error {
			return nil
		},
		deleteById: func(id int) (models.Item, error) {
			return mockItems[id-1], nil
		},
	}
}

// ERROR MOCK

// newMockItemErrorRepo returns a new instance of the mockItemErrorRepo
func newMockItemErrorRepo() *mockItemRepo {
	return &mockItemRepo{
		findAll: func(pagination models.Pagination) ([]models.Item, error) {
			return []models.Item{}, errors.New("error")
		},
		findByID: func(id int) (models.Item, error) {
			return models.Item{}, errors.New("error")
		},
		findByName: func(name string) (models.Item, error) {
			return models.Item{}, errors.New("error")
		},
		save: func(item models.Item) (models.Item, error) {
			return models.Item{}, errors.New("error")
		},
		update: func(item models.Item) (models.Item, error) {
			return models.Item{}, errors.New("error")
		},
		delete: func(item models.Item) error {
			return errors.New("error")
		},
		deleteById: func(id int) (models.Item, error) {
			return models.Item{}, errors.New("error")
		},
	}
}

// newMockItemSpecificErrorRepo returns a new instance of the mockItemSpecificErrorRepo
func newMockItemSpecificErrorRepo() *mockItemRepo {
	return &mockItemRepo{
		findAll: func(pagination models.Pagination) ([]models.Item, error) {
			return mockItems, nil
		},
		findByID: func(id int) (models.Item, error) {
			return mockItems[id-1], nil
		},
		findByName: func(name string) (models.Item, error) {
			var itm models.Item
			for _, item := range mockItems {
				if item.Name == name {
					itm = item
				}
			}
			return itm, nil
		},
		save: func(item models.Item) (models.Item, error) {
			return models.Item{}, errors.New("error")
		},
		update: func(item models.Item) (models.Item, error) {
			return models.Item{}, errors.New("error")
		},
		delete: func(item models.Item) error {
			return errors.New("error")
		},
		deleteById: func(id int) (models.Item, error) {
			return models.Item{}, errors.New("error")
		},
	}
}

// TestNewItemService is a test function for the NewItemService function
func TestNewItemService(t *testing.T) {
	mockRepo := newMockItemRepo()
	mockService := NewItemService(mockRepo)
	assert.NotNil(t, mockService)
	assert.IsType(t, itemService{}, mockService)
}

// TestCreateItem tests services.CreateItem function using a mock repository mockItemRepo and gin
func TestCreateItem(t *testing.T) {
	mockRepo := newMockItemRepo()
	mockService := NewItemService(mockRepo)

	mockItem := models.Item{
		Name:              "Item 6",
		Description:       "Description Item 6",
		Code:              "itm6",
		TotalQuantity:     600,
		AvailableQuantity: 600,
		Price:             59.99,
		Category:          "Category Test",
	}

	itemDTO, status, err := mockService.CreateItem(mockItem)
	var mockItemDTO models.ItemDTO
	automapper.Map(mockItem, &mockItemDTO)
	assert.NoError(t, err, "Error while creating item: %v", err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, mockItemDTO, itemDTO)
}

// TestCreateItem_SaveError tests services.CreateItem function using a mock repository mockItemErrorRepo and gin
func TestCreateItem_SaveError(t *testing.T) {
	mockRepo := newMockItemErrorRepo()
	mockService := NewItemService(mockRepo)

	mockItem := models.Item{
		Name:              "Item 6",
		Description:       "Description Item 6",
		Code:              "itm6",
		TotalQuantity:     600,
		AvailableQuantity: 600,
		Price:             59.99,
		Category:          "Category Test",
	}

	itemDTO, status, err := mockService.CreateItem(mockItem)
	assert.Error(t, err, "Error while creating item: %v", err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, models.ItemDTO{}, itemDTO)
}

// TestGetItem tests services.GetItem function using a mock repository mockItemRepo and gin
func TestGetItem(t *testing.T) {
	mockRepo := newMockItemRepo()
	mockService := NewItemService(mockRepo)

	itemDTO, status, err := mockService.GetItem(1)
	var mockItemDTO models.ItemDTO
	automapper.Map(mockItems[0], &mockItemDTO)
	assert.NoError(t, err, "Error while getting item: %v", err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, mockItemDTO, itemDTO)
}

// TestGetItem_FindByIDError tests services.GetItem function using a mock repository mockItemErrorRepo and gin
func TestGetItem_FindByIDError(t *testing.T) {
	mockRepo := newMockItemErrorRepo()
	mockService := NewItemService(mockRepo)

	itemDTO, status, err := mockService.GetItem(1)
	assert.Error(t, err, "Error while getting item: %v", err)
	assert.Equal(t, http.StatusNotFound, status)
	assert.Equal(t, models.ItemDTO{}, itemDTO)
}

// TestGetAllItems tests services.GetAllItems function using a mock repository mockItemRepo and gin
func TestGetAllItems(t *testing.T) {
	mockRepo := newMockItemRepo()
	mockService := NewItemService(mockRepo)

	itemsDTO, status, err := mockService.GetAllItems(models.Pagination{})
	var mockItemsDTO []models.ItemDTO
	automapper.Map(mockItems, &mockItemsDTO)
	assert.NoError(t, err, "Error while getting all items: %v", err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, mockItemsDTO, itemsDTO)
}

// TestGetAllItems_FindAllError tests services.GetAllItems function using a mock repository mockItemErrorRepo and gin
func TestGetAllItems_FindAllError(t *testing.T) {
	mockRepo := newMockItemErrorRepo()
	mockService := NewItemService(mockRepo)

	itemsDTO, status, err := mockService.GetAllItems(models.Pagination{})
	assert.Error(t, err, "Error while getting all items: %v", err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, []models.ItemDTO{}, itemsDTO)
}

// TestUpdateItem tests services.UpdateItem function using a mock repository mockItemRepo and gin
func TestUpdateItem(t *testing.T) {
	mockRepo := newMockItemRepo()
	mockService := NewItemService(mockRepo)

	mockItem := models.Item{
		Name:              "Item 6",
		Description:       "Description Item 6",
		Code:              "itm6",
		TotalQuantity:     600,
		AvailableQuantity: 600,
		Price:             59.99,
		Category:          "Category Test",
	}
	itemDTO, status, err := mockService.UpdateItem(1, mockItem)
	mockItem.ID = 1
	var mockItemDTO models.ItemDTO
	automapper.Map(mockItem, &mockItemDTO)
	assert.NoError(t, err, "Error while updating item: %v", err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, mockItemDTO, itemDTO)
}

// TestUpdateItem_FindByIDError tests services.UpdateItem function using a mock repository mockItemErrorRepo and gin
func TestUpdateItem_FindByIDError(t *testing.T) {
	mockRepo := newMockItemErrorRepo()
	mockService := NewItemService(mockRepo)

	mockItem := models.Item{
		Name:              "Item 6",
		Description:       "Description Item 6",
		Code:              "itm6",
		TotalQuantity:     600,
		AvailableQuantity: 600,
		Price:             59.99,
		Category:          "Category Test",
	}
	itemDTO, status, err := mockService.UpdateItem(1, mockItem)
	assert.Error(t, err, "Error while updating item: %v", err)
	assert.Equal(t, http.StatusNotFound, status)
	assert.Equal(t, models.ItemDTO{}, itemDTO)
}

// TestUpdateItem_UpdateError tests services.UpdateItem function using a mock repository mockItemErrorRepo and gin
func TestUpdateItem_UpdateError(t *testing.T) {
	mockRepo := newMockItemSpecificErrorRepo()
	mockService := NewItemService(mockRepo)

	mockItem := models.Item{
		Name:              "Item 6",
		Description:       "Description Item 6",
		Code:              "itm6",
		TotalQuantity:     600,
		AvailableQuantity: 600,
		Price:             59.99,
		Category:          "Category Test",
	}
	itemDTO, status, err := mockService.UpdateItem(1, mockItem)
	assert.Error(t, err, "Error while updating item: %v", err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, models.ItemDTO{}, itemDTO)
}

// TestDeleteItem tests services.DeleteItem function using a mock repository mockItemRepo and gin
func TestDeleteItem(t *testing.T) {
	mockRepo := newMockItemRepo()
	mockService := NewItemService(mockRepo)

	itemDTO, status, err := mockService.DeleteItem(1)
	var mockItemDTO models.ItemDTO
	automapper.Map(mockItems[0], &mockItemDTO)
	assert.NoError(t, err, "Error while deleting item: %v", err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, mockItemDTO, itemDTO)
}

// TestDeleteItem_FindByIDError tests services.DeleteItem function using a mock repository mockItemErrorRepo and gin
func TestDeleteItem_FindByIDError(t *testing.T) {
	mockRepo := newMockItemErrorRepo()
	mockService := NewItemService(mockRepo)

	itemDTO, status, err := mockService.DeleteItem(1)
	assert.Error(t, err, "Error while deleting item: %v", err)
	assert.Equal(t, http.StatusNotFound, status)
	assert.Equal(t, models.ItemDTO{}, itemDTO)
}

// TestDeleteItem_DeleteError tests services.DeleteItem function using a mock repository mockItemSpecificErrorRepo and gin
func TestDeleteItem_DeleteError(t *testing.T) {
	mockRepo := newMockItemSpecificErrorRepo()
	mockService := NewItemService(mockRepo)

	itemDTO, status, err := mockService.DeleteItem(1)
	assert.Error(t, err, "Error while deleting item: %v", err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, models.ItemDTO{}, itemDTO)
}
