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

func (r *ItemRepository) GetByID(id int) (*models.Iten, error) {
	var item models.Iten
	if err := config.DB.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ItemRepository) GetByCode(code string) (*models.Iten, error) {
	var item models.Iten
	if err := config.DB.Where("codigo = ?", code).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ItemRepository) Create(item *models.Iten) (*models.Iten, error) {
	if err := config.DB.Create(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (r *ItemRepository) Update(item *models.Iten) error {
	return config.DB.Save(item).Error
}

func (r *ItemRepository) Delete(id int) error {
	return config.DB.Delete(&models.Iten{}, id).Error
}
