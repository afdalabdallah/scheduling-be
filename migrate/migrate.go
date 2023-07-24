package main

import (
	"log"

	"github.com/afdalabdallah/backend-web/initializers"
	"github.com/afdalabdallah/backend-web/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabese()
}

func main() {

	initializers.DB.AutoMigrate(&models.Rumpun{})
	initializers.DB.AutoMigrate(&models.Matkul{})
	initializers.DB.AutoMigrate(&models.Dosen{})
	log.Println("Database migrated")
}
