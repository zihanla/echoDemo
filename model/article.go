package model

import (
	"errors"
	"time"
)

type Article struct {
	Id        int64     `json:"id"`               //
	Cid       int64     `json:"cid"`              // class id
	ClassName string    `json:"className" db:"-"` // tag - 表示忽略 // 分类名称
	Uid       int64     `json:"uid"`              //  user id
	Title     string    `json:"title"`            // 作者
	Origin    string    `json:"origin"`           // 来源
	Author    string    `json:"author"`           // 作者
	Content   string    `json:"content"`          // 内容
	Hits      int64     `json:"hits"`             // 点击量
	Utime     time.Time `json:"utime"`            // 发布时间
	Ctime     time.Time `json:"ctime"`            // 更改时间
	Status    bool      `json:"status"`           // 状态
}

// ArticleCount 文章总数
func ArticleCount() int {
	count := 0
	Db.Get(&count, "select count(*) from article")
	return count
}

// ArticlePage 文章分页
func ArticlePage(pi, ps int) ([]Article, error) {
	mods := make([]Article, 0, ps) // 创建一个ps大小的容器，现在使用
	err := Db.Select(&mods, "select * from article order by id desc limit ?,?", (pi-1)*ps, ps)
	cids := make([]int64, 0, ps)
	// 把class中的name数据拿过来 赋值给ClassName 优化查询
	for i := 0; i < len(mods); i++ {
		//mods[i].ClassName = ClassNameById(mods[i].Cid)
		if !inOf(mods[i].Cid, cids) {
			cids = append(cids, mods[i].Cid)
		}
	}
	//log.Println(cids)
	mp := ClassNameByIds(cids)
	for i := 0; i < len(mods); i++ {
		mods[i].ClassName = mp[mods[i].Cid]
		mods[i].Content = ""
	}
	return mods, err
}

func inOf(dst int64, arr []int64) bool {
	for i := 0; i < len(arr); i++ {
		if dst == arr[i] {
			return true
		}
	}
	return false
}

// ArticleDel 文章删除
func ArticleDel(id int64) error {
	tx, _ := Db.Begin() // 开启事务 保险柜
	result, err := tx.Exec("delete from article where id = ?", id)
	if err != nil {
		tx.Rollback() // 回滚 撤回操作
		return err
	}
	// 受影响的行数
	rows, _ := result.RowsAffected()
	if rows < 1 {
		tx.Rollback()
		return errors.New("rows affected < 1")
	}
	tx.Commit() // 操作没问题 提交操作
	return nil
}

// ArticleAdd 添加新闻
func ArticleAdd(mod *Article) error {
	tx, _ := Db.Beginx() // 开启事务 保险柜
	result, err := tx.NamedExec("insert into article (title,author,cid,content,hits,ctime,utime,origin,status,uid) values (:title,:author,:cid,:content,:hits,:ctime,:utime,:origin,:status,:uid)", mod)
	if err != nil {
		tx.Rollback() // 回滚 撤回之前操作
		return err
	}
	// 判断收影响行数
	rows, _ := result.RowsAffected()
	if rows < 1 {
		return errors.New("rows affected < 1")
	}
	tx.Commit() // 没有错误 提交操作
	return nil
}
