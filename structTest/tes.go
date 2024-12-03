package structTest

import "fmt"

func GetSum(a, b int) int {
	return a + b
}

type User struct {
	Name string
	Age  int
}

func (u *User) Test() {
	fmt.Println(u.Age)
	u.Test2()
}

func (u *User) Test2() {
	u.Name = "test"
	fmt.Println(u.Name)
}
