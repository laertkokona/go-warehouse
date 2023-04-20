package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/laertkokona/crud-test/helpers"
	"github.com/laertkokona/crud-test/utils"
	"net/http"
	"strings"
)

// AuthMiddleware is a middleware that checks for a valid JWT token
func AuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the authorization header from the request
		// check if the authorization header is empty
		// split the authorization header into an array of strings using the space as the delimiter
		// check if the array has a length of 2
		// validate the token
		// create a map of claims
		// set the role name to the role claim in the token
		// check if the role claim is empty or if the role claim is not equal to the role name in the token
		// next
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			helpers.FailedResponse(c, http.StatusUnauthorized, "missing authorization header", nil)
			c.Abort()
			return
		}
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			helpers.FailedResponse(c, http.StatusUnauthorized, "invalid token", nil)
			c.Abort()
			return
		}
		token, err := utils.ValidateToken(bearerToken[1])
		if err != nil {
			helpers.FailedResponse(c, http.StatusUnauthorized, "invalid token", err.Error())
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			helpers.FailedResponse(c, http.StatusUnauthorized, "invalid token", nil)
			c.Abort()
			return
		}
		role, ok := claims["role"].(string)
		if !ok {
			helpers.FailedResponse(c, http.StatusUnauthorized, "missing role claim", nil)
			c.Abort()
			return
		}
		if len(roles) > 0 && !contains(roles, role) {
			helpers.FailedResponse(c, http.StatusForbidden, "forbidden", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}

// contains checks if a string is in an array of strings
func contains(roles []string, role string) bool {
	// range through the array of strings
	// check if the string is equal to the role name
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}
