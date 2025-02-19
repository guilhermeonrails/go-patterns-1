package repositories

import (
	"myapi/internal/config"
	"myapi/internal/models"
)

type ItemRepository struct{}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{}
}

func (r *ItemRepository) ListAll() ([]models.Iten, error) {
	var items []models.Iten
	if err := config.DB.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
