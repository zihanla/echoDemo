package model

import (
	"errors"
	"time"
)

type Article struct {
	Id      int64     `json:"id"`      //
	Cid     int64     `json:"cid"`     // class id
	Uid     int64     `json:"uid"`     //  user id
	Title   string    `json:"title"`   // 作者
	Origin  string    `json:"origin"`  // 来源
	Author  string    `json:"author"`  // 作者
	Content string    `json:"content"` // 内容
	Hits    int64     `json:"hits"`    // 点击量
	Utime   time.Time `json:"utime"`   // 发布时间
	Ctime   time.Time `json:"ctime"`   // 更改时间
	Status  bool      `json:"status"`  // 状态
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
	return mods, err
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
