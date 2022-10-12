package model

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// User 表
type User struct {
	Id     int       `json:"id"`     // id
	Num    string    `json:"num"`    // 账号
	Name   string    `json:"name"`   // 用户名
	Pass   string    `json:"pass"`   // 密码
	Phone  string    `json:"phone"`  // 手机号
	Email  string    `json:"email"`  // 邮箱
	Status int       `json:"status"` // 账号状态
	Ctime  time.Time `json:"ctime"`  // 修改时间
}

// UserClaims token携带的数据
type UserClaims struct {
	Id   int    `json:"id"`
	Num  string `json:"num"`
	Name string `json:"name"`
	jwt.StandardClaims
}

// UserLogin 用户登录
func UserLogin(num string) (User, error) {
	mod := User{}
	// 通过账号查询数据
	err := Db.Get(&mod, "select * from user where num=? limit 1", num)
	return mod, err
}

// 用户列表分页  需要元素 当前页pi 每页大小ps  总共有多少条

// UserPage 用户列表分页
func UserPage(pi, ps int) ([]User, error) {
	// 创建一个容器装查询到的数据
	mods := make([]User, 0, ps) // 给定容器大小，但是现在不用
	err := Db.Select(&mods, "select * from user limit ?,?", (pi-1)*ps, ps)
	return mods, err
}

// UserCount 用户列表总数
func UserCount() int {
	// 创建一个值装查询到的数据
	count := 0
	Db.Get(&count, "select count(id) as count from user")
	return count
}

// UserDel 删除用户 id
func UserDel(id int64) error {
	tx, _ := Db.Begin() // 开启事务  保险箱
	result, err := tx.Exec("delete from user where id = ?", id)
	if err != nil {
		tx.Rollback() // 回滚 撤回刚才的操作
		return err
	}
	// 影响行数
	rows, _ := result.RowsAffected()
	if rows < 1 {
		tx.Rollback()
		return errors.New("rows affected < 1")
	}
	tx.Commit() // 提交命令
	return nil
}

// UserAdd 添加用户 id
func UserAdd(mod *User) error {
	tx, _ := Db.Begin() // 开启事务  保险箱
	result, err := tx.Exec("insert into user (num,`name`,pass,phone,email,ctime,status) values (?,?,?,?,?,?,?)", mod.Num, mod.Name, mod.Pass, mod.Phone, mod.Email, mod.Ctime, mod.Status)
	if err != nil {
		tx.Rollback() // 回滚 撤回刚才的操作
		return err
	}
	// 影响行数
	rows, _ := result.RowsAffected()
	if rows < 1 {
		tx.Rollback()
		return errors.New("rows affected < 1")
	}
	tx.Commit() // 提交命令
	return nil
}

// UserExist 判断用户是否存在
func UserExist(num string) bool {
	mod := User{}
	err := Db.Get(&mod, "select * from user where num = ?", num)
	if err != nil {
		return false
	}
	return true
}
