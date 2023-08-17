package userhandler

import (
	"fmt"

	"github.com/gibran-abdillah/rest-jwt/models"
	"github.com/gibran-abdillah/rest-jwt/serializers"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var user models.User
	var data serializers.CreateUser

	// put result from json to data

	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": "Cant parse the json!"})
		return
	}

	result := models.DB.Where("username = ?", data.Username).First(&user)

	// if success fetching the data, it means the username already exists in database

	if result.Error == nil {
		c.JSON(200, map[string]string{"message": "Sorry, that username already taked!"})
		return
	}

	// set password and set verified to false

	user.Username = data.Username
	user.Name = data.Name
	user.Password = data.Password
	user.SetPassword()

	if result := models.DB.Create(&user); result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": "Internal server error!"})
		return
	}
	fmt.Println(data)
	c.JSON(201, gin.H{"message": "User Created!"})
}
