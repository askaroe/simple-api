package main

import (
	"github.com/askaroe/simple-api/handlers"
	initialize2 "github.com/askaroe/simple-api/internal/initialize"
	"github.com/gin-gonic/gin"
)

func init() {
	initialize2.LoadEnvVariables()
	initialize2.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.POST("/api/v1/users", handlers.CreateUser)
	r.GET("/api/v1/users", handlers.GetAllUsers)
	r.GET("/api/v1/users/:id", handlers.GetUserById)
	r.PUT("/api/v1/users/:id", handlers.UpdateUser)
	r.DELETE("/api/v1/users/:id", handlers.DeletePost)
	r.Run() // listen and serve on 0.0.0.0:8080
}
