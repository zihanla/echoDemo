package control

import (
	"demo_1/model"
	"strconv"

	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
)

// ArticlePage 文章分页
func ArticlePage(ctx echo.Context) error {
	// 创建个分页容器 来装前端传来的数据
	ipt := Page{}
	// 接收前端传过来的值 并装进ipt容器里
	err := ctx.Bind(&ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	// 判断ipt容器里的值
	if ipt.Pi < 1 {
		ipt.Pi = 1
	}
	if ipt.Ps < 1 || ipt.Ps > 50 {
		ipt.Ps = 6
	}
	// 判断文章总数
	count := model.ArticleCount()
	if count < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据"))
	}
	// 拿ipt容器里的值 去数据库查询
	mods, err := model.ArticlePage(ipt.Pi, ipt.Ps)
	if err != nil {
		return ctx.JSON(utils.Fail("查询失败", err.Error()))
	}
	return ctx.JSON(utils.Page("文章数据", mods, count))
}

// ArticleDel 删除文章
func ArticleDel(ctx echo.Context) error {
	// 从前端path路径获取参数 并转换未int64
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	// 拿id去数据库 查值
	err = model.ArticleDel(id)
	if err != nil {
		return ctx.JSON(utils.Fail("删除失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("删除成功"))
}
