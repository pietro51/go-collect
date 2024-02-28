package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("hello world.")
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("err", err)
	}
	closer := resp.Body
	bytes, err := io.ReadAll(closer)
	fmt.Println(string(bytes))
}
