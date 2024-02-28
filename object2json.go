package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"` // 指定字段名称为 "name"
	Age  int    `json:"age"`
}

func main() {
	p := Person{
		Name: "Alice",
		Age:  25,
	}

	// 将Person对象转换为JSON字符串
	data, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
