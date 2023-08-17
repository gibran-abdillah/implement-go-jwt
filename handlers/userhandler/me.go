package userhandler

import (
	"github.com/gibran-abdillah/rest-jwt/helpers"
	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"message": "Not authorized"})
		return
	}
	if _, err := helpers.CheckJWT(tokenString); err == nil {
		c.JSON(200, gin.H{"Hello": "World"})
	}
	c.JSON(403, gin.H{"message": err})

}
