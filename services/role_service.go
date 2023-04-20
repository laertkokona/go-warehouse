package services

import (
	"github.com/laertkokona/crud-test/models"
	"github.com/laertkokona/crud-test/repositories"
	"github.com/laertkokona/crud-test/utils"
	"github.com/peteprogrammer/go-automapper"
	"net/http"
)

// RoleService interface like UserService
type RoleService interface {
	CreateRole(role models.Role) (models.RoleDTO, int, error)
	GetRole(id int) (models.RoleDTO, int, error)
	GetAllRoles() ([]models.RoleDTO, int, error)
	UpdateRole(id int, role models.Role) (models.RoleDTO, int, error)
	DeleteRole(id int) (models.RoleDTO, int, error)
}

// roleService struct like userService
type roleService struct {
	roleRepo repositories.RoleRepo
}

// NewRoleService returns a new instance of RoleService
func NewRoleService(repo repositories.RoleRepo) RoleService {
	return roleService{
		roleRepo: repo,
	}
}

// CreateRole method that takes a models.Role object and saves it to the database
func (r roleService) CreateRole(role models.Role) (models.RoleDTO, int, error) {
	// save the role to the database
	// return the role object
	role, err := r.roleRepo.Save(role)
	if err != nil {
		return models.RoleDTO{}, http.StatusInternalServerError, err
	}
	var returnRole models.RoleDTO
	automapper.Map(role, &returnRole)
	return returnRole, http.StatusOK, nil
}

// GetRole method that takes a role id and returns the role object
func (r roleService) GetRole(id int) (models.RoleDTO, int, error) {
	// get the role object from the database
	// return the role object
	role, err := r.roleRepo.FindByID(id)
	if err != nil {
		return models.RoleDTO{}, http.StatusInternalServerError, err
	}
	var returnRole models.RoleDTO
	automapper.Map(role, &returnRole)
	return returnRole, http.StatusOK, nil
}

// GetAllRoles method that returns all the roles
func (r roleService) GetAllRoles() ([]models.RoleDTO, int, error) {
	// get all the roles from the database
	// return the roles
	roles, err := r.roleRepo.FindAll()
	if err != nil {
		return []models.RoleDTO{}, http.StatusInternalServerError, err
	}
	var returnRoles []models.RoleDTO
	automapper.Map(roles, &returnRoles)
	return returnRoles, http.StatusOK, nil
}

// UpdateRole method that takes a role id and updates the role object
func (r roleService) UpdateRole(id int, role models.Role) (models.RoleDTO, int, error) {
	// get the role object from the database
	// unmarshal the role object
	// save the role to the database
	// return the role object
	roleDB, err := r.roleRepo.FindByID(id)
	if err != nil {
		return models.RoleDTO{}, http.StatusNotFound, err
	}
	utils.CopyNonEmptyFields(&roleDB, &role)
	role, err = r.roleRepo.Update(role)
	if err != nil {
		return models.RoleDTO{}, http.StatusInternalServerError, err
	}
	var returnRole models.RoleDTO
	automapper.Map(role, &returnRole)
	return returnRole, http.StatusOK, nil
}

// DeleteRole method that takes a role id and deletes the role object
func (r roleService) DeleteRole(id int) (models.RoleDTO, int, error) {
	// get the role object from the database
	// delete the role object from the database
	// return the role object
	role, err := r.roleRepo.FindByID(id)
	if err != nil {
		return models.RoleDTO{}, http.StatusNotFound, err
	}
	err = r.roleRepo.Delete(role)
	if err != nil {
		return models.RoleDTO{}, http.StatusInternalServerError, err
	}
	var returnRole models.RoleDTO
	automapper.Map(role, &returnRole)
	return returnRole, http.StatusOK, nil
}
