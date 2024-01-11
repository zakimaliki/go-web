package models

import (
	"golang-test/src/config"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
}

func PostUser(item *User) *gorm.DB {
	return config.DB.Create(&item)
}

func FindEmail(input *User) []User {
	items := []User{}
	config.DB.Raw("SELECT * FROM users WHERE email = ?", input.Email).Scan(&items)
	return items
}
