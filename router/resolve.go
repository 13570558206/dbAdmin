package router

import (
	"github.com/labstack/echo/v4"

	"dbAdmin/controller"
)

func AddResolveRoute(e *echo.Echo) {

	user := e.Group("/resolve")

	user.GET("/list", controller.TableList)
}
