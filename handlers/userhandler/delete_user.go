package userhandler

import (
	"strconv"

	"github.com/gibran-abdillah/rest-jwt/models"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var user models.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": "Invalid user id"})
		return
	}
	if err := models.DB.First(&user, id); err.Error != nil {
		c.AbortWithStatusJSON(404, gin.H{"message": "Invalid user"})
		return
	}
	if err := models.DB.Delete(&user, id); err.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": "internal server error!"})
		return
	}
	c.JSON(200, gin.H{"message": "deleted, bye!"})

}
