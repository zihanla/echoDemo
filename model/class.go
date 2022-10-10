package model

// Class 结构体
type Class struct {
	Id   int64  `json:"id"`   // id
	Name string `json:"name"` // 新闻名
	Desc string `json:"desc"` //  新闻描述
}

// ClassPage 分页
func ClassPage(pi, ps int) ([]Class, error) {
	mods := make([]Class, 0, ps)
	err := Db.Select(&mods, "select * from class limit ?,?", (pi-1)*ps, ps)
	return mods, err
}

// ClassCount 查询总数
func ClassCount() int {
	count := 0
	Db.Get(&count, "select count(id) as count from class")
	return count
}

// ClassAll 查询所有
func ClassAll() ([]Class, error) {
	mods := make([]Class, 0, 8)
	err := Db.Select(&mods, "select * from class")
	return mods, err
}

// ClassGet 查询某一条
func ClassGet(id int64) (*Class, error) {
	mod := &Class{}
	err := Db.Get(mod, "select * from class where id = ? limit 1", id)
	return mod, err
}

// 添加
// 修改
// 删除
