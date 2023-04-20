package services

import (
	"errors"
	"github.com/laertkokona/crud-test/models"
	"github.com/peteprogrammer/go-automapper"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

var mockRoles = []models.Role{
	{
		ID:   1,
		Name: "User",
	},
	{
		ID:   2,
		Name: "Admin",
	},
	{
		ID:   3,
		Name: "SysAdmin",
	},
}

var mockCreateRole = models.Role{
	ID:   4,
	Name: "Test",
}

// mockRoleRepo is a mock implementation of the repositories.RoleRepo interface
type mockRoleRepo struct {
	// findAll is a mock function with no given fields
	findAll func() ([]models.Role, error)
	// findByID is a mock function with given fields: id
	findByID func(id int) (models.Role, error)
	// findByName is a mock function with given fields: roleName
	findByName func(roleName string) (models.Role, error)
	// save is a mock function with given fields: role
	save func(role models.Role) (models.Role, error)
	// update is a mock function with given fields: role
	update func(role models.Role) (models.Role, error)
	// delete is a mock function with given fields: role
	delete func(role models.Role) error
	// deleteById is a mock function with given fields: id
	deleteById func(id int) (models.Role, error)
}

// FindAll is a mock function with given fields: pagination
func (_m *mockRoleRepo) FindAll() ([]models.Role, error) {
	return _m.findAll()
}

// FindByID is a mock function with given fields: id
func (_m *mockRoleRepo) FindByID(id int) (models.Role, error) {
	return _m.findByID(id)
}

// FindByName is a mock function with given fields: roleName
func (_m *mockRoleRepo) FindByName(roleName string) (models.Role, error) {
	return _m.findByName(roleName)
}

// Save is a mock function with given fields: role
func (_m *mockRoleRepo) Save(role models.Role) (models.Role, error) {
	return _m.save(role)
}

// Update is a mock function with given fields: role
func (_m *mockRoleRepo) Update(role models.Role) (models.Role, error) {
	log.Println("Update called:", role)
	return _m.update(role)
}

// Delete is a mock function with given fields: role
func (_m *mockRoleRepo) Delete(role models.Role) error {
	return _m.delete(role)
}

// DeleteById is a mock function with given fields: id
func (_m *mockRoleRepo) DeleteById(id int) (models.Role, error) {
	return _m.deleteById(id)
}

// NewMockRoleRepo returns a new instance of mockRoleRepo
func NewMockRoleRepo() *mockRoleRepo {
	return &mockRoleRepo{
		findAll: func() ([]models.Role, error) {
			return mockRoles, nil
		},
		findByID: func(id int) (models.Role, error) {
			return mockRoles[id-1], nil
		},
		findByName: func(name string) (models.Role, error) {
			var role models.Role
			for _, r := range mockRoles {
				if r.Name == name {
					role = r
					break
				}
			}
			return role, nil
		},
		save: func(role models.Role) (models.Role, error) {
			return role, nil
		},
		update: func(role models.Role) (models.Role, error) {
			role.ID = uint(1)
			return role, nil
		},
		delete: func(role models.Role) error {
			return nil
		},
		deleteById: func(id int) (models.Role, error) {
			return mockRoles[id-1], nil
		},
	}
}

// NewMockRoleErrorRepo returns a new instance of mockRoleRepo with error
func NewMockRoleErrorRepo() *mockRoleRepo {
	return &mockRoleRepo{
		findAll: func() ([]models.Role, error) {
			return []models.Role{}, errors.New("error")
		},
		findByID: func(id int) (models.Role, error) {
			return models.Role{}, errors.New("error")
		},
		findByName: func(name string) (models.Role, error) {
			return models.Role{}, errors.New("error")
		},
		save: func(role models.Role) (models.Role, error) {
			return models.Role{}, errors.New("error")
		},
		update: func(role models.Role) (models.Role, error) {
			return models.Role{}, errors.New("error")
		},
		delete: func(role models.Role) error {
			return errors.New("error")
		},
		deleteById: func(id int) (models.Role, error) {
			return models.Role{}, errors.New("error")
		},
	}
}

// NewMockRoleErrorRepo returns a new instance of mockRoleRepo with error
func NewMockRoleSpecificErrorRepo() *mockRoleRepo {
	return &mockRoleRepo{
		findAll: func() ([]models.Role, error) {
			return mockRoles, nil
		},
		findByID: func(id int) (models.Role, error) {
			return mockRoles[id-1], nil
		},
		findByName: func(name string) (models.Role, error) {
			var role models.Role
			for _, r := range mockRoles {
				if r.Name == name {
					role = r
					break
				}
			}
			return role, nil
		},
		save: func(role models.Role) (models.Role, error) {
			return models.Role{}, errors.New("error")
		},
		update: func(role models.Role) (models.Role, error) {
			return models.Role{}, errors.New("error")
		},
		delete: func(role models.Role) error {
			return errors.New("error")
		},
		deleteById: func(id int) (models.Role, error) {
			return models.Role{}, errors.New("error")
		},
	}
}

// TestNewRoleService tests the NewRoleService function
func TestNewRoleService(t *testing.T) {
	mockRepo := NewMockRoleRepo()
	mockService := NewRoleService(mockRepo)

	assert.NotNil(t, mockService)
	assert.IsType(t, roleService{}, mockService)
}

// TestCreateRole tests the CreateRole function using mockRoleRepo
func TestCreateRole(t *testing.T) {
	// create a mock role repo
	// create a mock role service
	// call the create role function
	// assert that the role returned is the same as the role passed in
	// assert that the status is OK
	// assert that there is no error
	mockRepo := NewMockRoleRepo()
	mockService := NewRoleService(mockRepo)
	role, status, err := mockService.CreateRole(mockCreateRole)
	var expectedDTO models.RoleDTO
	automapper.Map(mockCreateRole, &expectedDTO)
	assert.NoError(t, err, "Error creating role")
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, expectedDTO, role)
}

// TestCreateRole_SaveError tests the CreateRole function using mockRoleRepo
func TestCreateRole_SaveError(t *testing.T) {
	// create a mock role repo
	// create a mock role service
	// call the create role function
	// assert that the role returned is the same as the role passed in
	// assert that the status is InternalServerError
	// assert that there is no error
	mockRepo := NewMockRoleErrorRepo()
	mockService := NewRoleService(mockRepo)
	role, status, err := mockService.CreateRole(mockCreateRole)
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, models.RoleDTO{}, role)
}

// TestGetRole tests the GetRole function using mockRoleRepo
func TestGetRole(t *testing.T) {
	// create a mock role repo
	// create a mock role service
	// call the get role function
	// assert that the role returned is the same as the role passed in
	// assert that the status is OK
	// assert that there is no error
	mockRepo := NewMockRoleRepo()
	mockService := NewRoleService(mockRepo)
	role, status, err := mockService.GetRole(1)
	var expectedDTO models.RoleDTO
	automapper.Map(mockRoles[0], &expectedDTO)
	assert.NoError(t, err, "Error getting role")
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, expectedDTO, role)
}

// TestGetRole_FindByIDError tests the GetRole function using mockRoleRepo
func TestGetRole_FindByIDError(t *testing.T) {
	// create a mock role repo
	// create a mock role service
	// call the get role function
	// assert that the role returned is the same as the role passed in
	// assert that the status is InternalServerError
	// assert that there is no error
	mockRepo := NewMockRoleErrorRepo()
	mockService := NewRoleService(mockRepo)
	role, status, err := mockService.GetRole(1)
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, models.RoleDTO{}, role)
}

// TestGetAllRoles tests the GetAllRoles function using mockRoleRepo
func TestGetAllRoles(t *testing.T) {
	// create a mock role repo
	// create a mock role service
	// call the get all roles function
	// assert that the roles returned is the same as the roles passed in
	// assert that the status is OK
	// assert that there is no error
	mockRepo := NewMockRoleRepo()
	mockService := NewRoleService(mockRepo)
	roles, status, err := mockService.GetAllRoles()
	var expectedDTOs []models.RoleDTO
	automapper.Map(mockRoles, &expectedDTOs)
	assert.NoError(t, err, "Error getting all roles")
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, expectedDTOs, roles)
}

// TestGetAllRoles_FindAllError tests the GetAllRoles function using mockRoleRepo
func TestGetAllRoles_FindAllError(t *testing.T) {
	// create a mock role repo
	// create a mock role service
	// call the get all roles function
	// assert that the roles returned is the same as the roles passed in
	// assert that the status is InternalServerError
	// assert that there is no error
	mockRepo := NewMockRoleErrorRepo()
	mockService := NewRoleService(mockRepo)
	roles, status, err := mockService.GetAllRoles()
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, []models.RoleDTO{}, roles)
}

// TestUpdateRole tests the UpdateRole function using mockRoleRepo
func TestUpdateRole(t *testing.T) {
	// create a mock role repo
	// create a mock role service
	// call the update role function
	// assert that the role returned is the same as the role passed in
	// assert that the status is OK
	// assert that there is no error
	mockRepo := NewMockRoleRepo()
	mockService := NewRoleService(mockRepo)
	testRole := models.Role{
		Name: "TestTest",
	}
	role, status, err := mockService.UpdateRole(1, testRole)
	testRole.ID = uint(1)
	var expectedDTO models.RoleDTO
	automapper.Map(testRole, &expectedDTO)
	assert.NoError(t, err, "Error updating role")
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, expectedDTO, role)
}

// TestUpdateRole_FindByIDError tests the UpdateRole function using mockRoleRepo
func TestUpdateRole_FindByIDError(t *testing.T) {
	// create a mock role repo
	// create a mock role service
	// call the update role function
	// assert that the role returned is the same as the role passed in
	// assert that the status is NotFound
	// assert that there is no error
	mockRepo := NewMockRoleErrorRepo()
	mockService := NewRoleService(mockRepo)
	testRole := models.Role{
		Name: "TestTest",
	}
	role, status, err := mockService.UpdateRole(1, testRole)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, status)
	assert.Equal(t, models.RoleDTO{}, role)
}

// TestUpdateRole_UpdateError tests the UpdateRole function using mockRoleRepo
func TestUpdateRole_UpdateError(t *testing.T) {
	// create a mock role repo
	// create a mock role service
	// call the update role function
	// assert that the role returned is the same as the role passed in
	// assert that the status is InternalServerError
	// assert that there is no error
	mockRepo := NewMockRoleSpecificErrorRepo()
	mockService := NewRoleService(mockRepo)
	testRole := models.Role{
		Name: "TestTest",
	}
	role, status, err := mockService.UpdateRole(1, testRole)
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, models.RoleDTO{}, role)
}

// TestDeleteRole tests the DeleteRole function using mockRoleRepo
func TestDeleteRole(t *testing.T) {
	// create a mock role repo
	// create a mock role service
	// call the delete role function
	// assert that the status is OK
	// assert that there is no error
	mockRepo := NewMockRoleRepo()
	mockService := NewRoleService(mockRepo)
	role, status, err := mockService.DeleteRole(1)
	var expectedDTO models.RoleDTO
	automapper.Map(mockRoles[0], &expectedDTO)
	assert.NoError(t, err, "Error deleting role")
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, expectedDTO, role)
}

// TestDeleteRole_FindByIDError tests the DeleteRole function using mockRoleRepo
func TestDeleteRole_FindByIDError(t *testing.T) {
	// create a mock role repo
	// create a mock role service
	// call the delete role function
	// assert that the status is NotFound
	// assert that there is no error
	mockRepo := NewMockRoleErrorRepo()
	mockService := NewRoleService(mockRepo)
	role, status, err := mockService.DeleteRole(1)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, status)
	assert.Equal(t, models.RoleDTO{}, role)
}

// TestDeleteRole_DeleteError tests the DeleteRole function using mockRoleRepo
func TestDeleteRole_DeleteError(t *testing.T) {
	// create a mock role repo
	// create a mock role service
	// call the delete role function
	// assert that the status is InternalServerError
	// assert that there is no error
	mockRepo := NewMockRoleSpecificErrorRepo()
	mockService := NewRoleService(mockRepo)
	role, status, err := mockService.DeleteRole(1)
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, models.RoleDTO{}, role)
}
