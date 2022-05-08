package schema

import (
	"pantry/api/business/model"

	"gorm.io/gorm"
)

type Foodstuff struct {
	gorm.Model
	ID   uint `gorm:"primaryKey"`
	Name string
}

func ToModel(record Foodstuff) model.Foodstuff {
	return model.Foodstuff{
		ID:   int(record.ID),
		Name: record.Name,
	}
}

func FromModel(entity model.Foodstuff) Foodstuff {
	return Foodstuff{
		ID:   uint(entity.ID),
		Name: entity.Name,
	}
}
