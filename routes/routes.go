package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/laertkokona/crud-test/handlers"
	"github.com/laertkokona/crud-test/middleware"
	"github.com/laertkokona/crud-test/repositories"
	"github.com/laertkokona/crud-test/services"
	"github.com/laertkokona/crud-test/utils"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/laertkokona/crud-test/docs"
)

// SetupRoutes sets up the routes
func SetupRoutes(DB *gorm.DB) {
	// new gin engine
	router := gin.New()

	// new user repository
	userRepo := repositories.NewUserRepo(DB)
	// new role repository
	roleRepo := repositories.NewRoleRepo(DB)
	// new item repository
	itemRepo := repositories.NewItemRepo(DB)
	// new truck repository
	truckRepo := repositories.NewTruckRepo(DB)
	// new order repository
	orderRepo := repositories.NewOrderRepo(DB)

	// new service for the user repository
	userService := services.NewUserService(userRepo, roleRepo)
	// new service for the role repository
	roleService := services.NewRoleService(roleRepo)
	// new service for the item repository
	itemService := services.NewItemService(itemRepo)
	// new service for the truck repository
	truckService := services.NewTruckService(truckRepo)
	// new service for the order repository
	orderService := services.NewOrderService(orderRepo)

	// new handler for the user service
	userHandler := handlers.NewUserHandler(userService, roleService)
	// new handler for the item service
	itemHandler := handlers.NewItemHandler(itemService)
	// new handler for the order service
	orderHandler := handlers.NewOrderHandler(orderService)
	// new handler for the truck service
	truckHandler := handlers.NewTruckHandler(truckService)

	// adding the recovery and logger middleware to the router
	router.Use(gin.Recovery(), gin.Logger())

	// the user routes
	userRoutes := router.Group("/users")
	userRoutes.Use(middleware.AuthMiddleware(utils.GetRoleName(utils.SysAdmin)))
	{
		userRoutes.GET("/", userHandler.GetAllUsers)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	// the sign in and out routes
	router.POST("/signIn", userHandler.SignInUser)
	router.POST("/signOut", userHandler.SignOutUser)

	// the role routes
	roleRoutes := router.Group("/roles")
	// the auth middleware to protect the routes from unauthorized access
	roleRoutes.Use(middleware.AuthMiddleware(utils.GetRoleName(utils.SysAdmin)))
	{
		roleRoutes.GET("/", userHandler.GetAllRoles)
		roleRoutes.GET("/:id", userHandler.GetRole)
		roleRoutes.POST("/", userHandler.CreateRole)
		roleRoutes.PUT("/:id", userHandler.UpdateRole)
		roleRoutes.DELETE("/:id", userHandler.DeleteRole)
	}

	// the item routes
	itemRoutes := router.Group("/items")
	// the auth middleware to protect the routes from unauthorized access
	itemRoutes.Use(middleware.AuthMiddleware())
	{
		itemRoutes.GET("/", itemHandler.GetAllItems)
		itemRoutes.GET("/:id", itemHandler.GetItem)
		itemRoutes.POST("/", itemHandler.CreateItem)
		itemRoutes.PUT("/:id", itemHandler.UpdateItem)
		itemRoutes.DELETE("/:id", itemHandler.DeleteItem)
	}

	// the truck routes
	truckRoutes := router.Group("/trucks")
	// the auth middleware to protect the routes from unauthorized access
	truckRoutes.Use(middleware.AuthMiddleware(utils.GetRoleName(utils.Admin)))
	{
		truckRoutes.GET("/", truckHandler.GetAllTrucks)
		truckRoutes.GET("/:id", truckHandler.GetTruck)
		truckRoutes.POST("/", truckHandler.CreateTruck)
		truckRoutes.PUT("/:id", truckHandler.UpdateTruck)
		truckRoutes.DELETE("/:id", truckHandler.DeleteTruck)
	}

	// the order routes
	orderRoutes := router.Group("/orders")
	// the auth middleware to protect the routes from unauthorized access
	orderRoutes.Use(middleware.AuthMiddleware(utils.GetRoleName(utils.User)))
	{
		orderRoutes.GET("/", orderHandler.GetAllOrders)
		orderRoutes.GET("/:id", orderHandler.GetOrder)
		orderRoutes.POST("/", orderHandler.CreateOrder)
		orderRoutes.PUT("/:id", orderHandler.UpdateOrder)
		orderRoutes.DELETE("/:id", orderHandler.DeleteOrder)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// start the server
	err := router.Run(":8001")
	if err != nil {
		panic(err)
	}
}
