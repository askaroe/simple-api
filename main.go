package main

import (
	"github.com/askaroe/simple-api/controllers"
	"github.com/askaroe/simple-api/initialize"
	"github.com/gin-gonic/gin"
)

func init() {
	initialize.LoadEnvVariables()
	initialize.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.POST("/api/v1/users", controllers.CreateUser)
	r.GET("/api/v1/users", controllers.GetAllUsers)
	r.GET("/api/v1/users/:id", controllers.GetUserById)
	r.PUT("/api/v1/users/:id", controllers.UpdateUser)
	r.DELETE("/api/v1/users/:id", controllers.DeletePost)
	r.Run() // listen and serve on 0.0.0.0:8080
}
