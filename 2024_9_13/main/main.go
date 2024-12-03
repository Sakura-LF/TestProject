package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"time"
)

// 定义一个Person结构体
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 创建一些数据
	people := make([]Person, 10000)
	for i := range people {
		people[i] = Person{
			Name: fmt.Sprintf("Persons%d", i+1),
			Age:  i % 100,
		}
	}

	// 使用Gob进行序列化
	start := time.Now()
	_ = gobEncode(people)
	gobDuration := time.Since(start)
	fmt.Printf("Gob encoding took %s\n", gobDuration)

	// 使用JSON进行序列化
	start = time.Now()
	_, _ = json.Marshal(people)
	jsonDuration := time.Since(start)
	fmt.Printf("JSON encoding took %s\n", jsonDuration)

	// 检查序列化结果是否相同（为了演示，这里不做）
	// fmt.Printf("Gob data: %v\n", gobData)
	// fmt.Printf("JSON data: %v\n", jsonData)
}

// gobEncode 函数用于使用gob对数据进行编码
func gobEncode(data interface{}) []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}
