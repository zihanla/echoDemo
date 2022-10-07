package router

import (
	"demo_1/control"

	"github.com/labstack/echo"
)

// 编辑模式，为true Renderer 实时加载
var debug = true

func Run() {
	app := echo.New()
	app.Renderer = renderer
	// 静态文件访问
	app.Static("/static", "static")
	app.Static("/views", "views")
	// 静态页面
	app.GET("/", control.Index)
	app.GET("/login", control.LoginView)
	// admin分组 设置经过中间件的页面
	adm := app.Group("/admin", ServerHeader)
	adm.GET("/index.html", control.AdmView)
	// 方法
	app.POST("/api/login", control.UserLogin) // 用户登录
	app.Start(":8080")
}
