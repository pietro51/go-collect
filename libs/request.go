package libs

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

// GET请求
func Get(url string) string {
	fmt.Println("hello world.")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err", err)
	}
	closer := resp.Body
	bytes, _ := io.ReadAll(closer)
	return string(bytes)
}

// Post请求
func Post(url string, data string) string {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Fatal(err)
	}

	// appending to existing query args
	q := req.URL.Query()
	q.Add("foo", "bar")

	// assign encoded query string to http request
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return ""
	}

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	return string(responseBody)
}
