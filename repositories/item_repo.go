package repositories

import (
	"github.com/laertkokona/crud-test/models"
	"gorm.io/gorm"
)

// itemRepo struct
type itemRepo struct {
	DB *gorm.DB
}

// ItemRepo interface for item repository
type ItemRepo interface {
	FindAll(pagination models.Pagination) ([]models.Item, error)
	FindByID(int) (models.Item, error)
	FindByName(string) (models.Item, error)
	Save(models.Item) (models.Item, error)
	Update(models.Item) (models.Item, error)
	Delete(models.Item) error
	DeleteById(int) (models.Item, error)
}

// NewItemRepo returns a new instance of itemRepo
func NewItemRepo(db *gorm.DB) ItemRepo {
	return itemRepo{
		DB: db,
	}
}

// FindAll returns all items
func (p itemRepo) FindAll(pagination models.Pagination) ([]models.Item, error) {
	// If pagination is not set, return all items
	// If pagination is set, return items based on pagination
	var items []models.Item
	if pagination.Limit == 0 || pagination.Page == 0 {
		if err := p.DB.Find(&items).Error; err != nil {
			return nil, err
		}
		return items, p.DB.Find(&items).Error
	}
	if err := p.DB.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, p.DB.Offset((pagination.Page - 1) * pagination.Limit).Limit(pagination.Limit).Find(&items).Error
}

// FindByID returns an item by id
func (p itemRepo) FindByID(id int) (models.Item, error) {
	var item models.Item
	if err := p.DB.First(&item, id).Error; err != nil {
		return item, err
	}
	return item, p.DB.First(&item, id).Error
}

// FindByName returns an item by name
func (p itemRepo) FindByName(name string) (models.Item, error) {
	var item models.Item
	if err := p.DB.First(&item, "name=?", name).Error; err != nil {
		return item, err
	}
	return item, p.DB.First(&item, "name=?", name).Error
}

// Save saves an item
func (p itemRepo) Save(item models.Item) (models.Item, error) {
	return item, p.DB.Create(&item).Error
}

// Update updates an item
func (p itemRepo) Update(item models.Item) (models.Item, error) {
	return item, p.DB.Save(&item).Error
}

// Delete deletes an item
func (p itemRepo) Delete(item models.Item) error {
	return p.DB.Delete(&item).Error
}

// DeleteById deletes an item by id
func (p itemRepo) DeleteById(id int) (models.Item, error) {
	var item models.Item
	if err := p.DB.First(&item, id).Error; err != nil {
		return item, err
	}
	return item, p.DB.Delete(&item).Error
}
