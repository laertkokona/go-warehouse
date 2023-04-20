package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/laertkokona/crud-test/models"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

// GetHashPassword hash slice of bytes password and returns it
func GetHashPassword(password string) string {
	// hash password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

// HashPassword hash slice of bytes password
func HashPassword(password *[]byte) {
	// hash password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword(*password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	*password = hashedPassword
}

// ComparePassword compare hashed password with plain text password
func ComparePassword(hashedPassword, password []byte) bool {
	return bcrypt.CompareHashAndPassword(hashedPassword, password) == nil
}

// GenerateToken generates token with claims for user id
func GenerateToken(user models.User, role string) string {
	// create claims
	claims := jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 5).Unix(),
		"iat":  time.Now().Unix(),
		"sub":  user.ID,
		"user": user.Username,
		"role": role,
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		panic(err)
	}
	return t
}

// GenerateExpiredToken generates token with claims for user id and expiration time set to now
func GenerateExpiredToken() string {
	now := time.Now().Unix()
	// create claims
	claims := jwt.MapClaims{
		"exp": now,
		"iat": now,
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		panic(err)
	}
	return t
}

// ValidateToken validate the given token
func ValidateToken(token string) (*jwt.Token, error) {
	// 2nd arg function return secret key after checking if the signing method is HMAC and returned key is used by 'Parse' to decode the token)
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// nil secret key
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

// GetClaimsFromToken get claims from token
func GetClaimsFromToken(token string) (jwt.MapClaims, error) {
	// validate token
	t, err := ValidateToken(token)
	if err != nil {
		return nil, err
	}
	// get claims
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
