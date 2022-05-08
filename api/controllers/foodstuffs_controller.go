package controllers

import (
	"pantry/api/business/services"
	"pantry/api/controllers/dto"

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
	results := []dto.Foodstuff{}
	for _, item := range foodstuffs {
		foodstuff := dto.FromModel(*item)
		results = append(results, foodstuff)
	}
	return ctx.JSON(results)
}

func (c *controller) CreateFoodstuff(ctx *fiber.Ctx) error {
	data := dto.Foodstuff{}
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	foodstuff := dto.ToModel(data)
	result, err := c.foodstuffsService.CreateFoodstuff(&foodstuff)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(dto.FromModel(*result))
}
