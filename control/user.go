package control

import (
	"demo_1/model"
	"strconv"
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
	// 根据后台数据 给用户信息加密 再封装
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

// UserPage 用户分页
func UserPage(ctx echo.Context) error {
	// 创建一个容器装前端请求过来的数据
	ipt := Page{}
	// 获取前端发送过来的数据
	err := ctx.Bind(&ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	// 判断前端的传过来的值
	if ipt.Pi < 1 {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	if ipt.Ps < 1 || ipt.Ps > 50 {
		// 如果分页大小 小于1或大于50 就赋值分页大小为6
		ipt.Ps = 6
	}
	// 判断用户总数
	count := model.UserCount()
	if count < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据"))
	}
	// 操作数据库拿到前端要的值
	mods, err := model.UserPage(ipt.Pi, ipt.Ps)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到数据"))
	}
	return ctx.JSON(utils.Page("用户数据", mods, count))
}

// UserDel 删除用户
func UserDel(ctx echo.Context) error {
	// 从浏览器path路径中取id值 并转换为int64
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	uid, _ := ctx.Get("uid").(int64) // 用断言转换类型

	// 防止删除自己
	if uid == id {
		return ctx.JSON(utils.Fail("不能对自己操作", err.Error()))
	}

	err = model.UserDel(id)
	if err != nil {
		return ctx.JSON(utils.Fail("删除失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("删除成功"))
}

// UserAdd 添加用户
func UserAdd(ctx echo.Context) error {
	// 写个容器放前端发送的数据
	ipt := model.User{}
	// 接收从前端发送的数据 并装进容器里
	err := ctx.Bind(&ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	// 验证必要的数据 账号 用户名 密码
	if ipt.Num == "" {
		return ctx.JSON(utils.ErrIpt("用户账号不能为空"))
	}
	if ipt.Name == "" {
		return ctx.JSON(utils.ErrIpt("用户名不能为空"))
	}
	if ipt.Pass == "" {
		return ctx.JSON(utils.ErrIpt("密码不能为空"))
	}
	if model.UserExist(ipt.Num) {
		return ctx.JSON(utils.ErrIpt("账号重复"))
	}
	// TODO 小时 不准确 有时差
	ipt.Ctime = time.Now()
	// 操作数据库存入数据库
	err = model.UserAdd(&ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("添加失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("添加成功"))
}

// UserEdit 用户编辑
func UserEdit(ctx echo.Context) error {
	// 写一个容器装前端穿过来的数据
	ipt := model.User{}
	// 接收前端数据并装到ipt容器里
	err := ctx.Bind(&ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	// 验证信息
	if ipt.Name == "" {
		return ctx.JSON(utils.ErrIpt("用户名不能为空"))
	}
	// 提交数据到数据库
	err = model.UserEdit(&ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("修改失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("修改成功"))
}

// UserGet 通过id查询用户
func UserGet(ctx echo.Context) error {
	// 获取浏览器path路径的参数 并转换为int64
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	// 通过id查找数据库
	mod, err := model.UserGet(id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Succ("用户数据", mod))
}

type Page struct {
	Pi int `json:"pi"`
	Ps int `json:"ps"`
}

type login struct {
	Num  string `json:"num"`
	Pass string `json:"pass"`
}
