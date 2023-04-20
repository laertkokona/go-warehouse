package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/laertkokona/crud-test/models"
	"github.com/peteprogrammer/go-automapper"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
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

// mockItemService struct that implements the ItemService interface
type mockItemService struct {
	createItem  func(item models.Item) (models.ItemDTO, int, error)
	getItem     func(id int) (models.ItemDTO, int, error)
	getAllItems func(pagination models.Pagination) ([]models.ItemDTO, int, error)
	updateItem  func(id int, item models.Item) (models.ItemDTO, int, error)
	deleteItem  func(id int) (models.ItemDTO, int, error)
}

// CreateItem mock function
func (_m *mockItemService) CreateItem(item models.Item) (models.ItemDTO, int, error) {
	return _m.createItem(item)
}

// GetItem mock function
func (_m *mockItemService) GetItem(id int) (models.ItemDTO, int, error) {
	return _m.getItem(id)
}

// GetAllItems mock function
func (_m *mockItemService) GetAllItems(pagination models.Pagination) ([]models.ItemDTO, int, error) {
	return _m.getAllItems(pagination)
}

// UpdateItem mock function
func (_m *mockItemService) UpdateItem(id int, item models.Item) (models.ItemDTO, int, error) {
	return _m.updateItem(id, item)
}

// DeleteItem mock function
func (_m *mockItemService) DeleteItem(id int) (models.ItemDTO, int, error) {
	return _m.deleteItem(id)
}

// newMockItemService returns a new instance of mockItemService
func newMockItemService() *mockItemService {
	return &mockItemService{
		createItem: func(item models.Item) (models.ItemDTO, int, error) {
			var itemDTO models.ItemDTO
			automapper.Map(item, &itemDTO)
			return itemDTO, http.StatusOK, nil
		},
		getItem: func(id int) (models.ItemDTO, int, error) {
			var itemDTO models.ItemDTO
			automapper.Map(mockItems[id-1], &itemDTO)
			return itemDTO, http.StatusOK, nil
		},
		getAllItems: func(pagination models.Pagination) ([]models.ItemDTO, int, error) {
			var mockItemsDTO []models.ItemDTO
			automapper.Map(mockItems, &mockItems)
			return mockItemsDTO, http.StatusOK, nil
		},
		updateItem: func(id int, item models.Item) (models.ItemDTO, int, error) {
			var itemDTO models.ItemDTO
			item.ID = uint(id)
			automapper.Map(item, &itemDTO)
			return itemDTO, http.StatusOK, nil
		},
		deleteItem: func(id int) (models.ItemDTO, int, error) {
			var itemDTO models.ItemDTO
			automapper.Map(mockItems[id-1], &itemDTO)
			return itemDTO, http.StatusOK, nil
		},
	}
}

// newMockItemErrorService returns a new instance of mockItemService with errors
func newMockItemErrorService() *mockItemService {
	return &mockItemService{
		createItem: func(item models.Item) (models.ItemDTO, int, error) {
			return models.ItemDTO{}, http.StatusInternalServerError, errors.New("error while creating item")
		},
		getItem: func(id int) (models.ItemDTO, int, error) {
			return models.ItemDTO{}, http.StatusInternalServerError, errors.New("error while getting item")
		},
		getAllItems: func(pagination models.Pagination) ([]models.ItemDTO, int, error) {
			return []models.ItemDTO{}, http.StatusInternalServerError, errors.New("error while getting all items")
		},
		updateItem: func(id int, item models.Item) (models.ItemDTO, int, error) {
			return models.ItemDTO{}, http.StatusInternalServerError, errors.New("error while updating item")
		},
		deleteItem: func(id int) (models.ItemDTO, int, error) {
			return models.ItemDTO{}, http.StatusInternalServerError, errors.New("error while deleting item")
		},
	}
}

// TestCreateItem tests the CreateItem function
func TestCreateItem(t *testing.T) {
	mockService := newMockItemService()

	r := gin.Default()
	mockItemString, err := json.Marshal(mockItems[0])
	assert.NoError(t, err)
	itemHandler := NewItemHandler(mockService)
	r.POST("/items", itemHandler.CreateItem)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(mockItemString))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var item models.ItemDTO
	var mockItemDTO models.ItemDTO
	automapper.Map(mockItems[0], &mockItemDTO)
	err = json.Unmarshal(w.Body.Bytes(), &item)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.Equal(t, mockItemDTO, item)
}

// TestCreateItem_BindError tests the CreateItem function with a bind error
func TestCreateItem_BindError(t *testing.T) {
	mockService := newMockItemService()

	r := gin.Default()
	itemHandler := NewItemHandler(mockService)
	r.POST("/items", itemHandler.CreateItem)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}

// TestCreateItem_ServiceError tests the CreateItem function with a service error
func TestCreateItem_ServiceError(t *testing.T) {
	mockService := newMockItemErrorService()

	r := gin.Default()
	mockItemString, err := json.Marshal(mockItems[0])
	assert.NoError(t, err)
	itemHandler := NewItemHandler(mockService)
	r.POST("/items", itemHandler.CreateItem)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(mockItemString))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}

// TestGetItem tests the GetItem function
func TestGetItem(t *testing.T) {
	mockService := newMockItemService()

	r := gin.Default()
	itemHandler := NewItemHandler(mockService)
	r.GET("/items/:id", itemHandler.GetItem)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var item models.Item
	err := json.Unmarshal(w.Body.Bytes(), &item)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.Equal(t, mockItems[0], item)
}

// TestGetItem_InvalidIDError tests the GetItem function with an invalid id
func TestGetItem_InvalidIDError(t *testing.T) {
	mockService := newMockItemService()

	r := gin.Default()
	itemHandler := NewItemHandler(mockService)
	r.GET("/items/:id", itemHandler.GetItem)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items/invalid", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}

// TestGetItem_ServiceError tests the GetItem function with a service error
func TestGetItem_ServiceError(t *testing.T) {
	mockService := newMockItemErrorService()

	r := gin.Default()
	itemHandler := NewItemHandler(mockService)
	r.GET("/items/:id", itemHandler.GetItem)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}

// TestGetAllItems tests the GetAllItems function
func TestGetAllItems(t *testing.T) {
	mockService := newMockItemService()

	r := gin.Default()
	itemHandler := NewItemHandler(mockService)
	r.GET("/items", itemHandler.GetAllItems)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var items []models.ItemDTO
	err := json.Unmarshal(w.Body.Bytes(), &items)
	log.Println("Items: ", items)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.Equal(t, mockItems, items)
}

// TestGetAllItems_ServiceError tests the GetAllItems function with a service error
func TestGetAllItems_ServiceError(t *testing.T) {
	mockService := newMockItemErrorService()

	r := gin.Default()
	itemHandler := NewItemHandler(mockService)
	r.GET("/items", itemHandler.GetAllItems)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}

// TestUpdateItem tests the UpdateItem function
func TestUpdateItem(t *testing.T) {
	mockService := newMockItemService()

	r := gin.Default()
	mockItemString, err := json.Marshal(mockItems[0])
	assert.NoError(t, err)
	itemHandler := NewItemHandler(mockService)
	r.PUT("/items/:id", itemHandler.UpdateItem)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/items/1", bytes.NewBuffer(mockItemString))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var item models.Item
	err = json.Unmarshal(w.Body.Bytes(), &item)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.Equal(t, mockItems[0], item)
}

// TestUpdateItem_InvalidIDError tests the UpdateItem function with an invalid id
func TestUpdateItem_InvalidIDError(t *testing.T) {
	mockService := newMockItemService()

	r := gin.Default()
	mockItemString, err := json.Marshal(mockItems[0])
	assert.NoError(t, err)
	itemHandler := NewItemHandler(mockService)
	r.PUT("/items/:id", itemHandler.UpdateItem)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/items/invalid", bytes.NewBuffer(mockItemString))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}

// TestUpdateItem_InvalidBodyError tests the UpdateItem function with an invalid body
func TestUpdateItem_InvalidBodyError(t *testing.T) {
	mockService := newMockItemService()

	r := gin.Default()
	itemHandler := NewItemHandler(mockService)
	r.PUT("/items/:id", itemHandler.UpdateItem)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/items/1", bytes.NewBufferString("invalid"))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}

// TestUpdateItem_ServiceError tests the UpdateItem function with a service error
func TestUpdateItem_ServiceError(t *testing.T) {
	mockService := newMockItemErrorService()

	r := gin.Default()
	mockItemString, err := json.Marshal(mockItems[0])
	assert.NoError(t, err)
	itemHandler := NewItemHandler(mockService)
	r.PUT("/items/:id", itemHandler.UpdateItem)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/items/1", bytes.NewBuffer(mockItemString))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}

// TestDeleteItem tests the DeleteItem function
func TestDeleteItem(t *testing.T) {
	mockService := newMockItemService()

	r := gin.Default()
	itemHandler := NewItemHandler(mockService)
	r.DELETE("/items/:id", itemHandler.DeleteItem)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/items/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var item models.Item
	err := json.Unmarshal(w.Body.Bytes(), &item)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.Equal(t, mockItems[0], item)
}

// TestDeleteItem_InvalidIDError tests the DeleteItem function with an invalid id
func TestDeleteItem_InvalidIDError(t *testing.T) {
	mockService := newMockItemService()

	r := gin.Default()
	itemHandler := NewItemHandler(mockService)
	r.DELETE("/items/:id", itemHandler.DeleteItem)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/items/invalid", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}

// TestDeleteItem_ServiceError tests the DeleteItem function with a service error
func TestDeleteItem_ServiceError(t *testing.T) {
	mockService := newMockItemErrorService()

	r := gin.Default()
	itemHandler := NewItemHandler(mockService)
	r.DELETE("/items/:id", itemHandler.DeleteItem)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/items/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}
