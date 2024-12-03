package sliceTest

import (
	"testing"
	"time"
)

func Test_sliceTest(t *testing.T) {
	sliceTest()
}

func Test_slice(t *testing.T) {
	//var a []int
	//for value := range *a {
	//	t.Log(value)
	//}
}

func Test_slice2(t *testing.T) {
	ch := make(chan string, 5)

	go PutValue(&ch)

	for {

		time.Sleep(1 * time.Second)
		value, _ := <-ch

		t.Log(value)
	}
}

func PutValue(ch *chan string) {

	for i := 0; i < 5; i++ {
		*ch <- "hello"
	}

	close(*ch)
}
