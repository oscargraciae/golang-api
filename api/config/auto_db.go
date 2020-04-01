package config

import (
	"api-app/api/database"
	"api-app/api/models"
)

func Load() {
	db := database.Connect()

	defer db.Close()

	db.Debug().AutoMigrate(&models.User{}, &models.Business{})
}