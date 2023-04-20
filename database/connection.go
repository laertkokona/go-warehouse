package database

import (
	"fmt"
	"github.com/laertkokona/crud-test/initializers"
	"github.com/laertkokona/crud-test/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Connect connects to the postgres database by using the environment variables and returns the connection
func Connect(vars *initializers.Vars) *gorm.DB {
	// connect to the database using the environment variables
	// Migrate the models to the database
	connection, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", vars.PGHost, vars.PGUser, vars.PGPassword, vars.PGDatabase, vars.PGPort)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   vars.TablePrefix, // schema name
			SingularTable: false,
		}})

	if err != nil {
		panic("Could not connect to database")
	}

	Migrate(connection)
	return connection

}

// Migrate migrates the models to the database
func Migrate(connection *gorm.DB) {
	err := connection.AutoMigrate(&models.Role{})
	if err != nil {
		panic(err)
	}
	err = connection.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
	err = connection.AutoMigrate(&models.Item{})
	if err != nil {
		panic(err)
	}
	err = connection.AutoMigrate(&models.Truck{})
	if err != nil {
		panic(err)
	}
	err = connection.AutoMigrate(&models.OrderItem{})
	if err != nil {
		panic(err)
	}
	err = connection.AutoMigrate(&models.Order{})
	if err != nil {
		panic(err)
	}
}
