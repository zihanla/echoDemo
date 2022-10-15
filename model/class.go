package model

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

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

// ClassAdd 添加
func ClassAdd(mod *Class) error {
	tx, err := Db.Begin() // 开启事务 保险

	result, err := tx.Exec("insert  into class(`name`,`desc`) values(?,?)", mod.Name, mod.Desc)
	if err != nil {
		// 回滚 ctrl + z 撤回
		tx.Rollback()
		return err
	}
	rows, _ := result.RowsAffected()
	if rows < 1 {
		// 回滚
		tx.Rollback()
		return errors.New("rows affected < 1")
	}
	// 提交 保存正确的操作
	tx.Commit()
	return nil
}

// ClassEdit 修改分类
func ClassEdit(mod *Class) error {
	tx, err := Db.Begin() // 开启事务 保险

	result, err := tx.Exec("update class set `name`=?,`desc`=? where id=?", mod.Name, mod.Desc, mod.Id)
	if err != nil {
		// 回滚 ctrl + z 撤回
		tx.Rollback()
		return err
	}
	rows, _ := result.RowsAffected()
	if rows < 1 {
		// 回滚
		tx.Rollback()
		return errors.New("rows affected < 1")
	}
	// 提交 保存正确的操作
	tx.Commit()
	return nil
}

// ClassDel 删除选中分类
func ClassDel(id int64) error {
	tx, err := Db.Begin() // 开启事务 保险
	result, err := tx.Exec("delete from class where id =?", id)
	if err != nil {
		// 回滚 ctrl + z 撤回
		tx.Rollback()
		return err
	}
	rows, _ := result.RowsAffected()
	if rows < 1 {
		// 回滚
		tx.Rollback()
		return errors.New("rows affected < 1")
	}
	// 提交 保存正确的操作
	tx.Commit()
	return nil
}

// ClassNameById 通过id 查分类名
func ClassNameById(id int64) string {
	name := ""
	Db.Get(&name, "select name from class where id = ?", id)
	return name
}

// ClassNameByIds 查询分类 集合
func ClassNameByIds(ids []int64) map[int64]string {
	query, args, _ := sqlx.In("select * from class where id in (?)", ids)
	query = Db.Rebind(query)
	mods := make([]Class, 0, len(ids))
	Db.Select(&mods, query, args...)
	mp := make(map[int64]string)
	for i := 0; i < len(mods); i++ {
		mp[mods[i].Id] = mods[i].Name
	}
	//log.Println(mp)
	return mp
}
