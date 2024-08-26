package main

import (
	"restaurant-app/internal/database"
	"restaurant-app/internal/delivery/rest"
	menu "restaurant-app/internal/repository"
	rUsecase "restaurant-app/internal/usecase/resto"

	"github.com/labstack/echo"
)

const (
	dbAddress = "host=localhost user=postgres password=adminadmin dbname=restaurant-app port=5432 sslmode=disable TimeZone=Asia/Jakarta"
)

func main() {
	e := echo.New()

	db := database.GetDB(dbAddress)
	menuRepo := menu.NewRepository(db)

	restoUsecase := rUsecase.GetUsecase(menuRepo)

	h := rest.NewHandler(restoUsecase)
	rest.LoadRouter(e, h)

	e.Logger.Fatal(e.Start(":1323"))

}
