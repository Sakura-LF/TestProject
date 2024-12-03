package structTest

import (
	"fmt"
	"github.com/bytedance/sonic"
	"testing"
)

func TestGetSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "一:测试",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name)
			if got := GetSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSum1(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "一:a和b均为正数",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "二:a和b均为负数",
			args: args{
				a: -1,
				b: -2,
			},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName(t *testing.T) {
	user := &User{
		Name: "Sakura",
		Age:  10,
	}

	user.Test()
	fmt.Println(user)
}

func TestSt(t *testing.T) {
	var user User
	fmt.Println(user == nil)
	user2 := &User{}
	fmt.Println(user2 == nil)
	fmt.Println(user2)
	marshal, err := sonic.Marshal(user2)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}
