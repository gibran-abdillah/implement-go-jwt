package userhandler

import (
	"strconv"

	"github.com/gibran-abdillah/rest-jwt/helpers"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": "invalid user id"})
		return
	}

	user, err := helpers.Get_user(id)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"message": "user not found"})
		return
	}
	c.JSON(200, user)

}
