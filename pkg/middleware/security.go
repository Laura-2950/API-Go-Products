package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Authentication manages the security by validating the token
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "token not found")
			return
		}
		if token != os.Getenv("TOKEN") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token")
			return
		}
		c.Next()
	}
}
