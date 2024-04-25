package main

import (
	initialize2 "github.com/askaroe/simple-api/internal/initialize"
	"github.com/askaroe/simple-api/internal/models"
)

func init() {
	initialize2.LoadEnvVariables()
	initialize2.ConnectToDb()
}

func main() {
	err := initialize2.DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}
