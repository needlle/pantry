package api

import (
	"log"
	"pantry/api/business/services"
	"pantry/api/controllers"
	"pantry/api/repositories"
	"pantry/api/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Setup(router fiber.Router) {
	db, err := gorm.Open(sqlite.Open("data/pantry.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	foodstuffsRepository := repositories.NewFoodstuffsRepository(db)
	foodstuffsService := services.NewFoodstuffsService(foodstuffsRepository)
	foodstuffsController := controllers.NewFoodstuffsController(foodstuffsService)
	repositories.Setup(db)
	routes.Setup(router, foodstuffsController)
}
