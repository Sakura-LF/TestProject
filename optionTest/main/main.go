package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	// 使用函数写法定义的 option 模式
	user := NewUser(WithName("Sakura"), WithAge(20), WithTag("k1", "v1"))
	fmt.Println(user)
	// 不传入任何 option 的 user
	user2 := NewUser()
	fmt.Print(user2)
}

type User struct {
	Name string
	Age  int
	Tags map[string]string
}

// UserOption 第一种写法,函数写法
type UserOption func(u *User)

func WithName(name string) UserOption {
	return func(u *User) {
		u.Name = name
	}
}

func WithAge(age int) UserOption {
	return func(u *User) {
		u.Age = age
	}
}

func WithTag(k, v string) UserOption {
	return func(u *User) {
		if u.Tags == nil {
			u.Tags = make(map[string]string)
		}
		u.Tags[k] = v
	}
}

// 默认用户实现
func DefaultUser() *User {
	return &User{
		Name: uuid.NewString(),
		Age:  0,
	}
}

func NewUser(options ...UserOption) *User {
	// 这里还可以加一个默认user实现,用于用户不进行任何option的
	defaultUser := DefaultUser()
	// 遍历的时候直接对defaultUser进行修改
	//user := &User{}
	// 遍历函数的时候就已经执行了
	for _, option := range options {
		// 这个option其实就是WithName和WithAge的返回值
		// 也就是UseOption这个函数
		// 所以需要参数*User
		option(defaultUser)
	}
	return defaultUser
}
