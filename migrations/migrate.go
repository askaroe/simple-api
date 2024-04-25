package migrations

import (
	"github.com/askaroe/simple-api/initialize"
	"github.com/askaroe/simple-api/models"
)

func init() {
	initialize.LoadEnvVariables()
	initialize.ConnectToDb()
}

func main() {
	initialize.DB.AutoMigrate(&models.User{})
}
