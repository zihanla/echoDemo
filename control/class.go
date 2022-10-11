package control

import (
	"demo_1/model"
	"strconv"

	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
)

// ClassAll 查询所有分类
func ClassAll(ctx echo.Context) error {
	mods, err := model.ClassAll()
	if err != nil {
		ctx.JSON(utils.Fail("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Succ("分类数据", mods))
}

// ClassPage 分页查询
func ClassPage(ctx echo.Context) error {
	pi, err := strconv.Atoi(ctx.FormValue("pi"))
	if err != nil {
		return ctx.JSON(utils.ErrIpt("分页索引错误", err.Error()))
	}
	if pi < 1 {
		return ctx.JSON(utils.ErrIpt("分页索引范围错误"))
	}
	ps, err := strconv.Atoi(ctx.FormValue("ps"))
	if err != nil {
		return ctx.JSON(utils.ErrIpt("分页大小错误", err.Error()))
	}
	if ps < 1 || ps > 50 {
		// 如果分页大小 小于1或者大于50 赋值默认值为6
		ps = 6
	}
	count := model.ClassCount()
	if count < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到总数数据"))
	}
	mods, err := model.ClassPage(pi, ps)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Page("分类分页数据", mods, count))
}

// ClassAdd 添加分类
func ClassAdd(ctx echo.Context) error {
	// 创建class容器对象
	ipt := &model.Class{}
	// 接收客户端传过来的值
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	if ipt.Name == "" {
		return ctx.JSON(utils.ErrIpt("分类名称不能为空"))
	}
	// 把数据存入 数据库
	err = model.ClassAdd(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("添加失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("添加成功"))
}

// ClassEdit 添加分类
func ClassEdit(ctx echo.Context) error {
	// 创建class容器对象
	ipt := &model.Class{}
	// 接收客户端传过来的值
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	if ipt.Name == "" {
		return ctx.JSON(utils.ErrIpt("分类名称不能为空"))
	}
	// 把数据存入 数据库
	err = model.ClassEdit(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("修改失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("修改成功"))
}

// ClassDel 删除分类
func ClassDel(ctx echo.Context) error {
	// 通过param方法获取浏览器中的path路径参数 并把获取到的值转换为int64 类型
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	err = model.ClassDel(id)
	if err != nil {
		return ctx.JSON(utils.Fail("删除失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("删除成功"))
}

// ClassGet 查询分类
func ClassGet(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	mod, err := model.ClassGet(id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Succ("分类数据", mod))
}
