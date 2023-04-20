package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/laertkokona/crud-test/helpers"
	"github.com/laertkokona/crud-test/models"
	"github.com/laertkokona/crud-test/services"
	"github.com/peteprogrammer/go-automapper"
	"net/http"
	"strconv"
)

// UserHandler interface
type UserHandler interface {
	CreateUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	SignInUser(ctx *gin.Context)
	SignOutUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	CreateRole(ctx *gin.Context)
	GetRole(ctx *gin.Context)
	GetAllRoles(ctx *gin.Context)
	UpdateRole(ctx *gin.Context)
	DeleteRole(ctx *gin.Context)
}

// userHandler struct that implements UserHandler interface and has a userService field and a roleService field
type userHandler struct {
	userService services.UserService
	roleService services.RoleService
}

// CreateUser gets a user object from the request body and saves it to the database
//
// CreateUser godoc
// @Summary Create a new user
// @Description Create new user
// @Security 	 ApiKeyAuth
// @Tags User
// @ID create-user
// @Accept  json
// @Produce  json
// @Param user body models.User true "User object"
// @Success 200 {object} helpers.JSONSuccessResult{data=models.UserDTO}
// @Failure 400 {object} helpers.JSONBadRequestResult
// @Failure 401 {object} helpers.JSONUnauthorizedResult
// @Failure 500 {object} helpers.JSONInternalServerErrorResult
// @Router /users [post]
func (u userHandler) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		helpers.FailedResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	userDTO, status, err := u.userService.CreateUser(user)
	if err != nil {
		//ctx.JSON(status, gin.H{"error": err.Error()})
		helpers.FailedResponse(ctx, status, err.Error(), nil)
		return
	}
	//ctx.JSON(status, user)
	helpers.SuccessResponse(ctx, userDTO)
}

// GetUser gets a user id from the request params and returns the user object
//
// GetUser godoc
// @Summary      Get a user by id
// @Description  Get a user by id
// @Security 	 ApiKeyAuth
// @Tags User
// @ID           get-user
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Success      200 {object} helpers.JSONSuccessResult{data=models.UserDTO}
// @Failure      400 {object} helpers.JSONBadRequestResult
// @Failure      401 {object} helpers.JSONUnauthorizedResult
// @Failure      404 {object} helpers.JSONNotFoundResult
// @Failure      500 {object} helpers.JSONInternalServerErrorResult
// @Router       /users/{id} [get]
func (u userHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		helpers.FailedResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, status, err := u.userService.GetUser(intId)
	if err != nil {
		helpers.FailedResponse(ctx, status, err.Error(), nil)
		//ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	helpers.SuccessResponse(ctx, user)
	//ctx.JSON(status, user)
}

// GetAllUsers gets a page and limit from the request query and returns a list of users
//
// GetAllUsers godoc
// @Summary      List all users
// @Description  get users
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Tags User
// @Param 	  	 page query string false "Page number"
// @Param 	  	 limit query string false "Limit number"
// Param		 Bearer header string true "Bearer token"
// @Success      200 {array} helpers.JSONSuccessResult{data=[]models.UserDTO}
// @Failure      400 {object} helpers.JSONBadRequestResult
// @Failure      401 {object} helpers.JSONUnauthorizedResult
// @Failure      500 {object} helpers.JSONInternalServerErrorResult
// @Router       /users [get]
func (u userHandler) GetAllUsers(ctx *gin.Context) {
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		intPage = 0
		//ctx.JSON(400, gin.H{"error": err.Error()})
		//return
	}
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		intLimit = 0
		//ctx.JSON(400, gin.H{"error": err.Error()})
		//return
	}
	pagination := models.Pagination{
		Page:  intPage,
		Limit: intLimit,
	}
	users, status, err := u.userService.GetAllUsers(pagination)
	var usersDTO []models.UserDTO
	if err != nil {
		helpers.FailedResponse(ctx, status, err.Error(), nil)
		//ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	automapper.Map(users, &usersDTO)
	//ctx.JSON(status, usersDTO)
	helpers.SuccessResponse(ctx, usersDTO)
}

// SignInUser gets a user object from the request body and returns a token
//
// SignInUser godoc
// @Summary      Sign in user
// @Description  Sign in user
// @Accept       json
// @Produce      json
// @Tags Auth
// @Param        user body models.Login true "User object"
// @Success      200 {object} helpers.JSONSuccessResultNoData
// @Failure      400 {object} helpers.JSONBadRequestResult
// @Failure      401 {object} helpers.JSONUnauthorizedResult
// @Failure      500 {object} helpers.JSONInternalServerErrorResult
// @Router       /signIn [post]
func (u userHandler) SignInUser(ctx *gin.Context) {
	var loginUser models.Login
	if err := ctx.ShouldBindJSON(&loginUser); err != nil {
		helpers.FailedResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, status, err := u.userService.SignInUser(loginUser)
	if err != nil {
		helpers.FailedResponse(ctx, status, err.Error(), nil)
		return
	}
	ctx.Header("Authorization", token)
	helpers.SuccessResponseNoData(ctx)
}

// SignOutUser gets a user object from the request body and returns a token
//
// SignOutUser godoc
// @Summary      Sign out user
// @Description  Sign out user
// @Accept       json
// @Produce      json
// @Tags Auth
// @Success      200 {object} helpers.JSONSuccessResultNoData
// @Router       /signOut [post]
func (u userHandler) SignOutUser(ctx *gin.Context) {
	expToken := u.userService.SignOutUser()
	ctx.Header("Authorization", expToken)
	helpers.SuccessResponse(ctx, gin.H{"message": "User logged out successfully"})
	//ctx.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}

// UpdateUser gets a user object from the request body and returns the updated user object
//
// UpdateUser godoc
// @Summary      Update user
// @Description  Update user
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Tags User
// @Param        id path string true "User ID"
// @Param        user body models.User true "User object"
// @Success      200 {object} helpers.JSONSuccessResult{data=models.UserDTO}
// @Failure      400 {object} helpers.JSONBadRequestResult
// @Failure      401 {object} helpers.JSONUnauthorizedResult
// @Failure      500 {object} helpers.JSONInternalServerErrorResult
// @Router       /users/{id} [put]
func (u userHandler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		helpers.FailedResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		helpers.FailedResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userDTO, status, err := u.userService.UpdateUser(intId, user)
	if err != nil {
		helpers.FailedResponse(ctx, status, err.Error(), nil)
		//ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	helpers.SuccessResponse(ctx, userDTO)
	//ctx.JSON(status, user)
}

// DeleteUser gets a user object from the request body and returns the updated user object
//
// DeleteUser godoc
// @Summary      Delete user
// @Description  Delete user
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Tags User
// @Param        id path string true "User ID"
// @Success      200 {object} helpers.JSONSuccessResult{data=models.UserDTO}
// @Failure      400 {object} helpers.JSONBadRequestResult
// @Failure      401 {object} helpers.JSONUnauthorizedResult
// @Failure      500 {object} helpers.JSONInternalServerErrorResult
// @Router       /users/{id} [delete]
func (u userHandler) DeleteUser(ctx *gin.Context) {
	// get id from url
	// convert id to int
	// call service
	// return response
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		helpers.FailedResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, status, err := u.userService.DeleteUser(intId)
	if err != nil {
		helpers.FailedResponse(ctx, status, err.Error(), nil)
		//ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	helpers.SuccessResponse(ctx, user)
	//ctx.JSON(status, user)
}

// CreateRole gets a role object from the request body and returns the created role object
//
// CreateRole godoc
// @Summary      Create role
// @Description  Create role
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Tags Role
// @Param        role body models.Role true "Role object"
// @Success      200 {object} helpers.JSONSuccessResult{data=models.RoleDTO}
// @Failure      400 {object} helpers.JSONBadRequestResult
// @Failure      401 {object} helpers.JSONUnauthorizedResult
// @Failure      500 {object} helpers.JSONInternalServerErrorResult
// @Router       /roles [post]
func (u userHandler) CreateRole(ctx *gin.Context) {
	var role models.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		helpers.FailedResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userDTO, status, err := u.roleService.CreateRole(role)
	if err != nil {
		helpers.FailedResponse(ctx, status, err.Error(), nil)
		//ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	helpers.SuccessResponse(ctx, userDTO)
	//ctx.JSON(status, role)
}

// GetRole gets a role object from the request body and returns the created role object
//
// GetRole godoc
// @Summary      Get role
// @Description  Get role
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Tags Role
// @Param        id path string true "Role ID"
// @Success      200 {object} helpers.JSONSuccessResult{data=models.RoleDTO}
// @Failure      400 {object} helpers.JSONBadRequestResult
// @Failure      401 {object} helpers.JSONUnauthorizedResult
// @Failure      500 {object} helpers.JSONInternalServerErrorResult
// @Router       /roles/{id} [get]
func (u userHandler) GetRole(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		helpers.FailedResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	role, status, err := u.roleService.GetRole(intId)
	if err != nil {
		helpers.FailedResponse(ctx, status, err.Error(), nil)
		//ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	helpers.SuccessResponse(ctx, role)
	//ctx.JSON(status, role)
}

// GetAllRoles gets a role object from the request body and returns the created role object
//
// GetAllRoles godoc
// @Summary      Get all roles
// @Description  Get all roles
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Tags Role
// @Success      200 {object} helpers.JSONSuccessResult{data=[]models.RoleDTO}
// @Failure      400 {object} helpers.JSONBadRequestResult
// @Failure      401 {object} helpers.JSONUnauthorizedResult
// @Failure      500 {object} helpers.JSONInternalServerErrorResult
// @Router       /roles [get]
func (u userHandler) GetAllRoles(ctx *gin.Context) {
	roles, status, err := u.roleService.GetAllRoles()
	if err != nil {
		helpers.FailedResponse(ctx, status, err.Error(), nil)
		//ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	helpers.SuccessResponse(ctx, roles)
	//ctx.JSON(status, roles)
}

// UpdateRole gets a role object from the request body and returns the created role object
//
// UpdateRole godoc
// @Summary      Update role
// @Description  Update role
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Tags Role
// @Param        id path string true "Role ID"
// @Param        role body models.Role true "Role object"
// @Success      200 {object} helpers.JSONSuccessResult{data=models.RoleDTO}
// @Failure      400 {object} helpers.JSONBadRequestResult
// @Failure      401 {object} helpers.JSONUnauthorizedResult
// @Failure      500 {object} helpers.JSONInternalServerErrorResult
// @Router       /roles/{id} [put]
func (u userHandler) UpdateRole(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		helpers.FailedResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var role models.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		helpers.FailedResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roleDTO, status, err := u.roleService.UpdateRole(intId, role)
	if err != nil {
		helpers.FailedResponse(ctx, status, err.Error(), nil)
		//ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	helpers.SuccessResponse(ctx, roleDTO)
	//ctx.JSON(status, role)
}

// DeleteRole gets a role object from the request body and returns the created role object
//
// DeleteRole godoc
// @Summary      Delete role
// @Description  Delete role
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Tags Role
// @Param        id path string true "Role ID"
// @Success      200 {object} helpers.JSONSuccessResult{data=models.RoleDTO}
// @Failure      400 {object} helpers.JSONBadRequestResult
// @Failure      401 {object} helpers.JSONUnauthorizedResult
// @Failure      500 {object} helpers.JSONInternalServerErrorResult
// @Router       /roles/{id} [delete]
func (u userHandler) DeleteRole(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		helpers.FailedResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	role, status, err := u.roleService.DeleteRole(intId)
	if err != nil {
		helpers.FailedResponse(ctx, status, err.Error(), nil)
		//ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	helpers.SuccessResponse(ctx, role)
	//ctx.JSON(status, role)
}

func NewUserHandler(userService services.UserService, roleService services.RoleService) UserHandler {
	return userHandler{
		userService: userService,
		roleService: roleService,
	}
}
