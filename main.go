package main

import (
	"fmt"
)

func main() {
	m := map[string]string{"a": "123", "b": "456"}
	mm := m
	mm["b"] = "789"
	// m := make(map[string]string)
	// m["name"] = "张三"
	// m["age"] = "14"
	fmt.Println(mm)
	fmt.Println(m)

	// utils.Start()
	// utils.Test()
	// getResult := utils.Get("https://www.baidu.com/")
	// fmt.Println(getResult)
	// postResult := utils.Post("https://www.baidu.com/", "")
	// fmt.Println(postResult)

	// var jsonstr = `{
	// 	"code":200,
	// 	"message":"",
	// 	"data":[
	// 		{
	// 			"name":"zhangsan",
	// 			"age":15
	// 		},
	// 		{
	// 			"name":"lisi",
	// 			"age":25
	// 		},
	// 		{
	// 			"name":"wangwu",
	// 			"age":18
	// 		},
	// 		{
	// 			"name":"limazi",
	// 			"age":23
	// 		}
	// 	]
	// }`

}
