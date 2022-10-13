package control

import "github.com/labstack/echo"

func Index(ctx echo.Context) error {
	//return ctx.String(200, "hello zihan welcome use echo")
	// 重定向到登录界面
	return ctx.Redirect(302, "/login")
}
