package menu

import "restaurant-app/internal/model"

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
