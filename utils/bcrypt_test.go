package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/laertkokona/crud-test/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

// TestGetHashPassword tests the GetHashPassword function
func TestGetHashPassword(t *testing.T) {
	password := "password"
	hashedPassword := GetHashPassword(password)
	assert.True(t, ComparePassword([]byte(hashedPassword), []byte(password)), "Expected password to be valid, got invalid")
}

// TestHashPassword tests the HashPassword function
func TestHashPassword(t *testing.T) {
	password := []byte("password")
	HashPassword(&password)
	assert.True(t, ComparePassword(password, []byte("password")), "Expected password to be valid, got invalid")
}

// TestComparePasswords tests the ComparePasswords function
func TestComparePasswords(t *testing.T) {
	password := "password"
	hashedPassword := GetHashPassword(password)
	val := ComparePassword([]byte(hashedPassword), []byte(password))
	assert.True(t, val, "Expected password to be valid, got invalid")
}

// TestGenerateToken tests the GenerateToken function
func TestGenerateToken(t *testing.T) {
	token := GenerateToken(models.User{Model: gorm.Model{ID: 1}, Username: "username"}, "role")
	assert.NotEqual(t, token, "", "Expected token to be valid, got invalid")
	tokenClaims, err := GetClaimsFromToken(token)
	assert.NoError(t, err, "Expected no error, got", err)
	assert.True(t, tokenClaims["exp"].(float64) > float64(time.Now().Unix()), "Expected token to be valid, got expired")
}

// TestGenerateExpiredToken tests the GenerateExpiredToken function
func TestGenerateExpiredToken(t *testing.T) {
	token := GenerateExpiredToken()
	assert.NotEqual(t, token, "", "Expected token to be valid, got invalid")
	tokenClaims, err := GetClaimsFromToken(token)
	assert.NoError(t, err, "Expected no error, got", err)
	assert.True(t, tokenClaims["exp"].(float64) <= float64(time.Now().Unix()), "Expected token to be expired, got valid")
}

// TestValidateToken tests the ValidateToken function
func TestValidateToken(t *testing.T) {
	token := GenerateExpiredToken()
	tkn, err := ValidateToken(token)
	assert.NoError(t, err, "Expected error, got none")
	// assert that tkn type is tha same as *jwt.Token
	assert.IsType(t, &jwt.Token{}, tkn, "Expected type to be *jwt.Token, got", tkn)
}

// TestGetClaimsFromToken tests the GetClaimsFromToken function
func TestGetClaimsFromToken(t *testing.T) {
	token := GenerateExpiredToken()
	tokenClaims, err := GetClaimsFromToken(token)
	assert.NoError(t, err, "Expected no error, got", err)
	assert.NotEmptyf(t, tokenClaims["exp"], "Expected token to have expiration time, got empty")
	assert.NotEmptyf(t, tokenClaims["iat"], "Expected token to have issued at time, got empty")
}

// TestGetClaimsFromToken_ValidateTokenError tests the GetClaimsFromToken function
func TestGetClaimsFromToken_ValidateTokenError(t *testing.T) {
	token := generateBadToken()
	tokenClaims, err := GetClaimsFromToken(token)
	assert.Error(t, err, "Expected no error, got", err)
	assert.Empty(t, tokenClaims, "Expected token claims to be empty, got", tokenClaims)
}

func generateBadToken() string {
	// create claims
	claims := jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 5).Unix(),
		"iat":  time.Now().Unix(),
		"sub":  0,
		"user": "",
		"role": "",
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("What is this?"))
	if err != nil {
		panic(err)
	}
	return t
}
