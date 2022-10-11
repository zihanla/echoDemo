package router

import (
	"demo_1/control"

	"github.com/labstack/echo"
)

func AdmRouter(adm *echo.Group) {
	adm.GET("/index.html", control.AdmView)
	adm.POST("/class/add", control.ClassAdd)
	adm.POST("/class/edit", control.ClassEdit)
	adm.GET("/class/del/:id", control.ClassDel)
}
