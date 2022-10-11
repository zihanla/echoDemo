package router

import (
	"demo_1/model"
	"io"
	"text/template"

	"github.com/zxysilent/utils"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
)

// TemplateRenderer 模板结构体
type TemplateRenderer struct {
	templates *template.Template
}

// Render 实现模板方法
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	//if viewContext, isMap := data.(map[string]interface{}); isMap {
	//	viewContext["reverse"] = c.Echo().Reverse
	//}

	// 静态文件热加载
	if debug {
		t.templates = template.Must(template.ParseFiles("./views/login.html", "./views/index.html"))
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

// 定义一个全局 模板变量
// 只有在运行时才会赋值
var renderer = &TemplateRenderer{
	// 添加指定模板 进行静态渲染
	templates: template.Must(template.ParseFiles("./views/login.html", "./views/index.html")),
}

// ServerHeader 中间件
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Set(echo.HeaderServer, "Echo/666")
		// TODO 还需要解决直接从浏览器中拿去token值
		tokenString := ctx.FormValue("token")
		//log.Println(tokenString) // 打印token值
		claims := model.UserClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("123"), nil
		})

		if err == nil && token.Valid {
			// 验证成功
			return next(ctx)
		} else {
			// 验证失败
			return ctx.JSON(utils.ErrToken("token 验证失败"))
		}
	}
}
