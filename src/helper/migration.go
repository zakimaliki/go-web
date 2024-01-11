package helper

import (
	"golang-test/src/config"
	"golang-test/src/models"
)

func Migrate() {
	config.DB.AutoMigrate(&models.Product{})
	config.DB.AutoMigrate(&models.User{})
}
