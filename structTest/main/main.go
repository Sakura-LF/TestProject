package main

import (
	"fmt"
	"os"
)

type User struct {
	name string
	Age  int
}

func main() {
	//user := User{
	//	name: "Sakura",
	//	Age:  15,
	//}
	//fmt.Printf("%+v", user)
	//for i := 0; i < 10; i++ {
	//	// 打印i 的地址
	//	fmt.Printf("%p\n", &i)
	//	fmt.Println()
	//}
	fmt.Println(os.TempDir())
}
