package middleware

import "github.com/gin-gonic/gin"

// CSRFMiddleware is a middleware that checks if the request has a valid CSRF token
func CSRFMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the csrf token from the request header
		// check if the csrf token is empty
		// check if the csrf token is not equal to the csrf token in the session
		// next
		csrfToken := c.GetHeader("X-CSRF-Token")
		if csrfToken == "" {
			c.AbortWithStatusJSON(400, gin.H{"error": "missing csrf token"})
			return
		}
		if csrfToken != c.MustGet("csrf_token").(string) {
			c.AbortWithStatusJSON(400, gin.H{"error": "invalid csrf token"})
			return
		}
		c.Next()
	}
}
