package proxy

import (
	"log"
	"time"
)

// 代理模式:在不修改原始类代码的情况下,通过引入代理类来给原始类添加功能
// 常见应用场景:
// 1. 业务的非功能性开发:监控、统计、鉴权...
// 2. 缓存
// 3. RPC
// IUser IUser
type IUser interface {
	Login(username, password string) error
}

// User 用户
type User struct {
}

// Login 用户登录
func (u *User) Login(username, password string) error {
	// 不实现细节
	return nil
}

// UserProxy 代理类
type UserProxy struct {
	user IUser
}

// NewUserProxy NewUserProxy
func NewUserProxy(user IUser) *UserProxy {
	return &UserProxy{
		user: user,
	}
}

// Login 登录，和 user 实现相同的接口
func (p *UserProxy) Login(username, password string) error {
	// before 这里可能会有一些统计的逻辑
	start := time.Now()

	// 这里是原有的业务逻辑
	if err := p.user.Login(username, password); err != nil {
		return err
	}

	// after 这里可能也有一些监控统计的逻辑
	log.Printf("user login cost time: %s", time.Now().Sub(start))

	return nil
}
