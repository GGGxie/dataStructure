package facade

// 门面模式
// 将子系统的多个细粒度的接口封装成一个接口，提供给上层使用
// 假设现在我有一个网站，以前有登录和注册的流程，
// 登录的时候调用用户的查询接口，注册时调用用户的创建接口。
// 为了简化用户的使用流程，我们现在提供直接验证码登录/注册的功能，
// 如果该手机号已注册那么我们就走登录流程，如果该手机号未注册，那么我们就创建一个新的用户。

// IUser 用户接口
type IUser interface {
	Login(phone int, code int) (*User, error)
	Register(phone int, code int) (*User, error)
}

// IUserFacade 门面模式
type IUserFacade interface {
	LoginOrRegister(phone int, code int) error
}

// User 用户
type User struct {
	Name string
}

// UserService UserService
type UserService struct{}

// Login 登录
func (u UserService) Login(phone int, code int) (*User, error) {
	// 校验操作 ...
	return &User{Name: "test login"}, nil
}

// Register 注册
func (u UserService) Register(phone int, code int) (*User, error) {
	// 校验操作 ...
	// 创建用户
	return &User{Name: "test register"}, nil
}

// LoginOrRegister 登录或注册
func (u UserService) LoginOrRegister(phone int, code int) (*User, error) {
	user, err := u.Login(phone, code)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	return u.Register(phone, code)
}
