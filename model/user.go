package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
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
