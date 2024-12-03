package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/bytedance/sonic"
	"testing"
)

// 定义一个Person结构体
type Persons struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// generateData 生成测试数据
func generateData(n int) []Persons {
	people := make([]Persons, n)
	for i := range people {
		people[i] = Persons{
			Name: fmt.Sprintf("Persons%d", i+1),
			Age:  i % 100,
		}
	}
	return people
}

// BenchmarkGobEncoding 用于测试Gob序列化的性能
func BenchmarkGobEncoding(b *testing.B) {
	data := generateData(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		encoder := gob.NewEncoder(&buffer)
		if err := encoder.Encode(data); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkJsonEncoding 用于测试JSON序列化的性能
func BenchmarkJsonEncoding(b *testing.B) {
	data := generateData(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := sonic.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}
