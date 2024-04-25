package main

import (
	"github.com/askaroe/simple-api/initialize"
	"github.com/askaroe/simple-api/models"
)

func init() {
	initialize.LoadEnvVariables()
	initialize.ConnectToDb()
}

func main() {
	err := initialize.DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}
