package main

import (
	"goauth/controllers"
	"goauth/initializers"
	"goauth/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/me", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
