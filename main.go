package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MenuType string

const (
	MenuTypeFood  = "food"
	MenuTypeDrink = "drink"
)

const (
	dsn = "host=localhost user=postgres password=adminadmin dbname=restaurant-app port=5432 sslmode=disable TimeZone=Asia/Jakarta"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
	Type      MenuType
}

func seedDB(db *gorm.DB) {

	foodMenu := []MenuItem{
		{
			Name:      "Seblak",
			OrderCode: "seblak_ndower",
			Price:     10000,
			Type:      MenuTypeFood,
		},
		{
			Name:      "Bakso",
			OrderCode: "bakso_tanpa_tepung",
			Price:     20000,
			Type:      MenuTypeFood,
		},
	}

	drinkMenu := []MenuItem{
		{
			Name:      "Es Teh Manis",
			OrderCode: "s_teh",
			Price:     5000,
			Type:      MenuTypeDrink,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "s_jeruk",
			Price:     6000,
			Type:      MenuTypeDrink,
		},
	}

	db.Create(&foodMenu)
	db.Create(&drinkMenu)

	err := db.AutoMigrate(&MenuItem{})
	if err != nil {
		fmt.Println(err)
	}

}

func getMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	var menuData []MenuItem
	db.Where(MenuItem{Type: MenuType(menuType)}).Find(&menuData)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})

}

func main() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	seedDB(db)
	e := echo.New()
	e.GET("/menu", getMenu)
	e.Logger.Fatal(e.Start(":1323"))

}
