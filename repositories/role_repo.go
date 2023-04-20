package repositories

import (
	"github.com/laertkokona/crud-test/models"
	"gorm.io/gorm"
)

// roleRepo struct
type roleRepo struct {
	DB *gorm.DB
}

// RoleRepo interface
type RoleRepo interface {
	FindAll() ([]models.Role, error)
	FindByID(int) (models.Role, error)
	FindByName(string) (models.Role, error)
	Save(models.Role) (models.Role, error)
	Update(models.Role) (models.Role, error)
	Delete(models.Role) error
	DeleteById(int) (models.Role, error)
}

// NewRoleRepo returns a new instance of roleRepo
func NewRoleRepo(db *gorm.DB) RoleRepo {
	return roleRepo{
		DB: db,
	}
}

// FindAll returns all roles
func (r roleRepo) FindAll() ([]models.Role, error) {
	var roles []models.Role
	if err := r.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, r.DB.Find(&roles).Error
}

// FindByID returns a role by id
func (r roleRepo) FindByID(id int) (models.Role, error) {
	var role models.Role
	//if err := r.DB.Preload("Users").First(&role, id).Error; err != nil {
	//	return role, err
	//}
	return role, r.DB.First(&role, id).Error
}

// FindByName returns a role by name
func (r roleRepo) FindByName(name string) (models.Role, error) {
	var role models.Role
	if err := r.DB.First(&role, "name=?", name).Error; err != nil {
		return role, err
	}
	return role, r.DB.First(&role, "name=?", name).Error
}

// Save saves a role
func (r roleRepo) Save(role models.Role) (models.Role, error) {
	return role, r.DB.Create(&role).Error
}

// Update updates a role
func (r roleRepo) Update(role models.Role) (models.Role, error) {
	return role, r.DB.Save(&role).Error
}

// Delete deletes a role
func (r roleRepo) Delete(role models.Role) error {
	return r.DB.Delete(&role).Error
}

// DeleteById deletes a role by id
func (r roleRepo) DeleteById(id int) (models.Role, error) {
	var role models.Role
	if err := r.DB.First(&role, id).Error; err != nil {
		return role, err
	}
	return role, r.DB.Delete(&models.Role{}, id).Error
}
