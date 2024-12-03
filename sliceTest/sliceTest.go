package sliceTest

import (
	"fmt"
	"path/filepath"
	"time"
)

func sliceTest() {
	begin := time.Now()
	//fmt.Println(time.Now())
	defer func() {
		fmt.Println("函数耗时", time.Since(begin).Microseconds())
	}()
	s1 := make([]int, 0, 50)
	fmt.Println(s1)
	//clean := path.Clean("C:\\Users\\18084\\AppData\\Local\\Temp/000000000.data")
	//clean = path.Dir(clean)
	filepath.Join("C:\\Users\\18084\\AppData\\Local\\Temp/000000000.data", "test")
	clean := filepath.Clean("C:\\Users\\18084\\AppData\\Local\\Temp/0000000.data")
	fmt.Println(clean)
	//time.Sleep(time.Second * 2)

	//formatUint := strconv.FormatUint(1, 10)

}
