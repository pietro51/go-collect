package libs

import (
	"encoding/json"
	"fmt"
)

type JsonData struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Name string `json:"name"`
		Age  int64  `json:"age"`
	} `json:"data"`
}

// json解码
func Decode(jsonstr string) {
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

func Encode(jsonData JsonData) {
	// 将Person对象转换为JSON字符串
	data, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func Test() {
	fmt.Println("test")
}
