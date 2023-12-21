package models

import (
	"golang-test/src/config"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255)"`
	Price int
	Stock int
}

func SelectAll() *gorm.DB {
	// items := []Product{}
	// config.DB.Raw("SELECT * FROM products").Scan(&items)
	// return items
	items := []Product{}
	return config.DB.Find(&items)
}

func Select(id string) *gorm.DB {
	// items := []Product{}
	// config.DB.Raw("SELECT * FROM products WHERE id = ?", id).Scan(&items)
	// return items
	var item Product
	return config.DB.First(&item, "id = ?", id)
}

func Post(item *Product) *gorm.DB {
	// items := []Article{}
	// config.DB.Raw("INSERT INTO `products` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `price`, `stock`) VALUES(Null, NULL, NULL, NULL, ?, ?, ?)", Title, Slug, Desc).Scan(&items)
	// return items
	return config.DB.Create(&item)
}

func Updates(id string, newProduct *Product) *gorm.DB {
	// items := []Article{}
	// config.DB.Raw("UPDATE articles SET title = ?, slug = ? , description = ? WHERE id = ?", Title, Slug, Desc, Id).Scan(&items)
	// return items
	var item Product
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newProduct)
}

func Deletes(id string) *gorm.DB {
	// items := []Article{}
	// config.DB.Raw("DELETE FROM articles WHERE id = ?", Id).Scan(&items)
	// return items
	var item Product
	return config.DB.Delete(&item, "id = ?", id)
}

// var Products = []Product{
// 	Product{1, "baju", 200000, 12},
// 	Product{2, "kemeja", 100000, 8},
// 	Product{3, "jeans", 150000, 6},
// }
