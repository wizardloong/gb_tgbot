package repository

import (
	"telegram-bot/internal/domain"

	"gorm.io/gorm"
)

type GoodsRepository struct {
	db *gorm.DB
}

func NewGoodsRepository(db *gorm.DB) *GoodsRepository {
	return &GoodsRepository{db: db}
}

// Добавление товара
func (r *GoodsRepository) Create(goods *domain.Goods) error {
	return r.db.Create(goods).Error
}

// Получение товара по ID
func (r *GoodsRepository) GetByID(id uint) (*domain.Goods, error) {
	var goods domain.Goods
	err := r.db.First(&goods, id).Error
	return &goods, err
}

// Обновление товара
func (r *GoodsRepository) Update(goods *domain.Goods) error {
	return r.db.Save(goods).Error
}

// Удаление товара
func (r *GoodsRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Goods{}, id).Error
}

// Получение всех товаров
func (r *GoodsRepository) GetAll() ([]domain.Goods, error) {
	var goods []domain.Goods
	err := r.db.Find(&goods).Error
	return goods, err
}
