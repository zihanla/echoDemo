package router

import (
	"demo_1/control"

	"github.com/labstack/echo"
)

func ApiRouter(api *echo.Group) {
	api.POST("/login", control.UserLogin)         // 用户登录
	api.GET("/class/all", control.ClassAll)       // 查询所有分类
	api.GET("/class/page", control.ClassPage)     // 查询所有分类
	api.GET("/class/get/:id", control.ClassGet)   // 查询指定id分类
	api.GET("/article/page", control.ArticlePage) // 查询文章
}
