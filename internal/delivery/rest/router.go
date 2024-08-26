package rest

import "github.com/labstack/echo"

func LoadRouter(e *echo.Echo, h *handler) {
	e.GET("/menu", h.GetMenu)
}
