package initializers

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
)

type Vars struct {
	PGHost     string `env:"POSTGRES_HOST,required"`
	PGUser     string `env:"POSTGRES_USER,required"`
	PGPassword string `env:"POSTGRES_PASSWORD,required"`
	PGDatabase string `env:"POSTGRES_DB,required"`
	PGPort     string `env:"POSTGRES_PORT,required"`

	SecretKey   string `env:"SECRET_KEY,required"`
	Port        string `env:"PORT,required"`
	TablePrefix string `env:"TABLE_PREFIX,required"`
}

// LoadEnvVariables loads the environment variables
func LoadEnvVariables(envFilePath string) *Vars {
	// load the environment variables from the .env file
	// parse the environment variables and return them
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Error loading %s file", envFilePath)
		panic(err)
	}

	vars := Vars{}

	err = env.Parse(&vars)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
		panic(err)
	}

	return &vars
}
