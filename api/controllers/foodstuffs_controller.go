package controllers

import (
	"pantry/api/business/model"
	"pantry/api/business/services"

	"github.com/gofiber/fiber/v2"
)

type controller struct {
	foodstuffsService services.FoodstuffsService
}

type FoodstuffsController interface {
	GetFoodstuffs(*fiber.Ctx) error
	CreateFoodstuff(*fiber.Ctx) error
}

func NewFoodstuffsController(
	foodstuffsService services.FoodstuffsService,
) FoodstuffsController {
	return &controller{
		foodstuffsService,
	}
}

func (c *controller) GetFoodstuffs(ctx *fiber.Ctx) error {
	foodstuffs, err := c.foodstuffsService.GetFoodstuffs()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(foodstuffs)
}

func (c *controller) CreateFoodstuff(ctx *fiber.Ctx) error {
	foodstuff, err := c.foodstuffsService.CreateFoodstuff(&model.Foodstuff{Name: "Test"})
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(foodstuff)
}
