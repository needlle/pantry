package repositories

import (
	"pantry/api/repositories/schema"

	"gorm.io/gorm"
)

func Setup(db *gorm.DB) {
	db.AutoMigrate(&schema.Foodstuff{})
}
