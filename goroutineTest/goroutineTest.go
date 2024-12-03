package goroutineTest

import (
	"fmt"
	"time"
)

func GoTest() {

	go func() {
		for {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("hello")
		}
	}()
	return
}
