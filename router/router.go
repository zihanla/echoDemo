package router

import (
	"demo_1/control"

	"github.com/labstack/echo"
)

// 编辑模式，为true Renderer 实时加载
var debug = true

func Run() {
	app := echo.New()
	// 用于模板渲染
	app.Renderer = renderer
	// 静态文件访问
	app.Static("/static", "static")
	app.Static("/views", "views")
	// 静态页面
	app.GET("/", control.Index)
	app.GET("/login", control.LoginView)

	// admin分组 设置经过中间件的页面
	adm := app.Group("/admin", ServerHeader)
	AdmRouter(adm)

	// api 常规操作 不需要经过的中间验证的数据
	api := app.Group("/api")
	ApiRouter(api)

	// 服务端口
	app.Start(":8080")
}
