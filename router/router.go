package router

import (
	"github.com/labstack/echo/v4"
)

func Register() *echo.Echo {

	e := echo.New()

	AddResolveRoute(e)

	return e
}
