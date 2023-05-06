package main

import (
	"goauth/controllers"
	"goauth/initializers"

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

	r.Run()
}
