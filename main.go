package main

import (
	"github.com/laertkokona/crud-test/database"
	"github.com/laertkokona/crud-test/initializers"
	"github.com/laertkokona/crud-test/routes"
	"gorm.io/gorm"
	//"./docs"
)

var DB *gorm.DB
var vars *initializers.Vars

// init function
func init() {
	// Load environment variables
	vars := initializers.LoadEnvVariables(".env")

	// Connect to database
	DB = database.Connect(vars)
}

// main function
// @title Swagger Crud-Test API
// @version 0.1
// @description This is a sample CRUD server.

// @contact.name Laert Kokona
// @contact.url http://github.com/laertkokona
// @contact.email laerti98@gmail.com

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:8001
// @BasePath /
func main() {

	// Setup routes
	routes.SetupRoutes(DB)

	//docs.SwaggerInfo.Schemes = []string{"http", "https"}

}
