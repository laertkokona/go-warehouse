package repositories

import (
	"github.com/laertkokona/crud-test/models"
	"gorm.io/gorm"
)

// TruckRepo interface
type TruckRepo interface {
	FindAll(pagination models.Pagination) ([]models.Truck, error)
	FindByID(int) (models.Truck, error)
	Save(models.Truck) (models.Truck, error)
	Update(models.Truck) (models.Truck, error)
	Delete(models.Truck) error
	DeleteById(int) (models.Truck, error)
}

// truckRepo struct
type truckRepo struct {
	DB *gorm.DB
}

// NewTruckRepo returns a new instance of truckRepo
func NewTruckRepo(db *gorm.DB) TruckRepo {
	return truckRepo{
		DB: db,
	}
}

// FindAll returns all trucks
func (t truckRepo) FindAll(pagination models.Pagination) ([]models.Truck, error) {
	// If pagination is not set, return all trucks
	// If pagination is set, return trucks based on pagination
	var trucks []models.Truck
	if pagination.Limit == 0 || pagination.Page == 0 {
		//if err := t.DB.Find(&trucks).Error; err != nil {
		//	return nil, err
		//}
		return trucks, t.DB.Find(&trucks).Error
	}
	//if err := t.DB.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&trucks).Error; err != nil {
	//	return nil, err
	//}
	return trucks, t.DB.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&trucks).Error
}

// FindByID returns a truck by id
func (t truckRepo) FindByID(id int) (models.Truck, error) {
	var truck models.Truck
	if err := t.DB.First(&truck, id).Error; err != nil {
		return truck, err
	}
	return truck, t.DB.First(&truck, id).Error
}

// Save saves a truck
func (t truckRepo) Save(truck models.Truck) (models.Truck, error) {
	return truck, t.DB.Create(&truck).Error
}

// Update updates a truck
func (t truckRepo) Update(truck models.Truck) (models.Truck, error) {
	return truck, t.DB.Save(&truck).Error
}

// Delete deletes a truck
func (t truckRepo) Delete(truck models.Truck) error {
	return t.DB.Delete(&truck).Error
}

// DeleteById deletes a truck by id
func (t truckRepo) DeleteById(id int) (models.Truck, error) {
	var truck models.Truck
	if err := t.DB.First(&truck, id).Error; err != nil {
		return truck, err
	}
	return truck, t.DB.Delete(&truck).Error
}
