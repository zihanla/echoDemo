package control

import "github.com/labstack/echo"

// LoginView 登录页面
func LoginView(ctx echo.Context) error {
	return ctx.Render(200, "login.html", nil)
}

// AdmView 后台页面
func AdmView(ctx echo.Context) error {
	return ctx.Render(200, "index.html", nil)
}
