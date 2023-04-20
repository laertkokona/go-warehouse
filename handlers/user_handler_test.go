package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/laertkokona/crud-test/helpers"
	"github.com/laertkokona/crud-test/models"
	"github.com/laertkokona/crud-test/utils"
	"github.com/peteprogrammer/go-automapper"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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
		Password:  "Test1234!",
		RoleID:    1,
	},
	{
		Model:     mockModel[1],
		FirstName: "Admin",
		LastName:  "Test",
		Username:  "adminTest",
		Password:  "Test1234!",
		RoleID:    2,
	},
	{
		Model:     mockModel[2],
		FirstName: "SysAdmin",
		LastName:  "Test",
		Username:  "sysAdminTest",
		Password:  "Test1234!",
		RoleID:    3,
	},
}

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

// mockUserService is a mock implementation of the services.UserService interface
type mockUserService struct {
	createUser  func(user models.User) (models.UserDTO, int, error)
	getUser     func(id int) (models.UserDTO, int, error)
	getAllUsers func(pagination models.Pagination) ([]models.UserDTO, int, error)
	signInUser  func(loginUser models.Login) (string, int, error)
	signOutUser func() string
	updateUser  func(id int, user models.User) (models.UserDTO, int, error)
	deleteUser  func(id int) (models.UserDTO, int, error)
}

// mockRoleService is a mock implementation of the services.RoleService interface
type mockRoleService struct {
	createRole  func(role models.Role) (models.RoleDTO, int, error)
	getRole     func(id int) (models.RoleDTO, int, error)
	getAllRoles func() ([]models.RoleDTO, int, error)
	updateRole  func(id int, role models.Role) (models.RoleDTO, int, error)
	deleteRole  func(id int) (models.RoleDTO, int, error)
}

// CreateUser method that takes a models.User object and saves it to the database
func (m *mockUserService) CreateUser(user models.User) (models.UserDTO, int, error) {
	return m.createUser(user)
}

// GetUser method that takes a user id and returns the user object
func (m *mockUserService) GetUser(id int) (models.UserDTO, int, error) {
	return m.getUser(id)
}

// GetAllUsers method that takes a models.Pagination object and returns a slice of user objects
func (m *mockUserService) GetAllUsers(pagination models.Pagination) ([]models.UserDTO, int, error) {
	return m.getAllUsers(pagination)
}

// SignInUser method that takes a models.User object and returns a token
func (m *mockUserService) SignInUser(loginUser models.Login) (string, int, error) {
	return m.signInUser(loginUser)
}

// SignOutUser method that takes a token and returns a message
func (m *mockUserService) SignOutUser() string {
	return m.signOutUser()
}

// UpdateUser method that takes a user id and a user object and updates the user object in the database
func (m *mockUserService) UpdateUser(id int, user models.User) (models.UserDTO, int, error) {
	return m.updateUser(id, user)
}

// DeleteUser method that takes a user id and deletes the user object from the database
func (m *mockUserService) DeleteUser(id int) (models.UserDTO, int, error) {
	return m.deleteUser(id)
}

// CreateRole method that takes a models.Role object and saves it to the database
func (m *mockRoleService) CreateRole(role models.Role) (models.RoleDTO, int, error) {
	return m.createRole(role)
}

// GetRole method that takes a role id and returns the role object
func (m *mockRoleService) GetRole(id int) (models.RoleDTO, int, error) {
	return m.getRole(id)
}

// GetAllRoles method that takes a models.Pagination object and returns a slice of role objects
func (m *mockRoleService) GetAllRoles() ([]models.RoleDTO, int, error) {
	return m.getAllRoles()
}

// UpdateRole method that takes a role id and a role object and updates the role object in the database
func (m *mockRoleService) UpdateRole(id int, role models.Role) (models.RoleDTO, int, error) {
	return m.updateRole(id, role)
}

// DeleteRole method that takes a role id and deletes the role object from the database
func (m *mockRoleService) DeleteRole(id int) (models.RoleDTO, int, error) {
	return m.deleteRole(id)
}

// newMockUserService returns a new mockUserService
func newMockUserService() *mockUserService {
	return &mockUserService{
		createUser: func(user models.User) (models.UserDTO, int, error) {
			var userDTO models.UserDTO
			automapper.Map(user, &userDTO)
			return userDTO, http.StatusOK, nil
		},
		getUser: func(id int) (models.UserDTO, int, error) {
			var userDTO models.UserDTO
			automapper.Map(mockUsers[id-1], &userDTO)
			return userDTO, http.StatusOK, nil
		},
		getAllUsers: func(pagination models.Pagination) ([]models.UserDTO, int, error) {
			var mockUsersDTO []models.UserDTO
			automapper.Map(mockUsers, &mockUsersDTO)
			return mockUsersDTO, http.StatusOK, nil
		},
		signInUser: func(loginUser models.Login) (string, int, error) {
			var user models.User
			automapper.MapLoose(loginUser, &user)
			return utils.GenerateToken(user, utils.GetRoleName(utils.User)), http.StatusOK, nil
		},
		signOutUser: func() string {
			return utils.GenerateExpiredToken()
		},
		updateUser: func(id int, user models.User) (models.UserDTO, int, error) {
			mockUser := mockUsers[id-1]
			utils.CopyNonEmptyFields(&mockUser, &user)
			var userDTO models.UserDTO
			automapper.Map(mockUser, &userDTO)
			return userDTO, http.StatusOK, nil
		},
		deleteUser: func(id int) (models.UserDTO, int, error) {
			var userDTO models.UserDTO
			automapper.Map(mockUsers[id-1], &userDTO)
			return userDTO, http.StatusOK, nil
		},
	}
}

// newMockUserErrorService returns a new mockUserService with errors
func newMockUserErrorService() *mockUserService {
	return &mockUserService{
		createUser: func(user models.User) (models.UserDTO, int, error) {
			return models.UserDTO{}, http.StatusInternalServerError, errors.New("error creating user")
		},
		getUser: func(id int) (models.UserDTO, int, error) {
			return models.UserDTO{}, http.StatusInternalServerError, errors.New("error getting user")
		},
		getAllUsers: func(pagination models.Pagination) ([]models.UserDTO, int, error) {
			return []models.UserDTO{}, http.StatusInternalServerError, errors.New("error getting all users")
		},
		signInUser: func(loginUser models.Login) (string, int, error) {
			var user models.User
			automapper.MapLoose(loginUser, &user)
			return "", http.StatusInternalServerError, errors.New("error signing in user")
		},
		signOutUser: func() string {
			return utils.GenerateExpiredToken()
		},
		updateUser: func(id int, user models.User) (models.UserDTO, int, error) {
			return models.UserDTO{}, http.StatusInternalServerError, errors.New("error updating user")
		},
		deleteUser: func(id int) (models.UserDTO, int, error) {
			return models.UserDTO{}, http.StatusInternalServerError, errors.New("error deleting user")
		},
	}
}

// newMockRoleService returns a new mockRoleService
func newMockRoleService() *mockRoleService {
	return &mockRoleService{
		createRole: func(role models.Role) (models.RoleDTO, int, error) {
			var roleDTO models.RoleDTO
			automapper.Map(role, &roleDTO)
			return roleDTO, http.StatusOK, nil
		},
		getRole: func(id int) (models.RoleDTO, int, error) {
			var roleDTO models.RoleDTO
			automapper.Map(mockRoles[id-1], &roleDTO)
			return roleDTO, http.StatusOK, nil
		},
		getAllRoles: func() ([]models.RoleDTO, int, error) {
			var mockRolesDTO []models.RoleDTO
			automapper.Map(mockRoles, &mockRolesDTO)
			return mockRolesDTO, http.StatusOK, nil
		},
		updateRole: func(id int, role models.Role) (models.RoleDTO, int, error) {
			mockRole := mockRoles[id-1]
			utils.CopyNonEmptyFields(&mockRole, &role)
			var roleDTO models.RoleDTO
			automapper.Map(mockRole, &roleDTO)
			return roleDTO, http.StatusOK, nil
		},
		deleteRole: func(id int) (models.RoleDTO, int, error) {
			var roleDTO models.RoleDTO
			automapper.Map(mockRoles[id-1], &roleDTO)
			return roleDTO, http.StatusOK, nil
		},
	}
}

// newMockRoleErrorService returns a new mockRoleService with errors
func newMockRoleErrorService() *mockRoleService {
	return &mockRoleService{
		createRole: func(role models.Role) (models.RoleDTO, int, error) {
			return models.RoleDTO{}, http.StatusInternalServerError, errors.New("error creating role")
		},
		getRole: func(id int) (models.RoleDTO, int, error) {
			return models.RoleDTO{}, http.StatusInternalServerError, errors.New("error getting role")
		},
		getAllRoles: func() ([]models.RoleDTO, int, error) {
			return []models.RoleDTO{}, http.StatusInternalServerError, errors.New("error getting all roles")
		},
		updateRole: func(id int, role models.Role) (models.RoleDTO, int, error) {
			return models.RoleDTO{}, http.StatusInternalServerError, errors.New("error updating role")
		},
		deleteRole: func(id int) (models.RoleDTO, int, error) {
			return models.RoleDTO{}, http.StatusInternalServerError, errors.New("error deleting role")
		},
	}
}

// TestCreateUser tests the CreateUser method
func TestCreateUser(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	mockUserString, err := json.Marshal(mockUsers[0])
	assert.NoError(t, err, "Error marshalling mock user")

	r := gin.Default()
	r.POST("/users", userHandler.CreateUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(mockUserString))
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err = json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err, "Error unmarshalling result")
	userJSON, _ := json.Marshal(mockUsers[0])
	var user models.UserDTO
	err = json.Unmarshal(userJSON, &user)
	assert.NoError(t, err, "Error unmarshalling user")
	//automapper.MapLoose(result.Data, &user)

	var expectedDTO models.UserDTO
	automapper.Map(mockUsers[0], &expectedDTO)

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
	assert.NoError(t, err, "Error unmarshalling user")
	assert.Equal(t, expectedDTO, user, "User should be the same")
}

// TestCreateUser_InvalidJSONError tests the CreateUser method when the request body is not valid JSON
func TestCreateUser_InvalidJSONError(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.POST("/users", userHandler.CreateUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte("{{")))
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusBadRequest, w.Code, "Status code should be 400")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestCreateUser_ServiceError tests the CreateUser method when the service returns an error
func TestCreateUser_ServiceError(t *testing.T) {
	mockUserService := newMockUserErrorService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	mockUserString, err := json.Marshal(mockUsers[0])
	assert.NoError(t, err, "Error marshalling mock user")

	r := gin.Default()
	r.POST("/users", userHandler.CreateUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(mockUserString))
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err = json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Status code should be 400")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestGetUser tests the GetUser method
func TestGetUser(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.GET("/users/:id", userHandler.GetUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err, "Error unmarshalling user")
	userJSON, _ := json.Marshal(mockUsers[0])
	var user models.UserDTO
	err = json.Unmarshal(userJSON, &user)
	assert.NoError(t, err, "Error unmarshalling user")
	var expectedDTO models.UserDTO
	automapper.Map(mockUsers[0], &expectedDTO)
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
	assert.NoError(t, err, "Error unmarshalling user")
	assert.Equal(t, expectedDTO, user, "User should be the same")
}

// TestGetUser_InvalidIDError tests the GetUser method when the id is not an integer
func TestGetUser_InvalidIDError(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.GET("/users/:id", userHandler.GetUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/abc", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusBadRequest, w.Code, "Status code should be 400")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestGetUser_ServiceError tests the GetUser method when the service returns an error
func TestGetUser_ServiceError(t *testing.T) {
	mockUserService := newMockUserErrorService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.GET("/users/:id", userHandler.GetUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Status code should be 400")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestGetAllUsers tests the GetUsers method
func TestGetAllUsers(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.GET("/users", userHandler.GetAllUsers)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err, "Error unmarshalling user")
	usersJSON, _ := json.Marshal(result.Data)
	var users []models.UserDTO
	err = json.Unmarshal(usersJSON, &users)
	assert.NoError(t, err, "Error unmarshalling user")

	// Convert the mock users to DTOs
	var mockUsersDTO []models.UserDTO
	for _, user := range mockUsers {
		var userDTO models.UserDTO
		automapper.Map(user, &userDTO)
		mockUsersDTO = append(mockUsersDTO, userDTO)
	}

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
	assert.NoError(t, err, "Error unmarshalling users")
	assert.Equal(t, mockUsersDTO, users, "Users should be the same")
}

// TestGetAllUsers_ServiceError tests the GetUsers method when the service returns an error
func TestGetAllUsers_ServiceError(t *testing.T) {
	mockUserService := newMockUserErrorService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.GET("/users", userHandler.GetAllUsers)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Status code should be 400")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestSignInUser tests the SignInUser method
func TestSignInUser(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	mockLogInUser := models.Login{
		Username: "userTest",
		Password: "Test1234!",
	}

	mockUserString, _ := json.Marshal(mockLogInUser)

	r := gin.Default()
	r.POST("/users/signIn", userHandler.SignInUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/signIn", bytes.NewBuffer(mockUserString))
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NotEmpty(t, w.Header().Get("Authorization"), "Authorization header should not be empty")
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
	assert.NoError(t, err, "Error unmarshalling user")
	assert.NotEmpty(t, result.Message, "Success")
}

// TestSignInUser_InvalidUserError tests the SignInUser method when the user is invalid
func TestSignInUser_InvalidUserError(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	mockLogInUser := "//\\"

	r := gin.Default()
	r.POST("/users/signIn", userHandler.SignInUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/signIn", bytes.NewBuffer([]byte(mockLogInUser)))
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusBadRequest, w.Code, "Status code should be 400")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestSignInUser_ServiceError tests the SignInUser method when the service returns an error
func TestSignInUser_ServiceError(t *testing.T) {
	mockUserService := newMockUserErrorService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	mockLogInUser := models.Login{
		Username: "userTest",
		Password: "Test1234!",
	}

	mockUserString, _ := json.Marshal(mockLogInUser)
	r := gin.Default()
	r.POST("/users/signIn", userHandler.SignInUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/signIn", bytes.NewBuffer(mockUserString))
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Status code should be 400")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestSignOutUser tests the SignOutUser method
func TestSignOutUser(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.POST("/users/signOut", userHandler.SignOutUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/signOut", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
	assert.NoError(t, err, "Error unmarshalling user")
}

// TestUpdateUser tests the UpdateUser method
func TestUpdateUser(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	mockUser := models.User{
		FirstName: "UserTest",
		LastName:  "UserTest",
	}

	mockUserString, _ := json.Marshal(mockUser)

	r := gin.Default()
	r.PUT("/users/:id", userHandler.UpdateUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(mockUserString))
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err, "Error unmarshalling user")
	var user models.UserDTO
	userJSON, _ := json.Marshal(result.Data)
	err = json.Unmarshal(userJSON, &user)
	assert.NoError(t, err, "Error unmarshalling user")
	mockResult := mockUsers[0]
	utils.CopyNonEmptyFields(&mockResult, &mockUser)
	var expectedDTO models.UserDTO
	automapper.Map(mockResult, &expectedDTO)
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
	assert.NoError(t, err, "Error unmarshalling user")
	assert.NotEmpty(t, result.Message, "Success")
	assert.Equal(t, expectedDTO, user, "User should be the same")
}

// TestUpdateUser_InvalidIDError tests the UpdateUser method when the ID is invalid
func TestUpdateUser_InvalidIDError(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	mockUser := models.User{
		FirstName: "UserTest",
		LastName:  "UserTest",
	}

	mockUserString, _ := json.Marshal(mockUser)

	r := gin.Default()
	r.PUT("/users/:id", userHandler.UpdateUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/users/invalidID", bytes.NewBuffer(mockUserString))
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusBadRequest, w.Code, "Status code should be 400")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestUpdateUser_BindingError tests the UpdateUser method when the binding fails
func TestUpdateUser_BindingError(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.PUT("/users/:id", userHandler.UpdateUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer([]byte("")))
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusBadRequest, w.Code, "Status code should be 400")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestUpdateUser_ServiceError tests the UpdateUser method when the service fails
func TestUpdateUser_ServiceError(t *testing.T) {
	mockUserService := newMockUserErrorService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	mockUser := models.User{
		FirstName: "UserTest",
		LastName:  "UserTest",
	}

	mockUserString, _ := json.Marshal(mockUser)

	r := gin.Default()
	r.PUT("/users/:id", userHandler.UpdateUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/users/2", bytes.NewBuffer(mockUserString))
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Status code should be 500")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestDeleteUser tests the DeleteUser method
func TestDeleteUser(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.DELETE("/users/:id", userHandler.DeleteUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err, "Error unmarshalling response")
	var user models.UserDTO
	userJSON, _ := json.Marshal(result.Data)
	var expectedDTO models.UserDTO
	automapper.Map(mockUsers[0], &expectedDTO)
	err = json.Unmarshal(userJSON, &user)
	assert.NoError(t, err, "Error unmarshalling response")
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.Equal(t, result.Message, "Success")
	assert.Equal(t, expectedDTO, user, "Data should be nil")
}

// TestDeleteUser_InvalidIDError tests the DeleteUser method when the ID is invalid
func TestDeleteUser_InvalidIDError(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.DELETE("/users/:id", userHandler.DeleteUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/invalidID", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusBadRequest, w.Code, "Status code should be 400")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestDeleteUser_ServiceError tests the DeleteUser method when the service fails
func TestDeleteUser_ServiceError(t *testing.T) {
	mockUserService := newMockUserErrorService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.DELETE("/users/:id", userHandler.DeleteUser)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/2", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Status code should be 500")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestCreateRole tests the CreateRole method
func TestCreateRole(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	mockRole := models.Role{
		Name: "RoleTest",
	}

	mockRoleString, _ := json.Marshal(mockRole)

	r := gin.Default()
	r.POST("/roles", userHandler.CreateRole)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/roles", bytes.NewBuffer(mockRoleString))
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err, "Error unmarshalling response")
	var role models.Role
	roleJSON, _ := json.Marshal(result.Data)
	err = json.Unmarshal(roleJSON, &role)
	assert.NoError(t, err, "Error unmarshalling response")
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 201")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.Equal(t, result.Message, "Success")
	assert.Equal(t, mockRole, role, "Data should be nil")
}

// TestCreateRole_InvalidIDError tests the CreateRole method when the ID is invalid
func TestCreateRole_InvalidIDError(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	//mockRole := models.Role{
	//	Name: "RoleTest",
	//}

	//mockRoleString, _ := json.Marshal(mockRole)

	r := gin.Default()
	r.POST("/roles", userHandler.CreateRole)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/roles", bytes.NewBuffer([]byte("invalidRoleString")))
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusBadRequest, w.Code, "Status code should be 400")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestCreateRole_ServiceError tests the CreateRole method when the service fails
func TestCreateRole_ServiceError(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleErrorService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	mockRole := models.Role{
		Name: "RoleTest",
	}

	mockRoleString, _ := json.Marshal(mockRole)

	r := gin.Default()
	r.POST("/roles", userHandler.CreateRole)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/roles", bytes.NewBuffer(mockRoleString))
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Status code should be 500")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestGetRole tests the GetRole method
func TestGetRole(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.GET("/roles/:id", userHandler.GetRole)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/roles/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err, "Error unmarshalling response")
	var role models.Role
	roleJSON, _ := json.Marshal(result.Data)
	err = json.Unmarshal(roleJSON, &role)
	assert.NoError(t, err, "Error unmarshalling response")
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.Equal(t, result.Message, "Success")
	assert.Equal(t, mockRoles[0], role, "Data should be nil")
}

// TestGetRole_InvalidIDError tests the GetRole method when the ID is invalid
func TestGetRole_InvalidIDError(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.GET("/roles/:id", userHandler.GetRole)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/roles/invalidID", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusBadRequest, w.Code, "Status code should be 400")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestGetRole_ServiceError tests the GetRole method when the service fails
func TestGetRole_ServiceError(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleErrorService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.GET("/roles/:id", userHandler.GetRole)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/roles/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Status code should be 500")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestGetAllRoles tests the GetAllRoles method
func TestGetAllRoles(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.GET("/roles", userHandler.GetAllRoles)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/roles", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err, "Error unmarshalling response")
	var roles []models.Role
	rolesJSON, _ := json.Marshal(result.Data)
	err = json.Unmarshal(rolesJSON, &roles)
	assert.NoError(t, err, "Error unmarshalling response")
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.Equal(t, result.Message, "Success")
	assert.Equal(t, mockRoles, roles, "Data should be nil")
}

// TestGetAllRoles_ServiceError tests the GetAllRoles method when the service fails
func TestGetAllRoles_ServiceError(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleErrorService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.GET("/roles", userHandler.GetAllRoles)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/roles", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Status code should be 500")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestUpdateRole tests the UpdateRole method
func TestUpdateRole(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	mockRole := models.Role{
		Name: "RoleTest",
	}

	mockRoleString, _ := json.Marshal(mockRole)

	r := gin.Default()
	r.PUT("/roles/:id", userHandler.UpdateRole)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/roles/1", bytes.NewBuffer(mockRoleString))
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err, "Error unmarshalling role")
	var role models.RoleDTO
	roleJSON, _ := json.Marshal(result.Data)
	err = json.Unmarshal(roleJSON, &role)
	assert.NoError(t, err, "Error unmarshalling role")
	mockResult := mockRoles[0]
	utils.CopyNonEmptyFields(&mockResult, &mockRole)
	var expectedDTO models.RoleDTO
	automapper.Map(mockResult, &expectedDTO)
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
	assert.NoError(t, err, "Error unmarshalling role")
	assert.NotEmpty(t, result.Message, "Success")
	assert.Equal(t, expectedDTO, role, "Role should be the same")
}

// TestUpdateRole_InvalidIDError tests the UpdateRole method when the id is invalid
func TestUpdateRole_InvalidIDError(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.PUT("/roles/:id", userHandler.UpdateRole)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/roles/invalidID", strings.NewReader(`{"name": "newRole"}`))
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusBadRequest, w.Code, "Status code should be 400")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestUpdateRole_InvalidJSONError tests the UpdateRole method when the json is invalid
func TestUpdateRole_InvalidJSONError(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.PUT("/roles/:id", userHandler.UpdateRole)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/roles/1", strings.NewReader(`{"name": "newRole"`))
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusBadRequest, w.Code, "Status code should be 400")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestUpdateRole_ServiceError tests the UpdateRole method when the service fails
func TestUpdateRole_ServiceError(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleErrorService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.PUT("/roles/:id", userHandler.UpdateRole)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/roles/1", strings.NewReader(`{"name": "newRole"}`))
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Status code should be 500")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestDeleteRole tests the DeleteRole method
func TestDeleteRole(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.DELETE("/roles/:id", userHandler.DeleteRole)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/roles/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err, "Error unmarshalling response")
	var role models.Role
	roleJSON, _ := json.Marshal(result.Data)
	err = json.Unmarshal(roleJSON, &role)
	assert.NoError(t, err, "Error unmarshalling response")
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.Equal(t, result.Message, "Success")
	assert.Equal(t, mockRoles[0], role, "Data should be nil")
}

// TestDeleteRole_InvalidIDError tests the DeleteRole method when the id is invalid
func TestDeleteRole_InvalidIDError(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.DELETE("/roles/:id", userHandler.DeleteRole)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/roles/invalidID", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusBadRequest, w.Code, "Status code should be 400")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}

// TestDeleteRole_ServiceError tests the DeleteRole method when the service fails
func TestDeleteRole_ServiceError(t *testing.T) {
	mockUserService := newMockUserService()
	mockRoleService := newMockRoleErrorService()
	userHandler := NewUserHandler(mockUserService, mockRoleService)

	r := gin.Default()
	r.DELETE("/roles/:id", userHandler.DeleteRole)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/roles/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"), "Content-Type should be application/json")

	var result helpers.JSONResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Status code should be 500")
	assert.NoError(t, err, "Error unmarshalling response")
	assert.NotNil(t, result.Message, "Error should not be nil")
	assert.Nil(t, result.Data, "Data should be nil")
}
