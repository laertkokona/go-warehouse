package services

import (
	"github.com/laertkokona/crud-test/models"
	"github.com/laertkokona/crud-test/repositories"
	"github.com/laertkokona/crud-test/utils"
	"github.com/peteprogrammer/go-automapper"
	"net/http"
)

// ItemService interface with gin services
type ItemService interface {
	CreateItem(item models.Item) (models.ItemDTO, int, error)
	GetItem(id int) (models.ItemDTO, int, error)
	GetAllItems(pagination models.Pagination) ([]models.ItemDTO, int, error)
	UpdateItem(id int, item models.Item) (models.ItemDTO, int, error)
	DeleteItem(id int) (models.ItemDTO, int, error)
}

// itemService struct
type itemService struct {
	ItemRepo repositories.ItemRepo
}

// NewItemService returns a new instance of itemService
func NewItemService(itemRepo repositories.ItemRepo) ItemService {
	return itemService{
		ItemRepo: itemRepo,
	}
}

// CreateItem method that takes a models.Item object and saves it to the database
func (p itemService) CreateItem(item models.Item) (models.ItemDTO, int, error) {
	item, err := p.ItemRepo.Save(item)
	if err != nil {
		return models.ItemDTO{}, http.StatusInternalServerError, err
	}
	var itemDTO models.ItemDTO
	automapper.Map(item, &itemDTO)
	return itemDTO, http.StatusOK, nil
}

// GetItem method that takes an item id and returns the item object
func (p itemService) GetItem(id int) (models.ItemDTO, int, error) {
	item, err := p.ItemRepo.FindByID(id)
	if err != nil {
		return models.ItemDTO{}, http.StatusNotFound, err
	}
	var itemDTO models.ItemDTO
	automapper.Map(item, &itemDTO)
	return itemDTO, http.StatusOK, nil
}

// GetAllItems method that returns all items
func (p itemService) GetAllItems(pagination models.Pagination) ([]models.ItemDTO, int, error) {
	items, err := p.ItemRepo.FindAll(pagination)
	if err != nil {
		return []models.ItemDTO{}, http.StatusInternalServerError, err
	}
	var itemsDTO []models.ItemDTO
	automapper.Map(items, &itemsDTO)
	return itemsDTO, http.StatusOK, nil
}

// UpdateItem method that takes an item id and an item object and updates the item
func (p itemService) UpdateItem(id int, item models.Item) (models.ItemDTO, int, error) {

	itemDb, err := p.ItemRepo.FindByID(id)
	if err != nil {
		return models.ItemDTO{}, http.StatusNotFound, err
	}
	//err = json.Unmarshal([]byte(itemString), &itemDb)
	// set all the fields of the item that are not empty to the itemDb
	utils.CopyNonEmptyFields(&itemDb, &item)

	itemDb, err = p.ItemRepo.Update(itemDb)
	if err != nil {
		return models.ItemDTO{}, http.StatusInternalServerError, err
	}
	var itemDTO models.ItemDTO
	automapper.Map(itemDb, &itemDTO)
	return itemDTO, http.StatusOK, nil
}

// DeleteItem method that takes an item id and deletes the item
func (p itemService) DeleteItem(id int) (models.ItemDTO, int, error) {
	item, err := p.ItemRepo.FindByID(id)
	if err != nil {
		return models.ItemDTO{}, http.StatusNotFound, err
	}
	err = p.ItemRepo.Delete(item)
	if err != nil {
		return models.ItemDTO{}, http.StatusInternalServerError, err
	}
	var itemDTO models.ItemDTO
	automapper.Map(item, &itemDTO)
	return itemDTO, http.StatusOK, nil
}
