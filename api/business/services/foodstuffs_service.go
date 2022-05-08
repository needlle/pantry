package services

import (
	"pantry/api/business/model"
	"pantry/api/repositories"
)

type service struct {
	foodstuffsRepository repositories.FoodstuffsRepository
}

type FoodstuffsService interface {
	GetFoodstuffs() ([]*model.Foodstuff, error)
	CreateFoodstuff(*model.Foodstuff) (*model.Foodstuff, error)
}

func NewFoodstuffsService(foodstuffsRepository repositories.FoodstuffsRepository) FoodstuffsService {
	return &service{
		foodstuffsRepository,
	}
}

func (s *service) GetFoodstuffs() ([]*model.Foodstuff, error) {
	foodstuffs, err := s.foodstuffsRepository.FindAll()
	if err != nil {
		return []*model.Foodstuff{}, err
	}
	return foodstuffs, nil
}

func (s *service) CreateFoodstuff(data *model.Foodstuff) (*model.Foodstuff, error) {
	foodstuff, err := s.foodstuffsRepository.Create(data)
	if err != nil {
		return &model.Foodstuff{}, err
	}
	return foodstuff, nil
}
