package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"dbAdmin/logic"
)

func TableList(c echo.Context) error {
	return c.JSON(http.StatusOK, logic.NewResolveLogic(c).TableList())
}
