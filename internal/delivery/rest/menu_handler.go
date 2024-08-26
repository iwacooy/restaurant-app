package rest

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *handler) GetMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")

	menuData, err := h.restoUsecase.GetMenu(menuType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData})

}
