package services

import (
	"errors"
	"github.com/laertkokona/crud-test/models"
	"github.com/laertkokona/crud-test/repositories"
	"github.com/laertkokona/crud-test/utils"
	"github.com/peteprogrammer/go-automapper"
	"net/http"
)

// UserService interface
type UserService interface {
	CreateUser(user models.User) (models.UserDTO, int, error)
	GetUser(id int) (models.UserDTO, int, error)
	GetAllUsers(pagination models.Pagination) ([]models.UserDTO, int, error)
	SignInUser(loginUser models.Login) (string, int, error)
	SignOutUser() string
	UpdateUser(id int, user models.User) (models.UserDTO, int, error)
	DeleteUser(id int) (models.UserDTO, int, error)
}

// userService struct
type userService struct {
	userRepo repositories.UserRepo
	roleRepo repositories.RoleRepo
}

// NewUserService returns a new instance of UserService
func NewUserService(uRepo repositories.UserRepo, rRepo repositories.RoleRepo) UserService {
	return userService{
		userRepo: uRepo,
		roleRepo: rRepo,
	}
}

// CreateUser method that takes a models.User object, hashes its password and saves it to the database
func (u userService) CreateUser(user models.User) (models.UserDTO, int, error) {
	// get the user's password
	// hash the user's password
	// save the user to the database
	// set the user's password to an empty string
	// return the user object
	password := utils.GetHashPassword(user.Password)
	user.Password = password
	user, err := u.userRepo.Save(user)
	if err != nil {
		return models.UserDTO{}, http.StatusInternalServerError, err
	}
	user.Password = ""
	var returnUser models.UserDTO
	automapper.Map(user, &returnUser)
	return returnUser, http.StatusOK, nil
}

// GetUser method that takes a user id and returns the user object
func (u userService) GetUser(id int) (models.UserDTO, int, error) {
	// get the user object from the database
	// set the user's password to an empty string
	// return the user object
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return models.UserDTO{}, http.StatusInternalServerError, err
	}
	user.Password = ""
	var returnUser models.UserDTO
	automapper.Map(user, &returnUser)
	return returnUser, http.StatusOK, nil
}

// GetAllUsers method that returns all the users in the database
func (u userService) GetAllUsers(pagination models.Pagination) ([]models.UserDTO, int, error) {
	// get all the users from the database
	// set the user's password to an empty string
	// return the user object
	users, err := u.userRepo.FindAll(pagination)
	if err != nil {
		return []models.UserDTO{}, http.StatusInternalServerError, err
	}
	for i := range users {
		users[i].Password = ""
	}
	var returnUsers []models.UserDTO
	automapper.Map(users, &returnUsers)
	//for _, user := range users {
	//	var returnUser models.UserDTO
	//	automapper.Map(user, &returnUser)
	//	returnUsers = append(returnUsers, returnUser)
	//}
	return returnUsers, http.StatusOK, nil
}

// UpdateUser method that takes a user id and updates the user object in the database
func (u userService) UpdateUser(id int, user models.User) (models.UserDTO, int, error) {
	// find the user object in the database
	// unmarshal the user object from the request body
	// set the user's id to the id of the user object in the database
	// set the user's password to the password of the user object in the database
	// update the user object in the database
	// set the user's password to an empty string
	// return the user object
	userDb, err := u.userRepo.FindByID(id)
	if err != nil {
		return models.UserDTO{}, http.StatusNotFound, err
	}
	utils.CopyNonEmptyFields(&userDb, &user)
	userDb, err = u.userRepo.Update(userDb)
	if err != nil {
		return models.UserDTO{}, http.StatusInternalServerError, err
	}
	var returnUser models.UserDTO
	automapper.Map(userDb, &returnUser)
	return returnUser, http.StatusOK, nil
}

// DeleteUser method that takes a user id and deletes the user object from the database
func (u userService) DeleteUser(id int) (models.UserDTO, int, error) {
	// find the user object in the database
	// delete the user object from the database
	// set the user's password to an empty string
	// return the user object
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return models.UserDTO{}, http.StatusNotFound, err
	}
	err = u.userRepo.Delete(user)
	if err != nil {
		return models.UserDTO{}, http.StatusInternalServerError, err
	}
	var returnUser models.UserDTO
	automapper.Map(user, &returnUser)
	return returnUser, http.StatusOK, nil
}

// SignInUser method that takes a models.User object and returns a token
func (u userService) SignInUser(loginUser models.Login) (string, int, error) {
	// find the user object in the database
	// compare the user's password with the password in the database
	// if the passwords don't match, return an error
	// find the user's role in the database
	// generate a token
	// return the token
	userDb, err := u.userRepo.FindByUsername(loginUser.Username)
	if err != nil {
		return "", http.StatusNotFound, err
	}
	if !utils.ComparePassword([]byte(userDb.Password), []byte(loginUser.Password)) {
		return "", http.StatusUnauthorized, errors.New("invalid password")
	}
	role, err := u.roleRepo.FindByID(userDb.RoleID)
	if err != nil {
		return "", http.StatusNotFound, err
	}
	token := utils.GenerateToken(userDb, role.Name)
	return token, http.StatusOK, nil
}

// SignOutUser method that returns an expired token
func (u userService) SignOutUser() string {
	return utils.GenerateExpiredToken()
}
