package main

import (
	"encoding/json"
	"fmt"
)

var jsonstr = `{
	"code":200,
	"message":"",
	"data":[
		{
			"name":"zhangsan",
			"age":15
		},
		{
			"name":"lisi",
			"age":25
		},
		{
			"name":"wangwu",
			"age":18
		},
		{
			"name":"limazi",
			"age":23
		}
	]
}`

type JsonData struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Name string `json:"name"`
		Age  int64  `json:"age"`
	} `json:"data"`
}

// json解码
func JsonDecode() {
	//json解码
	jsondata := JsonData{}
	if err := json.Unmarshal([]byte(jsonstr), &jsondata); err != nil {
		fmt.Printf("json.Unmarshal failed, err=%v data=%s\n", err, jsonstr)
		return
	}
	fmt.Printf("code:%d,message:%s\n", jsondata.Code, jsondata.Message)
	for _, row := range jsondata.Data {
		fmt.Printf("name:%s\n", row.Name)
		fmt.Printf("age:%d\n", row.Age)
	}
}

func main() {
	JsonDecode()
}
