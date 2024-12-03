package goroutineTest

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGoTest(t *testing.T) {
	GoTest()
	fmt.Println("hello1111111111")
	time.Sleep(1 * time.Minute)
	context.Background()
}
