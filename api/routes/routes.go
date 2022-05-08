package routes

import (
	"pantry/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(
	router fiber.Router,
	foodstuffsController controllers.FoodstuffsController,
) {
	router.Get("/foodstuffs", foodstuffsController.GetFoodstuffs)
	router.Post("/foodstuffs", foodstuffsController.CreateFoodstuff)
}
