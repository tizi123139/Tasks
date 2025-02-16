package main

import(
	"net/http"
	"fmt"
	"io/ioutil"
)//io/ioutil用于读取响应内容，net/http用于发送 HTTP 请求

func main() {
    // 定义请求的URL
    url := "http://localhost:8080"

    // 发送HTTP GET请求
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }
    defer resp.Body.Close()

    // 读取响应内容，将内容读取到body中
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return
    }

    // 输出响应内容
    fmt.Println("Response status:", resp.Status)
    fmt.Println("Response body:", string(body))
}