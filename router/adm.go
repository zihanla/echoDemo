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

	adm.GET("/user/page", control.UserPage)
	adm.GET("/user/del/:id", control.UserDel)
	adm.POST("/user/add", control.UserAdd)
	adm.GET("/user/get/:id", control.UserGet)
	adm.POST("/user/edit", control.UserEdit)
	adm.GET("/article/del/:id", control.ArticleDel)
}
