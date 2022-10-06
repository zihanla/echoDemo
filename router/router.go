package router

import (
	"demo_1/control"
	"github.com/labstack/echo"
)

func Run() {
	app := echo.New()
	// 静态文件访问
	app.Static("/static", "static")
	app.Static("/views", "views")
	// 页面
	app.GET("/", control.Index)
	// 方法
	app.POST("/api/login", control.UserLogin) // 用户登录
	app.Start(":8080")
}
