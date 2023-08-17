package userhandler

import (
	"fmt"
	"strconv"

	"github.com/gibran-abdillah/rest-jwt/helpers"
	"github.com/gibran-abdillah/rest-jwt/models"
	"github.com/gin-gonic/gin"
)

func Edit(c *gin.Context) {
	var data models.User
	var another_user models.User

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": "Invalid user id"})
		return
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": "Cant parse json"})
		return
	}
	result, err := helpers.Get_user(id)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": "Error Occured"})
		return
	}
	if data.Username != result.Username {
		if err := models.DB.Where("username = ?", data.Username).First(&another_user); err.Error == nil {
			c.AbortWithStatusJSON(403, gin.H{"message": "username already taked"})
			return
		}
		fmt.Printf("No cheating :)")
	}

	result.Name = data.Name
	result.Username = data.Username

	models.DB.Save(&result)
	c.JSON(203, result)

}
