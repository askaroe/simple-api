package main

import (
	"fmt"
	"github.com/askaroe/simple-api/internal/initialize"
	"github.com/askaroe/simple-api/internal/models"
)

func init() {
	initialize.LoadEnvVariables()
	initialize.ConnectToDb()
}

func main() {
	if err := initialize.DB.Migrator().DropTable(&models.User{}); err != nil {
		fmt.Println(err)
		return
	}

}
