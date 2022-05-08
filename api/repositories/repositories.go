package repositories

import "gorm.io/gorm"

func Setup(db *gorm.DB) {
	db.AutoMigrate(&Foodstuff{})
}
