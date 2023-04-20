package services

import (
	"github.com/laertkokona/crud-test/models"
	"github.com/laertkokona/crud-test/repositories"
	"github.com/laertkokona/crud-test/utils"
	"net/http"
)

// OrderService interface using gin context
type OrderService interface {
	CreateOrder(order models.Order) (models.Order, int, error)
	GetOrder(id int) (models.Order, int, error)
	GetAllOrders(pagination models.Pagination) ([]models.Order, int, error)
	UpdateOrder(id int, order models.Order) (models.Order, int, error)
	DeleteOrder(id int) (models.Order, int, error)
}

// orderService struct
type orderService struct {
	OrderRepo repositories.OrderRepo
}

// NewOrderService returns a new instance of orderService
func NewOrderService(orderRepo repositories.OrderRepo) OrderService {
	return orderService{
		OrderRepo: orderRepo,
	}
}

// CreateOrder method that takes a models.Order object and saves it to the database
func (p orderService) CreateOrder(order models.Order) (models.Order, int, error) {
	// call the order repository to save the order
	// return the order object
	order, err := p.OrderRepo.Save(order)
	if err != nil {
		return order, http.StatusInternalServerError, err
	}
	return order, http.StatusOK, nil
}

// GetOrder method that takes an order id and returns the order object
func (p orderService) GetOrder(id int) (models.Order, int, error) {
	// call the order repository to get the order
	// return the order object
	order, err := p.OrderRepo.FindByID(id)
	if err != nil {
		return order, http.StatusInternalServerError, err
	}
	return order, http.StatusOK, nil
}

// GetAllOrders method that returns all the orders
func (p orderService) GetAllOrders(pagination models.Pagination) ([]models.Order, int, error) {
	// call the order repository to get all the orders
	// return the orders
	orders, err := p.OrderRepo.FindAll(pagination)
	if err != nil {
		return orders, http.StatusInternalServerError, err
	}
	return orders, http.StatusOK, nil
}

// UpdateOrder method that takes an order id and a models.Order object and updates the order
func (p orderService) UpdateOrder(id int, order models.Order) (models.Order, int, error) {
	// call the order repository to update the order
	// return the order object
	orderDb, err := p.OrderRepo.FindByID(id)
	if err != nil {
		return orderDb, http.StatusNotFound, err
	}
	//var comparableOrderDb models.ComparableOrder
	//var comparableOrder models.ComparableOrder
	//comparableOrderDb = models.ComparableOrder{}

	utils.CopyNonEmptyFields(&orderDb, &order)
	orderDb, err = p.OrderRepo.Update(orderDb)
	if err != nil {
		return orderDb, http.StatusInternalServerError, err
	}
	return orderDb, http.StatusOK, nil
}

// DeleteOrder method that takes an order id and deletes the order
func (p orderService) DeleteOrder(id int) (models.Order, int, error) {
	item, err := p.OrderRepo.FindByID(id)
	if err != nil {
		return item, http.StatusNotFound, err
	}
	err = p.OrderRepo.Delete(item)
	if err != nil {
		return item, http.StatusInternalServerError, err
	}
	return item, http.StatusOK, nil
}
