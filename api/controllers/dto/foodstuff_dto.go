package dto

import "pantry/api/business/model"

type Foodstuff struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ToModel(data Foodstuff) model.Foodstuff {
	return model.Foodstuff{
		ID:   data.ID,
		Name: data.Name,
	}
}

func FromModel(entity model.Foodstuff) Foodstuff {
	return Foodstuff{
		ID:   entity.ID,
		Name: entity.Name,
	}
}
