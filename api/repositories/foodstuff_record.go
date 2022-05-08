package repositories

import (
	"pantry/api/business/model"

	"gorm.io/gorm"
)

type Foodstuff struct {
	gorm.Model
	ID   uint `gorm:"primaryKey"`
	Name string
}

func toModel(record Foodstuff) model.Foodstuff {
	return model.Foodstuff{
		ID:   int(record.ID),
		Name: record.Name,
	}
}

func fromModel(record model.Foodstuff) Foodstuff {
	return Foodstuff{
		ID:   uint(record.ID),
		Name: record.Name,
	}
}
