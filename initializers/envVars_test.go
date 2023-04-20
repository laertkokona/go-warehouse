package initializers

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var vars *Vars

// TestLoadEnvVariables tests the LoadEnvVariables function
func TestLoadEnvVariables(t *testing.T) {
	// Define the expected values for the environment variables
	expectedVars := &Vars{
		PGHost:     "localhost",
		PGUser:     "postgres",
		PGPassword: "postgres",
		PGDatabase: "test",
		PGPort:     "5432",

		SecretKey:   "secret",
		Port:        "8001",
		TablePrefix: "test-go-warehouse.",
	}

	m := map[string]string{
		"POSTGRES_HOST":     expectedVars.PGHost,
		"POSTGRES_USER":     expectedVars.PGUser,
		"POSTGRES_PASSWORD": expectedVars.PGPassword,
		"POSTGRES_DB":       expectedVars.PGDatabase,
		"POSTGRES_PORT":     expectedVars.PGPort,
		"SECRET_KEY":        expectedVars.SecretKey,
		"PORT":              expectedVars.Port,
		"TABLE_PREFIX":      expectedVars.TablePrefix,
	}

	// Path to the .env.test file
	testEnvPath := "../.env.test"

	// Create a new .env file for the test
	err := godotenv.Write(m, testEnvPath)
	require.NoError(t, err)

	// Remove the .env file at the end of the test
	defer func() {
		err := os.Remove(testEnvPath)
		assert.NoError(t, err)
	}()

	// Call the function being tested
	vars := LoadEnvVariables(testEnvPath)

	// Use assert to check that the function returns the expected values
	assert.Equal(t, expectedVars.PGHost, vars.PGHost)
	assert.Equal(t, expectedVars.PGUser, vars.PGUser)
	assert.Equal(t, expectedVars.PGDatabase, vars.PGDatabase)
	assert.Equal(t, expectedVars.PGPassword, vars.PGPassword)
	assert.Equal(t, expectedVars.PGPort, vars.PGPort)
	assert.Equal(t, expectedVars.SecretKey, vars.SecretKey)
	assert.Equal(t, expectedVars.Port, vars.Port)
}
