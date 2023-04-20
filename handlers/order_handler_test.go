package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/laertkokona/crud-test/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mockModels = []gorm.Model{
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
	{
		ID: uint(6),
	},
}
var mockOrders = []models.Order{
	{
		Model: mockModels[0],
		Code:  "ord1",
		OrderItems: []models.OrderItem{
			{
				Model:    mockModels[0],
				ItemId:   1,
				Quantity: 10,
			},
			{
				Model:    mockModels[1],
				ItemId:   2,
				Quantity: 20,
			},
		},
	},
	{
		Model: mockModels[1],
		Code:  "ord2",
		OrderItems: []models.OrderItem{
			{
				Model:    mockModels[2],
				ItemId:   3,
				Quantity: 30,
			},
			{
				Model:    mockModels[3],
				ItemId:   4,
				Quantity: 40,
			},
		},
	},
}

// mockOrderService is a mock implementation of the services.OrderService interface
type mockOrderService struct {
	createOrder  func(order models.Order) (models.Order, int, error)
	getOrder     func(id int) (models.Order, int, error)
	getAllOrders func(pagination models.Pagination) ([]models.Order, int, error)
	updateOrder  func(id int, order models.Order) (models.Order, int, error)
	deleteOrder  func(id int) (models.Order, int, error)
}

// CreateOrder is a mock implementation of the services.OrderService.CreateOrder method
func (m *mockOrderService) CreateOrder(order models.Order) (models.Order, int, error) {
	return m.createOrder(order)
}

// GetOrder is a mock implementation of the services.OrderService.GetOrder method
func (m *mockOrderService) GetOrder(id int) (models.Order, int, error) {
	return m.getOrder(id)
}

// GetAllOrders is a mock implementation of the services.OrderService.GetAllOrders method
func (m *mockOrderService) GetAllOrders(pagination models.Pagination) ([]models.Order, int, error) {
	return m.getAllOrders(pagination)
}

// UpdateOrder is a mock implementation of the services.OrderService.UpdateOrder method
func (m *mockOrderService) UpdateOrder(id int, order models.Order) (models.Order, int, error) {
	return m.updateOrder(id, order)
}

// DeleteOrder is a mock implementation of the services.OrderService.DeleteOrder method
func (m *mockOrderService) DeleteOrder(id int) (models.Order, int, error) {
	return m.deleteOrder(id)
}

// newMockOrderService returns a new instance of mockOrderService
func newMockOrderService() *mockOrderService {
	return &mockOrderService{
		createOrder: func(order models.Order) (models.Order, int, error) {
			return order, http.StatusOK, nil
		},
		getOrder: func(id int) (models.Order, int, error) {
			return mockOrders[id-1], http.StatusOK, nil
		},
		getAllOrders: func(pagination models.Pagination) ([]models.Order, int, error) {
			return mockOrders, http.StatusOK, nil
		},
		updateOrder: func(id int, order models.Order) (models.Order, int, error) {
			return order, http.StatusOK, nil
		},
		deleteOrder: func(id int) (models.Order, int, error) {
			return mockOrders[id-1], http.StatusOK, nil
		},
	}
}

// NewMockOrderErrorService returns a new instance of mockOrderService with errors
func NewMockOrderErrorService() *mockOrderService {
	return &mockOrderService{
		createOrder: func(order models.Order) (models.Order, int, error) {
			return order, http.StatusInternalServerError, errors.New("error creating order")
		},
		getOrder: func(id int) (models.Order, int, error) {
			return models.Order{}, http.StatusInternalServerError, errors.New("error getting order")
		},
		getAllOrders: func(pagination models.Pagination) ([]models.Order, int, error) {
			return []models.Order{}, http.StatusInternalServerError, errors.New("error getting all orders")
		},
		updateOrder: func(id int, order models.Order) (models.Order, int, error) {
			return models.Order{}, http.StatusInternalServerError, errors.New("error updating order")
		},
		deleteOrder: func(id int) (models.Order, int, error) {
			return models.Order{}, http.StatusInternalServerError, errors.New("error deleting order")
		},
	}
}

// TestCreateOrder tests the CreateOrder method
func TestCreateOrder(t *testing.T) {
	mockOrderService := newMockOrderService()

	r := gin.Default()
	mockOrderString, err := json.Marshal(mockOrders[0])
	assert.NoError(t, err)
	orderHandler := NewOrderHandler(mockOrderService)
	r.POST("/orders", orderHandler.CreateOrder)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(mockOrderString))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var order models.Order
	err = json.Unmarshal(w.Body.Bytes(), &order)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.Equal(t, mockOrders[0], order)
}

// TestCreateOrder_BindError tests the CreateOrder method with a bind error
func TestCreateOrder_BindError(t *testing.T) {
	mockOrderService := newMockOrderService()

	r := gin.Default()
	orderHandler := NewOrderHandler(mockOrderService)
	r.POST("/orders", orderHandler.CreateOrder)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer([]byte("{")))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}

// TestCreateOrder_ServiceError tests the CreateOrder method with a service error
func TestCreateOrder_ServiceError(t *testing.T) {
	mockOrderService := NewMockOrderErrorService()

	r := gin.Default()
	mockOrderString, err := json.Marshal(mockOrders[0])
	assert.NoError(t, err)
	orderHandler := NewOrderHandler(mockOrderService)
	r.POST("/orders", orderHandler.CreateOrder)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(mockOrderString))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}

// TestGetOrder tests the GetOrder method
func TestGetOrder(t *testing.T) {
	mockOrderService := newMockOrderService()

	r := gin.Default()
	orderHandler := NewOrderHandler(mockOrderService)
	r.GET("/orders/:id", orderHandler.GetOrder)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/orders/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var order models.Order
	err := json.Unmarshal(w.Body.Bytes(), &order)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.Equal(t, mockOrders[0], order)
}

// TestGetOrder_ServiceError tests the GetOrder method with a service error
func TestGetOrder_ServiceError(t *testing.T) {
	mockOrderService := NewMockOrderErrorService()

	r := gin.Default()
	orderHandler := NewOrderHandler(mockOrderService)
	r.GET("/orders/:id", orderHandler.GetOrder)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/orders/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}

// TestGetOrder_InvalidIDError tests the GetOrder method with an invalid id
func TestGetOrder_InvalidIDError(t *testing.T) {
	mockOrderService := newMockOrderService()

	r := gin.Default()
	orderHandler := NewOrderHandler(mockOrderService)
	r.GET("/orders/:id", orderHandler.GetOrder)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/orders/invalid", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}

// TestGetAllOrders tests the GetAllOrders method
func TestGetAllOrders(t *testing.T) {
	mockOrderService := newMockOrderService()

	r := gin.Default()
	orderHandler := NewOrderHandler(mockOrderService)
	r.GET("/orders", orderHandler.GetAllOrders)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/orders", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var orders []models.Order
	err := json.Unmarshal(w.Body.Bytes(), &orders)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.Equal(t, mockOrders, orders)
}

// TestGetAllOrders_ServiceError tests the GetAllOrders method with a service error
func TestGetAllOrders_ServiceError(t *testing.T) {
	mockOrderService := NewMockOrderErrorService()

	r := gin.Default()
	orderHandler := NewOrderHandler(mockOrderService)
	r.GET("/orders", orderHandler.GetAllOrders)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/orders", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}

// TestUpdateOrder tests the UpdateOrder method
func TestUpdateOrder(t *testing.T) {
	mockOrderService := newMockOrderService()

	r := gin.Default()
	mockOrderString, err := json.Marshal(mockOrders[0])
	assert.NoError(t, err)
	orderHandler := NewOrderHandler(mockOrderService)
	r.PUT("/orders/:id", orderHandler.UpdateOrder)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/orders/1", bytes.NewBuffer(mockOrderString))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var order models.Order
	err = json.Unmarshal(w.Body.Bytes(), &order)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.Equal(t, mockOrders[0], order)
}

// TestUpdateOrder_InvalidIDError tests the UpdateOrder method with an invalid id
func TestUpdateOrder_InvalidIDError(t *testing.T) {
	mockOrderService := newMockOrderService()

	r := gin.Default()
	mockOrderString, err := json.Marshal(mockOrders[0])
	assert.NoError(t, err)
	orderHandler := NewOrderHandler(mockOrderService)
	r.PUT("/orders/:id", orderHandler.UpdateOrder)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/orders/invalid", bytes.NewBuffer(mockOrderString))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}

// TestUpdateOrder_InvalidJSONError tests the UpdateOrder method with an invalid json
func TestUpdateOrder_InvalidJSONError(t *testing.T) {
	mockOrderService := newMockOrderService()

	r := gin.Default()
	orderHandler := NewOrderHandler(mockOrderService)
	r.PUT("/orders/:id", orderHandler.UpdateOrder)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/orders/1", bytes.NewBufferString("invalid"))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}

// TestUpdateOrder_ServiceError tests the UpdateOrder method with a service error
func TestUpdateOrder_ServiceError(t *testing.T) {
	mockOrderService := NewMockOrderErrorService()

	r := gin.Default()
	mockOrderString, err := json.Marshal(mockOrders[0])
	assert.NoError(t, err)
	orderHandler := NewOrderHandler(mockOrderService)
	r.PUT("/orders/:id", orderHandler.UpdateOrder)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/orders/1", bytes.NewBuffer(mockOrderString))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}

// TestDeleteOrder tests the DeleteOrder method
func TestDeleteOrder(t *testing.T) {
	mockOrderService := newMockOrderService()

	r := gin.Default()
	orderHandler := NewOrderHandler(mockOrderService)
	r.DELETE("/orders/:id", orderHandler.DeleteOrder)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/orders/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var order models.Order
	err := json.Unmarshal(w.Body.Bytes(), &order)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.Equal(t, mockOrders[0], order)
}

// TestDeleteOrder_InvalidIDError tests the DeleteOrder method with an invalid id
func TestDeleteOrder_InvalidIDError(t *testing.T) {
	mockOrderService := newMockOrderService()

	r := gin.Default()
	orderHandler := NewOrderHandler(mockOrderService)
	r.DELETE("/orders/:id", orderHandler.DeleteOrder)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/orders/invalid", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}

// TestDeleteOrder_ServiceError tests the DeleteOrder method with a service error
func TestDeleteOrder_ServiceError(t *testing.T) {
	mockOrderService := NewMockOrderErrorService()

	r := gin.Default()
	orderHandler := NewOrderHandler(mockOrderService)
	r.DELETE("/orders/:id", orderHandler.DeleteOrder)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/orders/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Error while unmarshalling response: %v", err)
	assert.NotNil(t, response["error"])
}
