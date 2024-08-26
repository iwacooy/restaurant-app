package resto

import (
	"restaurant-app/internal/model"
	menu "restaurant-app/internal/repository"
)

type restoUsecase struct {
	menuRepo menu.Repository
}

func GetUsecase(menuRepo menu.Repository) Usecase {
	return &restoUsecase{menuRepo: menuRepo}
}

func (r *restoUsecase) GetMenu(menuType string) ([]model.MenuItem, error) {
	return r.menuRepo.GetMenu(menuType)
}
