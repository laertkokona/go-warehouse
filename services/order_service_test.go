package services

import (
	"errors"
	"github.com/laertkokona/crud-test/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
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

// mockOrderRepo is a mock implementation of the repositories.OrderRepo interface
type mockOrderRepo struct {
	// findAll is a mock function with given fields: pagination
	findAll func(pagination models.Pagination) ([]models.Order, error)
	// findByID is a mock function with given fields: id
	findByID func(id int) (models.Order, error)
	// save is a mock function with given fields: order
	save func(order models.Order) (models.Order, error)
	// update is a mock function with given fields: order
	update func(order models.Order) (models.Order, error)
	// delete is a mock function with given fields: order
	delete func(order models.Order) error
	// deleteById is a mock function with given fields: id
	deleteById func(id int) (models.Order, error)
}

// FindAll is a mock function with given fields: pagination
func (_m *mockOrderRepo) FindAll(pagination models.Pagination) ([]models.Order, error) {
	return _m.findAll(pagination)
}

// FindByID is a mock function with given fields: id
func (_m *mockOrderRepo) FindByID(id int) (models.Order, error) {
	return _m.findByID(id)
}

// Save is a mock function with given fields: order
func (_m *mockOrderRepo) Save(order models.Order) (models.Order, error) {
	return _m.save(order)
}

// Update is a mock function with given fields: order
func (_m *mockOrderRepo) Update(order models.Order) (models.Order, error) {
	return _m.update(order)
}

// Delete is a mock function with given fields: order
func (_m *mockOrderRepo) Delete(order models.Order) error {
	return _m.delete(order)
}

// DeleteById is a mock function with given fields: id
func (_m *mockOrderRepo) DeleteById(id int) (models.Order, error) {
	return _m.deleteById(id)
}

// newMockOrderRepo returns a new mockOrderRepo
func newMockOrderRepo() *mockOrderRepo {
	return &mockOrderRepo{
		findAll: func(pagination models.Pagination) ([]models.Order, error) {
			return mockOrders, nil
		},
		findByID: func(id int) (models.Order, error) {
			var order models.Order
			for _, o := range mockOrders {
				if o.ID == uint(id) {
					order = o
				}
			}
			return order, nil
		},
		save: func(order models.Order) (models.Order, error) {
			return order, nil
		},
		update: func(order models.Order) (models.Order, error) {
			order.ID = 1
			return order, nil
		},
		delete: func(order models.Order) error {
			return nil
		},
		deleteById: func(id int) (models.Order, error) {
			var order models.Order
			for _, o := range mockOrders {
				if o.ID == uint(id) {
					order = o
				}
			}
			return order, nil
		},
	}
}

// ERROR MOCK

// newMockOrderErrorRepo returns a new mockOrderErrorRepo
func newMockOrderErrorRepo() *mockOrderRepo {
	return &mockOrderRepo{
		findAll: func(pagination models.Pagination) ([]models.Order, error) {
			return []models.Order{}, errors.New("error")
		},
		findByID: func(id int) (models.Order, error) {
			return models.Order{}, errors.New("error")
		},
		save: func(order models.Order) (models.Order, error) {
			return models.Order{}, errors.New("error")
		},
		update: func(order models.Order) (models.Order, error) {
			return models.Order{}, errors.New("error")
		},
		delete: func(order models.Order) error {
			return errors.New("error")
		},
		deleteById: func(id int) (models.Order, error) {
			return models.Order{}, errors.New("error")
		},
	}
}

// newMockOrderSpecificErrorRepo returns a new mockOrderErrorRepo
func newMockOrderSpecificErrorRepo() *mockOrderRepo {
	return &mockOrderRepo{
		findAll: func(pagination models.Pagination) ([]models.Order, error) {
			return mockOrders, nil
		},
		findByID: func(id int) (models.Order, error) {
			var order models.Order
			for _, o := range mockOrders {
				if o.ID == uint(id) {
					order = o
				}
			}
			return order, nil
		},
		save: func(order models.Order) (models.Order, error) {
			return models.Order{}, errors.New("error")
		},
		update: func(order models.Order) (models.Order, error) {
			return models.Order{}, errors.New("error")
		},
		delete: func(order models.Order) error {
			return errors.New("error")
		},
		deleteById: func(id int) (models.Order, error) {
			return models.Order{}, errors.New("error")
		},
	}
}

// TestNewOrderService test the NewOrderService function
func TestNewOrderService(t *testing.T) {
	mockOrderRepo := newMockOrderRepo()
	mockService := NewOrderService(mockOrderRepo)

	assert.NotNil(t, mockService)
	assert.IsType(t, orderService{}, mockService)
}

// TestCreateOrder test the CreateOrder function using mockOrderRepo
func TestCreateOrder(t *testing.T) {
	mockOrderRepo := newMockOrderRepo()
	mockService := NewOrderService(mockOrderRepo)

	mockOrder := models.Order{
		Code: "ord3",
		OrderItems: []models.OrderItem{
			{
				ItemId:   5,
				Quantity: 50,
			},
		},
	}

	order, status, err := mockService.CreateOrder(mockOrder)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, mockOrder, order)
}

// TestCreateOrder_SaveError test the CreateOrder function using mockOrderErrorRepo
func TestCreateOrder_SaveError(t *testing.T) {
	mockOrderRepo := newMockOrderErrorRepo()
	mockService := NewOrderService(mockOrderRepo)

	mockOrder := models.Order{
		Code: "ord3",
		OrderItems: []models.OrderItem{
			{
				ItemId:   5,
				Quantity: 50,
			},
		},
	}

	order, status, err := mockService.CreateOrder(mockOrder)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, models.Order{}, order)
}

// TestGetOrder test the GetOrder function using mockOrderRepo
func TestGetOrder(t *testing.T) {
	mockOrderRepo := newMockOrderRepo()
	mockService := NewOrderService(mockOrderRepo)

	order, status, err := mockService.GetOrder(1)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, mockOrders[0], order)
}

// TestGetOrder_FindByIdError test the GetOrder function using mockOrderErrorRepo
func TestGetOrder_FindByIdError(t *testing.T) {
	mockOrderRepo := newMockOrderErrorRepo()
	mockService := NewOrderService(mockOrderRepo)

	order, status, err := mockService.GetOrder(1)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, models.Order{}, order)
}

// TestGetAllOrders test the GetAllOrders function using mockOrderRepo
func TestGetAllOrders(t *testing.T) {
	mockOrderRepo := newMockOrderRepo()
	mockService := NewOrderService(mockOrderRepo)

	orders, status, err := mockService.GetAllOrders(models.Pagination{})
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, mockOrders, orders)
}

// TestGetAllOrders_FindAllError test the GetAllOrders function using mockOrderErrorRepo
func TestGetAllOrders_FindAllError(t *testing.T) {
	mockOrderRepo := newMockOrderErrorRepo()
	mockService := NewOrderService(mockOrderRepo)

	orders, status, err := mockService.GetAllOrders(models.Pagination{})
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, []models.Order{}, orders)
}

// TestUpdateOrder test the UpdateOrder function using mockOrderRepo
func TestUpdateOrder(t *testing.T) {
	mockOrderRepo := newMockOrderRepo()
	mockService := NewOrderService(mockOrderRepo)

	mockOrder := models.Order{
		Code: "ord3",
		OrderItems: []models.OrderItem{
			{
				ItemId:   5,
				Quantity: 50,
			},
		},
	}

	order, status, err := mockService.UpdateOrder(1, mockOrder)
	mockOrder.ID = uint(1)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, mockOrder, order)
}

// TestUpdateOrder_FindByIdError test the UpdateOrder function using mockOrderErrorRepo
func TestUpdateOrder_FindByIdError(t *testing.T) {
	mockOrderRepo := newMockOrderErrorRepo()
	mockService := NewOrderService(mockOrderRepo)

	mockOrder := models.Order{
		Code: "ord3",
		OrderItems: []models.OrderItem{
			{
				ItemId:   5,
				Quantity: 50,
			},
		},
	}

	order, status, err := mockService.UpdateOrder(1, mockOrder)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, status)
	assert.Equal(t, models.Order{}, order)
}

// TestUpdateOrder_UpdateError test the UpdateOrder function using mockOrderSpecificErrorRepo
func TestUpdateOrder_UpdateError(t *testing.T) {
	mockOrderRepo := newMockOrderSpecificErrorRepo()
	mockService := NewOrderService(mockOrderRepo)

	mockOrder := models.Order{
		Code: "ord3",
		OrderItems: []models.OrderItem{
			{
				ItemId:   5,
				Quantity: 50,
			},
		},
	}

	order, status, err := mockService.UpdateOrder(1, mockOrder)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, models.Order{}, order)
}

// TestDeleteOrder test the DeleteOrder function using mockOrderRepo
func TestDeleteOrder(t *testing.T) {
	mockOrderRepo := newMockOrderRepo()
	mockService := NewOrderService(mockOrderRepo)

	order, status, err := mockService.DeleteOrder(1)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, mockOrders[0], order)
}

// TestDeleteOrder_FindByIdError test the DeleteOrder function using mockOrderErrorRepo
func TestDeleteOrder_FindByIdError(t *testing.T) {
	mockOrderRepo := newMockOrderErrorRepo()
	mockService := NewOrderService(mockOrderRepo)

	order, status, err := mockService.DeleteOrder(1)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, status)
	assert.Equal(t, models.Order{}, order)
}

// TestDeleteOrder_DeleteError test the DeleteOrder function using mockOrderSpecificErrorRepo
func TestDeleteOrder_DeleteError(t *testing.T) {
	mockOrderRepo := newMockOrderSpecificErrorRepo()
	mockService := NewOrderService(mockOrderRepo)

	order, status, err := mockService.DeleteOrder(1)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, mockOrders[0], order)
}

//// TestCreateOrder test the CreateOrder function using mockOrderRepo and gin
//func TestCreateOrder(t *testing.T) {
//	mockOrderRepo := newMockOrderRepo()
//	mockService := NewOrderService(mockOrderRepo)
//
//	r := gin.Default()
//	r.POST("/orders", mockService.CreateOrder)
//
//	mockOrder := models.Order{
//		Code: "ord3",
//		OrderItems: []models.OrderItem{
//			{
//				ItemId:   5,
//				Quantity: 50,
//			},
//		},
//	}
//
//	stringMockOrder, err := json.Marshal(mockOrder)
//	if err != nil {
//		t.Errorf("Error marshalling order: %v", err)
//	}
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(stringMockOrder))
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusOK, w.Code)
//	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
//
//	var order models.Order
//	if err := json.Unmarshal(w.Body.Bytes(), &order); err != nil {
//		t.Errorf("Error unmarshalling order: %v", err)
//	}
//
//	assert.Equal(t, mockOrder, order)
//}
//
//// TestCreateOrder_BindError test the CreateOrder function using mockOrderRepo and gin
//func TestCreateOrder_BindError(t *testing.T) {
//	mockOrderRepo := newMockOrderRepo()
//	mockService := NewOrderService(mockOrderRepo)
//
//	r := gin.Default()
//	r.POST("/orders", mockService.CreateOrder)
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("POST", "/orders", nil)
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusBadRequest, w.Code)
//	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
//
//	var response map[string]string
//	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
//		t.Errorf("Error unmarshalling response: %v", err)
//	}
//
//	assert.NotNil(t, response["error"])
//}
//
//// TestCreateOrder_SaveError test the CreateOrder function using mockOrderRepo and gin
//func TestCreateOrder_SaveError(t *testing.T) {
//	mockOrderRepo := newMockOrderErrorRepo()
//	mockService := NewOrderService(mockOrderRepo)
//
//	r := gin.Default()
//	r.POST("/orders", mockService.CreateOrder)
//
//	mockOrder := models.Order{
//		Code: "ord3",
//		OrderItems: []models.OrderItem{
//			{
//				ItemId:   5,
//				Quantity: 50,
//			},
//		},
//	}
//
//	stringMockOrder, err := json.Marshal(mockOrder)
//	if err != nil {
//		t.Errorf("Error marshalling order: %v", err)
//	}
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(stringMockOrder))
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusInternalServerError, w.Code)
//	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
//
//	var response map[string]string
//	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
//		t.Errorf("Error unmarshalling response: %v", err)
//	}
//
//	assert.NotNil(t, response["error"])
//}
//
//// TestGetAllOrders test the GetAllOrders function using mockOrderRepo and gin
//func TestGetAllOrders(t *testing.T) {
//	mockOrderRepo := newMockOrderRepo()
//	mockService := NewOrderService(mockOrderRepo)
//
//	r := gin.Default()
//	r.GET("/orders", mockService.GetAllOrders)
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("GET", "/orders", nil)
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusOK, w.Code)
//	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
//
//	var orders []models.Order
//	if err := json.Unmarshal(w.Body.Bytes(), &orders); err != nil {
//		t.Errorf("Error unmarshalling orders: %v", err)
//	}
//
//	assert.Equal(t, mockOrders, orders)
//}
//
//// TestGetAllOrders_FindAllError test the GetAllOrders function using mockOrderRepo and gin
//func TestGetAllOrders_FindAllError(t *testing.T) {
//	mockOrderRepo := newMockOrderErrorRepo()
//	mockService := NewOrderService(mockOrderRepo)
//
//	r := gin.Default()
//	r.GET("/orders", mockService.GetAllOrders)
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("GET", "/orders", nil)
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusInternalServerError, w.Code)
//	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
//
//	var response map[string]string
//	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
//		t.Errorf("Error unmarshalling response: %v", err)
//	}
//
//	assert.NotNil(t, response["error"])
//}
//
//// TestGetOrder test the GetOrder function using mockOrderRepo and gin
//func TestGetOrder(t *testing.T) {
//	mockOrderRepo := newMockOrderRepo()
//	mockService := NewOrderService(mockOrderRepo)
//
//	r := gin.Default()
//	r.GET("/orders/:id", mockService.GetOrder)
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("GET", "/orders/1", nil)
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusOK, w.Code)
//	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
//
//	var order models.Order
//	if err := json.Unmarshal(w.Body.Bytes(), &order); err != nil {
//		t.Errorf("Error unmarshalling order: %v", err)
//	}
//
//	assert.Equal(t, mockOrders[0], order)
//}
//
//// TestGetOrder_InvalidID test the GetOrder function using mockOrderRepo and gin
//func TestGetOrder_InvalidID(t *testing.T) {
//	mockOrderRepo := newMockOrderRepo()
//	mockService := NewOrderService(mockOrderRepo)
//
//	r := gin.Default()
//	r.GET("/orders/:id", mockService.GetOrder)
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("GET", "/orders/abc", nil)
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusBadRequest, w.Code)
//	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
//
//	var response map[string]string
//	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
//		t.Errorf("Error unmarshalling response: %v", err)
//	}
//
//	assert.NotNil(t, response["error"])
//}
//
//// TestGetOrder_FindError test the GetOrder function using mockOrderRepo and gin
//func TestGetOrder_FindError(t *testing.T) {
//	mockOrderRepo := newMockOrderErrorRepo()
//	mockService := NewOrderService(mockOrderRepo)
//
//	r := gin.Default()
//	r.GET("/orders/:id", mockService.GetOrder)
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("GET", "/orders/1", nil)
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusInternalServerError, w.Code)
//	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
//
//	var response map[string]string
//	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
//		t.Errorf("Error unmarshalling response: %v", err)
//	}
//
//	assert.NotNil(t, response["error"])
//}
//
//// TestUpdateOrder test the UpdateOrder function using mockOrderRepo and gin
//func TestUpdateOrder(t *testing.T) {
//	mockOrderRepo := newMockOrderRepo()
//	mockService := NewOrderService(mockOrderRepo)
//
//	r := gin.Default()
//	r.PUT("/orders/:id", mockService.UpdateOrder)
//
//	mockOrder := models.Order{
//		Code: "ord3",
//		OrderItems: []models.OrderItem{
//			{
//				ItemId:   5,
//				Quantity: 50,
//			},
//		},
//	}
//
//	stringMockOrder, err := json.Marshal(mockOrder)
//	if err != nil {
//		t.Errorf("Error marshalling order: %v", err)
//	}
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("PUT", "/orders/1", bytes.NewBuffer(stringMockOrder))
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusOK, w.Code)
//	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
//
//	var order models.Order
//	if err := json.Unmarshal(w.Body.Bytes(), &order); err != nil {
//		t.Errorf("Error unmarshalling order: %v", err)
//	}
//
//	assert.Equal(t, mockOrder, order)
//}
//
//// TestUpdateOrder_InvalidID test the UpdateOrder function using mockOrderRepo and gin
//func TestUpdateOrder_InvalidID(t *testing.T) {
//	mockOrderRepo := newMockOrderRepo()
//	mockService := NewOrderService(mockOrderRepo)
//
//	r := gin.Default()
//	r.PUT("/orders/:id", mockService.UpdateOrder)
//
//	mockOrder := models.Order{
//		Code: "ord3",
//		OrderItems: []models.OrderItem{
//			{
//				ItemId:   5,
//				Quantity: 50,
//			},
//		},
//	}
//
//	stringMockOrder, err := json.Marshal(mockOrder)
//	if err != nil {
//		t.Errorf("Error marshalling order: %v", err)
//	}
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("PUT", "/orders/abc", bytes.NewBuffer(stringMockOrder))
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusBadRequest, w.Code)
//	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
//
//	var response map[string]string
//	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
//		t.Errorf("Error unmarshalling response: %v", err)
//	}
//
//	assert.NotNil(t, response["error"])
//}
//
//// TestUpdateOrder_FindError test the UpdateOrder function using mockOrderRepo and gin
//func TestUpdateOrder_FindError(t *testing.T) {
//	mockOrderRepo := newMockOrderErrorRepo()
//	mockService := NewOrderService(mockOrderRepo)
//
//	r := gin.Default()
//	r.PUT("/orders/:id", mockService.UpdateOrder)
//
//	mockOrder := models.Order{
//		Code: "ord3",
//		OrderItems: []models.OrderItem{
//			{
//				ItemId:   5,
//				Quantity: 50,
//			},
//		},
//	}
//
//	stringMockOrder, err := json.Marshal(mockOrder)
//	if err != nil {
//		t.Errorf("Error marshalling order: %v", err)
//	}
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("PUT", "/orders/1", bytes.NewBuffer(stringMockOrder))
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusNotFound, w.Code)
//	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
//
//	var response map[string]string
//	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
//		t.Errorf("Error unmarshalling response: %v", err)
//	}
//
//	assert.NotNil(t, response["error"])
//}
//
//// TestUpdateOrder_BindError test the UpdateOrder function using mockOrderRepo and gin
//func TestUpdateOrder_BindError(t *testing.T) {
//	mockOrderRepo := newMockOrderRepo()
//	mockService := NewOrderService(mockOrderRepo)
//
//	r := gin.Default()
//	r.PUT("/orders/:id", mockService.UpdateOrder)
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("PUT", "/orders/1", bytes.NewBuffer([]byte("")))
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusBadRequest, w.Code)
//	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
//
//	var response map[string]string
//	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
//		t.Errorf("Error unmarshalling response: %v", err)
//	}
//
//	assert.NotNil(t, response["error"])
//}
//
//// TestUpdateOrder_UpdateError test the UpdateOrder function using mockOrderRepo and gin
//func TestUpdateOrder_UpdateError(t *testing.T) {
//	mockOrderRepo := newMockOrderSpecificErrorRepo()
//	mockService := NewOrderService(mockOrderRepo)
//
//	r := gin.Default()
//	r.PUT("/orders/:id", mockService.UpdateOrder)
//
//	mockOrder := models.Order{
//		Code: "ord3",
//		OrderItems: []models.OrderItem{
//			{
//				ItemId:   5,
//				Quantity: 50,
//			},
//		},
//	}
//
//	stringMockOrder, err := json.Marshal(mockOrder)
//	if err != nil {
//		t.Errorf("Error marshalling order: %v", err)
//	}
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("PUT", "/orders/1", bytes.NewBuffer(stringMockOrder))
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusInternalServerError, w.Code)
//	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
//
//	var response map[string]string
//	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
//		t.Errorf("Error unmarshalling response: %v", err)
//	}
//
//	assert.NotNil(t, response["error"])
//}
//
//// TestDeleteOrder test the DeleteOrder function using mockOrderRepo and gin
//func TestDeleteOrder(t *testing.T) {
//	mockOrderRepo := newMockOrderRepo()
//	mockService := NewOrderService(mockOrderRepo)
//
//	r := gin.Default()
//	r.DELETE("/orders/:id", mockService.DeleteOrder)
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("DELETE", "/orders/1", nil)
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusOK, w.Code)
//	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
//
//	var order models.Order
//	if err := json.Unmarshal(w.Body.Bytes(), &order); err != nil {
//		t.Errorf("Error unmarshalling order: %v", err)
//	}
//
//	assert.Equal(t, mockOrders[0], order)
//}
//
//// TestDeleteOrder_InvalidID test the DeleteOrder function using mockOrderRepo and gin
//func TestDeleteOrder_InvalidID(t *testing.T) {
//	mockOrderRepo := newMockOrderRepo()
//	mockService := NewOrderService(mockOrderRepo)
//
//	r := gin.Default()
//	r.DELETE("/orders/:id", mockService.DeleteOrder)
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("DELETE", "/orders/abc", nil)
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusBadRequest, w.Code)
//	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
//
//	var response map[string]string
//	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
//		t.Errorf("Error unmarshalling response: %v", err)
//	}
//
//	assert.NotNil(t, response["error"])
//}
//
//// TestDeleteOrder_FindError test the DeleteOrder function using mockOrderRepo and gin
//func TestDeleteOrder_FindError(t *testing.T) {
//	mockOrderRepo := newMockOrderErrorRepo()
//	mockService := NewOrderService(mockOrderRepo)
//
//	r := gin.Default()
//	r.DELETE("/orders/:id", mockService.DeleteOrder)
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("DELETE", "/orders/1", nil)
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusNotFound, w.Code)
//	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
//
//	var response map[string]string
//	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
//		t.Errorf("Error unmarshalling response: %v", err)
//	}
//
//	assert.NotNil(t, response["error"])
//}
//
//// TestDeleteOrder_DeleteError test the DeleteOrder function using mockOrderRepo and gin
//func TestDeleteOrder_DeleteError(t *testing.T) {
//	mockOrderRepo := newMockOrderSpecificErrorRepo()
//	mockService := NewOrderService(mockOrderRepo)
//
//	r := gin.Default()
//	r.DELETE("/orders/:id", mockService.DeleteOrder)
//
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("DELETE", "/orders/1", nil)
//	r.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusInternalServerError, w.Code)
//	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
//
//	var response map[string]string
//	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
//		t.Errorf("Error unmarshalling response: %v", err)
//	}
//
//	assert.NotNil(t, response["error"])
//}
