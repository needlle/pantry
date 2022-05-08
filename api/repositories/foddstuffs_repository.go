package repositories

import (
	"pantry/api/business/model"
	"pantry/api/repositories/schema"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

type FoodstuffsRepository interface {
	FindAll() ([]*model.Foodstuff, error)
	Create(foodstuff *model.Foodstuff) (*model.Foodstuff, error)
	// Update(foodstuff *model.Foodstuff, id int) (*model.Foodstuff, error)
	// FindByID(username string) (*model.Foodstuff, error)
}

func NewFoodstuffsRepository(db *gorm.DB) FoodstuffsRepository {
	return &repository{
		DB: db,
	}
}

func (r repository) FindAll() ([]*model.Foodstuff, error) {
	records := []*schema.Foodstuff{}
	results := []*model.Foodstuff{}
	r.DB.Find(&records)
	for _, item := range records {
		foodstuff := schema.ToModel(*item)
		results = append(results, &foodstuff)
	}
	return results, nil
}

func (r repository) Create(foodstuff *model.Foodstuff) (*model.Foodstuff, error) {
	record := schema.FromModel(*foodstuff)
	if err := r.DB.Create(&record).Error; err != nil {
		return &model.Foodstuff{}, err
	}
	result := schema.ToModel(record)
	return &result, nil
}
