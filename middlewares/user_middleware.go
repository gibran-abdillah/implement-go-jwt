package middlewares

import (
	"fmt"

	"github.com/gibran-abdillah/rest-jwt/helpers"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Im authorized")
		cookie, err := c.Cookie("Authorization")
		if err != nil {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		if _, err := helpers.CheckJWT(cookie); err != nil {
			c.JSON(403, gin.H{"message": "Invalid cookies token!"})
			c.Abort()
			return
		}

		c.Next() // Call this to proceed to the next middleware/handler
	}
}
