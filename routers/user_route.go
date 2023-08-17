package routers

import (
	"github.com/gibran-abdillah/rest-jwt/handlers/userhandler"
	"github.com/gibran-abdillah/rest-jwt/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupUserRoute(router *gin.RouterGroup) {
	// we need to modify the value of *gin.Engine, which is from `gin.RouterGroup()`

	// create new group for users and put middleware
	// the endpoint will be like /api/users/<endpoint>

	users := router.Group("/users")
	users.Use(middlewares.AuthMiddleware())

	// Only authorized user can handle requests
	{
		users.GET("", userhandler.List)
		users.POST("", userhandler.Create)
		users.GET("/me", userhandler.GetMe)
		users.GET("/:id", userhandler.Get)
		users.PUT("/:id", userhandler.Edit)
		users.DELETE("/:id", userhandler.Delete)
	}

}
