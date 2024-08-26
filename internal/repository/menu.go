package menu

import (
	"restaurant-app/internal/model"

	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *menuRepo {
	return &menuRepo{db: db}
}

func (r *menuRepo) GetMenu(menuType string) ([]model.MenuItem, error) {
	var menuData []model.MenuItem
	if err := r.db.Where(model.MenuItem{Type: model.MenuType(menuType)}).Find(&menuData).Error; err != nil {
		return nil, err
	}

	return menuData, nil
}
