package control

import (
	"demo_1/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
)

// UserLogin 登录逻辑
func UserLogin(ctx echo.Context) error {
	// 创建一个结构体 接收请求中的值
	ipt := login{}
	// 从请求中拿值
	err := ctx.Bind(&ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	// 根据请求中的提交的账号 去数据库中查找是否有对应 账号
	mod, err := model.UserLogin(ipt.Num)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("用户名错误", err.Error()))
	}
	// 找到对于账号后 验证密码
	if mod.Pass != ipt.Pass {
		return ctx.JSON(utils.ErrIpt("密码错误"))
	}

	claims := model.UserClaims{
		Id:   mod.Id,
		Num:  mod.Num,
		Name: mod.Name,
		StandardClaims: jwt.StandardClaims{
			// 过期时间 = 现在时间 + 两小时 // Unix 时间戳 从计算机建立开始到现在
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
	}
	// 加密算法 加密数据
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("123"))
	//fmt.Printf("%v %v", ss, err)

	return ctx.JSON(utils.Succ("登录成功", ss))
}

type login struct {
	Num  string `json:"num"`
	Pass string `json:"pass"`
}
