package userhandler

import (
	"github.com/gibran-abdillah/rest-jwt/models"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(200, users)
}
