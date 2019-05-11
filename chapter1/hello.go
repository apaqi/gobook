package main

import "fmt" // 我们需要 fmt 包中的 Println 函数

func main() {
    fmt.Println("Hello, world. 你好,世界!")
}
//如何构建
//切到该文件所在目录
// go build hello.go  --构建
// ./hello            --运行