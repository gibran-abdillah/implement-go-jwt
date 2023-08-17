package routers

import (
	"github.com/gibran-abdillah/rest-jwt/handlers/authhandler"
	"github.com/gin-gonic/gin"
)

func InitAuthRoute(r *gin.RouterGroup) {
	r.POST("/auth/login", authhandler.Login)
	r.POST("/auth/register", authhandler.Register)

}
