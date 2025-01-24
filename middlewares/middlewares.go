package middlewares

import (
	"net/http"
	"user/server/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	// Retrieve token from request
	jwtToken := c.Request.Header.Get("Authorization")
	if jwtToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized."})
		c.Abort()
	}

	// Verify token
	userJwt, err := utils.VerifyToken(jwtToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized."})
		c.Abort()
	}

	c.Set("userJwt", userJwt)
}
