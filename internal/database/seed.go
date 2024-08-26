package database

import (
	"restaurant-app/internal/model"
	"restaurant-app/internal/model/constant"

	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {

	foodMenu := []model.MenuItem{
		{
			Name:      "Seblak",
			OrderCode: "seblak_ndower",
			Price:     10000,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Bakso",
			OrderCode: "bakso_tanpa_tepung",
			Price:     20000,
			Type:      constant.MenuTypeFood,
		},
	}

	drinkMenu := []model.MenuItem{
		{
			Name:      "Es Teh Manis",
			OrderCode: "s_teh",
			Price:     5000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "s_jeruk",
			Price:     6000,
			Type:      constant.MenuTypeDrink,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)

	}

}
