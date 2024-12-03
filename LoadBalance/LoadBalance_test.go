package LoadBalance

import (
	"strconv"
	"testing"
)

func TestNewRoundRobinLoadBalance(t *testing.T) {
	balance := NewRoundRobinLoadBalance()
	for i := 0; i < 5; i++ {
		balance.Add("127.0.0.1:" + strconv.Itoa(i))
	}
	//fmt.Println(balance)
	for i := 0; i < 10; i++ {
		t.Log(balance.Get())
	}
}
