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
	results, err := s.foodstuffsRepository.FindAll()
	if err != nil {
		return []*model.Foodstuff{}, err
	}
	return results, nil
}

func (s *service) CreateFoodstuff(foodstuff *model.Foodstuff) (*model.Foodstuff, error) {
	result, err := s.foodstuffsRepository.Create(foodstuff)
	if err != nil {
		return &model.Foodstuff{}, err
	}
	return result, nil
}
