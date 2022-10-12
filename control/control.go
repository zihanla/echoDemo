package control

import "github.com/labstack/echo"

func Index(ctx echo.Context) error {
	return ctx.String(200, "hello zihan welcome use echo")
}
