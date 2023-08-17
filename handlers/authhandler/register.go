package authhandler

import (
	"github.com/gibran-abdillah/rest-jwt/models"
	"github.com/gibran-abdillah/rest-jwt/serializers"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var data serializers.CreateUser
	var user models.User

	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": "Invalid json request!"})
		return
	}
	if result := models.DB.Where("username = ?", data.Username).First(&user); result.Error == nil {
		c.AbortWithStatusJSON(403, gin.H{"message": "username already exists!"})
		return
	}
	user.Username = data.Username
	user.Password = data.Password
	user.Name = data.Name

	user.SetPassword()

	if result := models.DB.Create(&user); result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": "Internal server error!"})
		return
	}

	c.JSON(201, gin.H{"message": "Created", "user": user})
}
