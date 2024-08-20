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

func seedDB() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	foodMenu := []MenuItem{
		{
			Name:      "Seblak",
			OrderCode: "seblak_ndower",
			Price:     10000,
		},
		{
			Name:      "Bakso",
			OrderCode: "bakso_tanpa_tepung",
			Price:     20000,
		},
	}

	drinkMenu := []MenuItem{
		{
			Name:      "Es Teh Manis",
			OrderCode: "s_teh",
			Price:     5000,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "s_jeruk",
			Price:     6000,
		},
	}

	db.Create(&foodMenu)
	db.Create(&drinkMenu)

	err = db.AutoMigrate(&MenuItem{})
	if err != nil {
		fmt.Println(err)
	}

}

func getFoodMenu(c echo.Context) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	var menuData []MenuItem

	db.Where(MenuItem{Type: MenuTypeFood}).Find(&menuData)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}

func getDrinkMenu(c echo.Context) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	var menuData []MenuItem

	db.Where(MenuItem{Type: MenuTypeDrink}).Find(&menuData)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}
func main() {
	e := echo.New()
	e.GET("/menu/food", getFoodMenu)
	e.GET("/menu/drink", getDrinkMenu)
	e.Logger.Fatal(e.Start(":1323"))
	seedDB()
}
