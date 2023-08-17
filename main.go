package main

import (
	"fmt"

	"github.com/gibran-abdillah/rest-jwt/helpers"
	"github.com/gibran-abdillah/rest-jwt/models"
	"github.com/gibran-abdillah/rest-jwt/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// connecting and migrating

	models.Connect()

	result, err := helpers.Get_user(1)
	if err != nil {
		user := models.User{Username: "gibranabdillah", Verified: true, Password: "HelloWorld"}
		user.SetPassword()

		models.DB.Create(&user)
	} else {
		fmt.Println("Nil cuy", result)
	}

	router := gin.Default()
	api_router := router.Group("/api")
	// add the routers in /routers/*.go to api_router
	routers.SetupUserRoute(api_router)
	routers.InitAuthRoute(api_router)

	router.Run()

}
