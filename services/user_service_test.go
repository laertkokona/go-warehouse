package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/laertkokona/crud-test/models"
	"github.com/laertkokona/crud-test/utils"
	"github.com/peteprogrammer/go-automapper"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math"
	"net/http"
	"testing"
	"time"
)

var mockModel = []gorm.Model{
	{ID: 1},
	{ID: 2},
	{ID: 3},
}
var mockUsers = []models.User{
	{
		Model:     mockModel[0],
		FirstName: "User",
		LastName:  "Test",
		Username:  "userTest",
		RoleID:    1,
	},
	{
		Model:     mockModel[1],
		FirstName: "Admin",
		LastName:  "Test",
		Username:  "adminTest",
		RoleID:    2,
	},
	{
		Model:     mockModel[2],
		FirstName: "SysAdmin",
		LastName:  "Test",
		Username:  "sysAdminTest",
		RoleID:    3,
	},
}

// mockUserRepo is a mock implementation of the repositories.UserRepo interface
type mockUserRepo struct {
	// findAll is a mock function with given fields: pagination
	findAll func(pagination models.Pagination) ([]models.User, error)
	// findByID is a mock function with given fields: id
	findByID func(id int) (models.User, error)
	// findByUsername is a mock function with given fields: username
	findByUsername func(username string) (models.User, error)
	// save is a mock function with given fields: user
	save func(user models.User) (models.User, error)
	// update is a mock function with given fields: user
	update func(user models.User) (models.User, error)
	// delete is a mock function with given fields: user
	delete func(user models.User) error
	// deleteById is a mock function with given fields: id
	deleteById func(id int) (models.User, error)
}

// FindAll is a mock function with given fields: pagination
func (_m *mockUserRepo) FindAll(pagination models.Pagination) ([]models.User, error) {
	return _m.findAll(pagination)
}

// FindByID is a mock function with given fields: id
func (_m *mockUserRepo) FindByID(id int) (models.User, error) {
	return _m.findByID(id)
}

// FindByUsername is a mock function with given fields: username
func (_m *mockUserRepo) FindByUsername(username string) (models.User, error) {
	return _m.findByUsername(username)
}

// Save is a mock function with given fields: user
func (_m *mockUserRepo) Save(user models.User) (models.User, error) {
	return _m.save(user)
}

// Update is a mock function with given fields: user
func (_m *mockUserRepo) Update(user models.User) (models.User, error) {
	return _m.update(user)
}

// Delete is a mock function with given fields: user
func (_m *mockUserRepo) Delete(user models.User) error {
	return _m.delete(user)
}

// DeleteById is a mock function with given fields: id
func (_m *mockUserRepo) DeleteById(id int) (models.User, error) {
	return _m.deleteById(id)
}

// newMockUserRepo returns a new instance of mockUserRepo
func newMockUserRepo() *mockUserRepo {
	return &mockUserRepo{
		findAll: func(pagination models.Pagination) ([]models.User, error) {
			return mockUsers, nil
		},
		findByID: func(id int) (models.User, error) {
			return mockUsers[id-1], nil
		},
		findByUsername: func(username string) (models.User, error) {
			var user models.User
			for _, u := range mockUsers {
				if u.Username == username {
					user = u
				}
			}
			user.Password = utils.GetHashPassword("Test1234!")
			return user, nil
		},
		save: func(user models.User) (models.User, error) {
			user.Password = ""
			return user, nil
		},
		update: func(user models.User) (models.User, error) {
			user.ID = 1
			user.Password = ""
			return user, nil
		},
		delete: func(user models.User) error {
			return nil
		},
		deleteById: func(id int) (models.User, error) {
			return mockUsers[id-1], nil
		},
	}
}

// newMockUserErrorRepo returns a new instance of mockUserRepo with error
func newMockUserErrorRepo() *mockUserRepo {
	return &mockUserRepo{
		findAll: func(pagination models.Pagination) ([]models.User, error) {
			return []models.User{}, errors.New("error")
		},
		findByID: func(id int) (models.User, error) {
			return models.User{}, errors.New("error")
		},
		findByUsername: func(username string) (models.User, error) {
			return models.User{}, errors.New("error")
		},
		save: func(user models.User) (models.User, error) {
			return models.User{}, errors.New("error")
		},
		update: func(user models.User) (models.User, error) {
			return models.User{}, errors.New("error")
		},
		delete: func(user models.User) error {
			return errors.New("error")
		},
		deleteById: func(id int) (models.User, error) {
			return models.User{}, errors.New("error")
		},
	}
}

// newMockUserSpecificErrorRepo returns a new instance of mockUserRepo with error
func newMockUserSpecificErrorRepo() *mockUserRepo {
	return &mockUserRepo{
		findAll: func(pagination models.Pagination) ([]models.User, error) {
			for _, u := range mockUsers {
				u.Password = ""
			}
			return mockUsers, nil
		},
		findByID: func(id int) (models.User, error) {
			return mockUsers[id-1], nil
		},
		findByUsername: func(username string) (models.User, error) {
			var user models.User
			for _, u := range mockUsers {
				if u.Username == username {
					user = u
				}
			}
			return user, nil
		},
		save: func(user models.User) (models.User, error) {
			return models.User{}, errors.New("error")
		},
		update: func(user models.User) (models.User, error) {
			return models.User{}, errors.New("error")
		},
		delete: func(user models.User) error {
			return errors.New("error")
		},
		deleteById: func(id int) (models.User, error) {
			return models.User{}, errors.New("error")
		},
	}
}

// TestNewUserService is a test function for NewUserService
func TestNewUserService(t *testing.T) {
	mockUserRepo := newMockUserRepo()
	mockRoleRepo := NewMockRoleRepo()
	mockService := NewUserService(mockUserRepo, mockRoleRepo)
	assert.NotNil(t, mockService)
	assert.IsType(t, userService{}, mockService)
}

// TestCreateUser is a test function for CreateUser
func TestCreateUser(t *testing.T) {
	mockUserRepo := newMockUserRepo()
	mockRoleRepo := NewMockRoleRepo()
	mockService := NewUserService(mockUserRepo, mockRoleRepo)
	mockUser := models.User{
		FirstName: "User",
		LastName:  "Test",
		Username:  "testUser",
		Password:  utils.GetHashPassword("Test1234!"),
		RoleID:    1,
	}
	user, status, err := mockService.CreateUser(mockUser)
	mockUser.Password = ""
	var expectedDTO models.UserDTO
	automapper.Map(mockUser, &expectedDTO)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, expectedDTO, user)
}

// TestCreateUser_SaveError is a test function for CreateUser with error
func TestCreateUser_SaveError(t *testing.T) {
	mockUserRepo := newMockUserErrorRepo()
	mockRoleRepo := NewMockRoleRepo()
	mockService := NewUserService(mockUserRepo, mockRoleRepo)
	mockUser := models.User{
		FirstName: "User",
		LastName:  "Test",
		Username:  "testUser",
		Password:  utils.GetHashPassword("Test1234!"),
		RoleID:    1,
	}
	user, status, err := mockService.CreateUser(mockUser)
	mockUser.Password = ""
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, models.UserDTO{}, user)
}

// TestGetUser is a test function for GetUsers
func TestGetUser(t *testing.T) {
	mockUserRepo := newMockUserRepo()
	mockRoleRepo := NewMockRoleRepo()
	mockService := NewUserService(mockUserRepo, mockRoleRepo)
	user, status, err := mockService.GetUser(1)
	var expectedDTO models.UserDTO
	automapper.Map(mockUsers[0], &expectedDTO)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, expectedDTO, user)
}

// TestGetUser_FindByIDError is a test function for GetUsers with error
func TestGetUser_FindByIDError(t *testing.T) {
	mockUserRepo := newMockUserErrorRepo()
	mockRoleRepo := NewMockRoleRepo()
	mockService := NewUserService(mockUserRepo, mockRoleRepo)
	user, status, err := mockService.GetUser(1)
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, models.UserDTO{}, user)
}

// TestGetAllUsers is a test function for GetAllUsers
func TestGetAllUsers(t *testing.T) {
	mockUserRepo := newMockUserRepo()
	mockRoleRepo := NewMockRoleRepo()
	mockService := NewUserService(mockUserRepo, mockRoleRepo)
	users, status, err := mockService.GetAllUsers(models.Pagination{Page: 1, Limit: 10})
	var expectedDTOs []models.UserDTO
	automapper.Map(mockUsers, &expectedDTOs)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, expectedDTOs, users)
}

// TestGetAllUsers_FindAllError is a test function for GetAllUsers with error
func TestGetAllUsers_FindAllError(t *testing.T) {
	mockUserRepo := newMockUserErrorRepo()
	mockRoleRepo := NewMockRoleRepo()
	mockService := NewUserService(mockUserRepo, mockRoleRepo)
	users, status, err := mockService.GetAllUsers(models.Pagination{})
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, []models.UserDTO{}, users)
}

// TestUpdateUser is a test function for UpdateUser
func TestUpdateUser(t *testing.T) {
	mockUserRepo := newMockUserRepo()
	mockRoleRepo := NewMockRoleRepo()
	mockService := NewUserService(mockUserRepo, mockRoleRepo)
	mockUser := models.User{
		FirstName: "Test",
		LastName:  "Test",
		Username:  "testTest",
		RoleID:    1,
	}
	resUser, status, err := mockService.UpdateUser(1, mockUser)
	mockUser.ID = 1
	var expectedDTO models.UserDTO
	automapper.Map(mockUser, &expectedDTO)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, expectedDTO, resUser)
}

// TestUpdateUser_FindByIDError is a test function for UpdateUser with error
func TestUpdateUser_FindByIdError(t *testing.T) {
	mockUserRepo := newMockUserErrorRepo()
	mockRoleRepo := NewMockRoleRepo()
	mockService := NewUserService(mockUserRepo, mockRoleRepo)
	mockUser := models.User{
		FirstName: "Test",
		LastName:  "Test",
		Username:  "testTest",
		RoleID:    1,
	}
	resUser, status, err := mockService.UpdateUser(1, mockUser)
	mockUser.ID = 1
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, status)
	assert.Equal(t, models.UserDTO{}, resUser)
}

// TestUpdateUser_UpdateError is a test function for UpdateUser with error
func TestUpdateUser_UpdateError(t *testing.T) {
	mockUserRepo := newMockUserSpecificErrorRepo()
	mockRoleRepo := NewMockRoleRepo()
	mockService := NewUserService(mockUserRepo, mockRoleRepo)
	mockUser := models.User{
		FirstName: "Test",
		LastName:  "Test",
		Username:  "testTest",
	}
	resUser, status, err := mockService.UpdateUser(1, mockUser)
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, models.UserDTO{}, resUser)
}

// TestDeleteUser is a test function for DeleteUser
func TestDeleteUser(t *testing.T) {
	mockUserRepo := newMockUserRepo()
	mockRoleRepo := NewMockRoleRepo()
	mockService := NewUserService(mockUserRepo, mockRoleRepo)
	user, status, err := mockService.DeleteUser(1)
	var expectedDTO models.UserDTO
	automapper.Map(mockUsers[0], &expectedDTO)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, expectedDTO, user)
}

// TestDeleteUser_FindByIDError is a test function for DeleteUser with error
func TestDeleteUser_FindByIdError(t *testing.T) {
	mockUserRepo := newMockUserErrorRepo()
	mockRoleRepo := NewMockRoleRepo()
	mockService := NewUserService(mockUserRepo, mockRoleRepo)
	user, status, err := mockService.DeleteUser(1)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, status)
	assert.Equal(t, models.UserDTO{}, user)
}

// TestDeleteUser_DeleteError is a test function for DeleteUser with error
func TestDeleteUser_DeleteError(t *testing.T) {
	mockUserRepo := newMockUserSpecificErrorRepo()
	mockRoleRepo := NewMockRoleRepo()
	mockService := NewUserService(mockUserRepo, mockRoleRepo)
	user, status, err := mockService.DeleteUser(1)
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, models.UserDTO{}, user)
}

// TestSignInUser is a test function for SignInUser
func TestSignInUser(t *testing.T) {
	mockUserRepo := newMockUserRepo()
	mockRoleRepo := NewMockRoleRepo()
	mockService := NewUserService(mockUserRepo, mockRoleRepo)
	user := models.Login{
		Username: "sysAdminTest",
		Password: "Test1234!",
	}
	token, status, err := mockService.SignInUser(user)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotEmpty(t, token)
	jwtToken, err := utils.ValidateToken(token)
	assert.NoError(t, err)
	assert.True(t, jwtToken.Valid)
}

// TestSignInUser_FindByUsernameError is a test function for SignInUser with error
func TestSignInUser_FindByUsernameError(t *testing.T) {
	mockUserRepo := newMockUserErrorRepo()
	mockRoleRepo := NewMockRoleRepo()
	mockService := NewUserService(mockUserRepo, mockRoleRepo)
	user := models.Login{
		Username: "sysAdminTest",
		Password: "Test1234!",
	}
	token, status, err := mockService.SignInUser(user)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, status)
	assert.Empty(t, token)
}

// TestSignInUser_ValidatePasswordError is a test function for SignInUser with error
func TestSignInUser_ValidatePasswordError(t *testing.T) {
	mockUserRepo := newMockUserRepo()
	mockRoleRepo := NewMockRoleRepo()
	mockService := NewUserService(mockUserRepo, mockRoleRepo)
	user := models.Login{
		Username: "sysAdminTest",
		Password: "11!",
	}
	token, status, err := mockService.SignInUser(user)
	assert.Error(t, err)
	assert.Equal(t, http.StatusUnauthorized, status)
	assert.Empty(t, token)
}

// TestSignInUser_FindByUsernameError is a test function for SignInUser with error
func TestSignInUser_FindByIDError(t *testing.T) {
	mockUserRepo := newMockUserRepo()
	mockRoleRepo := NewMockRoleErrorRepo()
	mockService := NewUserService(mockUserRepo, mockRoleRepo)
	user := models.Login{
		Username: "sysAdminTest",
		Password: "Test1234!",
	}
	token, status, err := mockService.SignInUser(user)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, status)
	assert.Empty(t, token)
}

// TestSignOutUser is a test function for SignOutUser
func TestSignOutUser(t *testing.T) {
	mockUserRepo := newMockUserRepo()
	mockRoleRepo := NewMockRoleRepo()
	mockService := NewUserService(mockUserRepo, mockRoleRepo)
	token := mockService.SignOutUser()
	jwtToken, err := utils.ValidateToken(token)
	assert.NoError(t, err)
	assert.True(t, jwtToken.Valid)
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, claims["exp"], claims["iat"])
	sec, dec := math.Modf(claims["exp"].(float64))
	assert.True(t, time.Now().After(time.Unix(int64(sec), int64(dec*float64(time.Second)))))
}
