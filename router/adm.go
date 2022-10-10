package router

import (
	"demo_1/control"

	"github.com/labstack/echo"
)

func AdmRouter(adm *echo.Group) {
	adm.GET("/index.html", control.AdmView)
}
