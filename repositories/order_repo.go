package repositories

import (
	"github.com/laertkokona/crud-test/models"
	"gorm.io/gorm"
)

// OrderRepo interface
type OrderRepo interface {
	FindAll(pagination models.Pagination) ([]models.Order, error)
	FindByID(int) (models.Order, error)
	Save(models.Order) (models.Order, error)
	Update(models.Order) (models.Order, error)
	Delete(models.Order) error
	DeleteById(int) (models.Order, error)
}

// orderRepo struct
type orderRepo struct {
	DB *gorm.DB
}

// NewOrderRepo returns a new instance of orderRepo
func NewOrderRepo(db *gorm.DB) OrderRepo {
	return orderRepo{
		DB: db,
	}
}

// FindAll returns all orders
func (o orderRepo) FindAll(pagination models.Pagination) ([]models.Order, error) {
	// If pagination is not set, return all orders
	// If pagination is set, return orders based on pagination
	var orders []models.Order
	if pagination.Limit == 0 || pagination.Page == 0 {
		//if err := o.DB.Preload("OrderItems").Find(&orders).Error; err != nil {
		//	return nil, err
		//}
		return orders, o.DB.Preload("OrderItems").Find(&orders).Error
	}
	//if err := o.DB.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Preload("OrderItems").Find(&orders).Error; err != nil {
	//	return nil, err
	//}
	return orders, o.DB.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Preload("OrderItems").Find(&orders).Error
}

// FindByID returns an order by id
func (o orderRepo) FindByID(id int) (models.Order, error) {
	var order models.Order
	if err := o.DB.Preload("OrderItems").First(&order, id).Error; err != nil {
		return order, err
	}
	return order, o.DB.Preload("OrderItems").First(&order, id).Error
}

// Save saves an order
func (o orderRepo) Save(order models.Order) (models.Order, error) {
	return order, o.DB.Create(&order).Error
}

// Update updates an order
func (o orderRepo) Update(order models.Order) (models.Order, error) {
	return order, o.DB.Save(&order).Error
}

// Delete deletes an order
func (o orderRepo) Delete(order models.Order) error {
	return o.DB.Delete(&order).Error
}

// DeleteById deletes an order by id
func (o orderRepo) DeleteById(id int) (models.Order, error) {
	var order models.Order
	if err := o.DB.First(&order, id).Error; err != nil {
		return order, err
	}
	return order, o.DB.Delete(&order).Error
}
