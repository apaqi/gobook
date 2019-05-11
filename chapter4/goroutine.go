package main

import (
	"fmt"
	"time"
)

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}
/**
在 Go 中可以并发执行的活动单元称之为 goroutine。当一个 Go 程序启动时，一个执行 main function 的 goroutine 会被创建，
称之为 main goroutine。创建新的 goroutine 可以使用 go 语句，像这样: go f()，其中 f 是一个函数。
使用 go 语句开启一个新的 goroutine 之后，go 语句之后的函数调用将在新的 goroutine 中执行，而不会阻塞当前的程序执行。
 */
func main() {
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
	//sleep 3秒，要不主线程执行完成退出，而Add还没执行完成
	time.Sleep(time.Duration(3)*time.Second)

}
