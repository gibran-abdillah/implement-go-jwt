package authhandler

import (
	"github.com/gibran-abdillah/rest-jwt/helpers"
	"github.com/gibran-abdillah/rest-jwt/models"
	"github.com/gibran-abdillah/rest-jwt/serializers"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var stored_data serializers.UserInput
	var user models.User
	if err := c.ShouldBindJSON(&stored_data); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": "Failed to parse json"})
		return
	}
	result := models.DB.First(&user, "username = ?", stored_data.Username)
	if result.Error != nil {
		c.JSON(404, gin.H{"message": "User not found? wanna register?"})
		return
	}
	if !user.ValidatePassword(stored_data.Password) {
		c.JSON(200, gin.H{"message": "NO BROO!"})
		return
	}
	token_string, err := helpers.GenerateJwtToken(user.Username)

	if err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}

	c.SetCookie("Authorization", token_string, 3600, "/", "", false, true)
	c.JSON(200, gin.H{"message": "Success Logged in!", "token": token_string})

	// generate jwt token

}
